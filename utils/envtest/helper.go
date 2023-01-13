// Copyright 2022 OnMetal authors
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

package envtest

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-getter/v2"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

var log = controllerruntime.Log.WithName("test-env")

func GetPath(ctx context.Context, dst string, path string) (string, error) {
	res, err := getter.DefaultClient.Get(ctx, &getter.Request{
		Src:     path,
		Dst:     dst,
		GetMode: getter.ModeDir,
		Inplace: true,
	})
	if err != nil {
		return "", fmt.Errorf("error getting path %q: %w", path, err)
	}

	return res.Dst, nil
}

func IterateGetPaths(
	ctx context.Context,
	dst string,
	paths []string,
	f func(path, resolved string, err error) error,
) error {
	for _, path := range paths {
		resolved, err := GetPath(ctx, dst, path)
		if err := f(path, resolved, err); err != nil {
			return err
		}
	}
	return nil
}

func ReadDirRegularFiles(name string) ([]os.FileInfo, error) {
	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, fmt.Errorf("error reading entries of path %q: %w", name, err)
	}

	var files []os.FileInfo
	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}

		file, err := entry.Info()
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", entry.Name(), err)
		}

		files = append(files, file)
	}
	return files, nil
}

func IterateDirFiles(name string, f func(info os.FileInfo, err error) error) error {
	entries, err := os.ReadDir(name)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}

		info, err := entry.Info()
		if err := f(info, err); err != nil {
			return err
		}
	}

	return nil
}

// mergePaths merges two string slices containing paths.
// This function makes no guarantees about order of the merged slice.
func mergePaths(s1, s2 []string) []string {
	m := make(map[string]struct{})
	for _, s := range s1 {
		m[s] = struct{}{}
	}
	for _, s := range s2 {
		m[s] = struct{}{}
	}
	merged := make([]string, len(m))
	i := 0
	for key := range m {
		merged[i] = key
		i++
	}
	return merged
}

// mergeAPIServices merges two APIService slices using their names.
// This function makes no guarantees about order of the merged slice.
func mergeAPIServices(s1, s2 []*apiregistrationv1.APIService) []*apiregistrationv1.APIService {
	m := make(map[string]*apiregistrationv1.APIService)
	for _, obj := range s1 {
		m[obj.GetName()] = obj
	}
	for _, obj := range s2 {
		m[obj.GetName()] = obj
	}
	merged := make([]*apiregistrationv1.APIService, len(m))
	i := 0
	for _, obj := range m {
		merged[i] = obj.DeepCopy()
		i++
	}
	return merged
}
