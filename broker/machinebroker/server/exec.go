// Copyright 2023 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"io"
	"net/http"

	"github.com/go-logr/logr"
	computev1alpha1 "github.com/onmetal/onmetal-api/api/compute/v1alpha1"
	"github.com/onmetal/onmetal-api/client-go/onmetalapi"
	onmetalapiclientgoscheme "github.com/onmetal/onmetal-api/client-go/onmetalapi/scheme"
	ori "github.com/onmetal/onmetal-api/ori/apis/machine/v1alpha1"
	remotecommandserver "github.com/onmetal/onmetal-api/poollet/machinepoollet/ori/streaming/remotecommand"
	"k8s.io/client-go/tools/remotecommand"
)

func (s *Server) Exec(ctx context.Context, req *ori.ExecRequest) (*ori.ExecResponse, error) {
	machineID := req.MachineId
	log := s.loggerFrom(ctx, "MachineID", machineID)

	log.V(1).Info("Inserting request into cache")
	token, err := s.execRequestCache.Insert(req)
	if err != nil {
		return nil, err
	}

	log.V(1).Info("Returning url with token")
	return &ori.ExecResponse{
		Url: s.buildURL("exec", token),
	}, nil
}

func (s *Server) ServeExec(w http.ResponseWriter, req *http.Request, token string) {
	ctx := req.Context()
	log := logr.FromContextOrDiscard(ctx)

	request, ok := s.execRequestCache.Consume(token)
	if !ok {
		log.V(1).Info("Rejecting unknown / expired token")
		http.NotFound(w, req)
		return
	}

	onmetalClientset, err := onmetalapi.NewForConfig(s.cluster.Config())
	if err != nil {
		log.Error(err, "Error getting onmetal api clientset for config")
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	reqURL := onmetalClientset.ComputeV1alpha1().RESTClient().
		Post().
		Namespace(s.cluster.Namespace()).
		Resource("machines").
		Name(request.MachineId).
		SubResource("exec").
		VersionedParams(&computev1alpha1.MachineExecOptions{}, onmetalapiclientgoscheme.ParameterCodec).
		URL()

	executor, err := remotecommand.NewSPDYExecutor(s.cluster.Config(), http.MethodGet, reqURL)
	if err != nil {
		log.Error(err, "Error creating remote command executor")
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	exec := executorExec{executor}
	handler, err := remotecommandserver.NewExecHandler(exec, remotecommandserver.ExecHandlerOptions{})
	if err != nil {
		log.Error(err, "Error creating exec handler")
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	handler.Handle(w, req, remotecommandserver.ExecOptions{})
}

type executorExec struct {
	executor remotecommand.Executor
}

func (e executorExec) Exec(ctx context.Context, in io.Reader, out io.WriteCloser, resize remotecommand.TerminalSizeQueue) error {
	return e.executor.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdin:             in,
		Stdout:            out,
		Stderr:            nil,
		Tty:               true, // TODO: Obtain this value properly
		TerminalSizeQueue: resize,
	})
}
