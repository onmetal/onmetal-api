#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
source "$SCRIPT_DIR/common.sh"

VGOPATH="$VGOPATH"
MODELS_SCHEMA="$MODELS_SCHEMA"
CLIENT_GEN="$CLIENT_GEN"
DEEPCOPY_GEN="$DEEPCOPY_GEN"
LISTER_GEN="$LISTER_GEN"
INFORMER_GEN="$INFORMER_GEN"
DEFAULTER_GEN="$DEFAULTER_GEN"
CONVERSION_GEN="$CONVERSION_GEN"
OPENAPI_GEN="$OPENAPI_GEN"
APPLYCONFIGURATION_GEN="$APPLYCONFIGURATION_GEN"

VIRTUAL_GOPATH="$(mktemp -d)"
trap 'rm -rf "$GOPATH"' EXIT

# Setup virtual GOPATH so the codegen tools work as expected.
(cd "$SCRIPT_DIR/../onmetal-apiserver"; go mod download && "$VGOPATH" "$VIRTUAL_GOPATH")

export GOROOT="${GOROOT:-"$(go env GOROOT)"}"
export GOPATH="$VIRTUAL_GOPATH"
export GO111MODULE=off

echo "${BOLD}onmetal-apiserver${NORMAL}"

echo "Generating ${BLUE}deepcopy${NORMAL}"
"$DEEPCOPY_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "$(qualify-gs "github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis" "compute ipam networking storage")" \
  -O zz_generated.deepcopy

echo "Generating ${BLUE}defaulter${NORMAL}"
"$DEFAULTER_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "$(qualify-gvs "github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis" "compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")" \
  -O zz_generated.defaults

echo "Generating ${BLUE}conversion${NORMAL}"
"$CONVERSION_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "$(qualify-gs "github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis" "compute ipam networking storage")" \
  --input-dirs "$(qualify-gvs "github.com/onmetal/onmetal-api/onmetal-apiserver/internal/apis" "compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")" \
  -O zz_generated.conversion
