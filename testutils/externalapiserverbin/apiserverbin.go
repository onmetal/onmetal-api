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

// Package externalapiserverbin is a test-only package that provides the path to a
// compiled binary of the onmetal-api API server. This is to speed up tests
// by not requiring the compilation each time.
// Caution: Only external packages should use this package.
package externalapiserverbin

import (
	"bytes"
	"fmt"
	cp "github.com/otiai10/copy"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	Path string
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("apiserverbin: unable to determine filename")
	}

	tmpDir, err := os.MkdirTemp("", "onmetal-api")
	if err != nil {
		panic(fmt.Sprintf("failed to create temp directory: %v", err))
	}
	//defer func() { _ = os.RemoveAll(tmpDir) }()

	moduleRoot := filepath.Join(filename, "..", "..", "..")
	Path = filepath.Join(tmpDir, "testbin", "apiserver")

	if err := cp.Copy(moduleRoot, tmpDir, cp.Options{AddPermission: 0666}); err != nil {
		panic(fmt.Sprintf("failed to copy moduleroot: %v", err))
	}

	var out bytes.Buffer
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = tmpDir
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("error running command: %w\noutput: %s", err, out.String()))
	}

	out.Reset()
	cmd = exec.Command("go", "build", "-o",
		Path,
		"main.go",
	)
	cmd.Dir = filepath.Join(tmpDir, "cmd", "apiserver")
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		panic(fmt.Errorf("error running command: %w\noutput: %s", err, out.String()))
	}
}
