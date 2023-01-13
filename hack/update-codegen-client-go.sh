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
echo "VIRTUAL_GOPATH is $VIRTUAL_GOPATH"
#trap 'rm -rf "$GOPATH"' EXIT

# Setup virtual GOPATH so the codegen tools work as expected.
(cd "$SCRIPT_DIR/../client-go"; go mod download && "$VGOPATH" "$VIRTUAL_GOPATH")

export GOROOT="${GOROOT:-"$(go env GOROOT)"}"
export GOPATH="$VIRTUAL_GOPATH"
export GO111MODULE=off

echo "${BOLD}client-go${NORMAL}"

echo "Generating ${BLUE}openapi${NORMAL}"
"$OPENAPI_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "$(qualify-gvs "github.com/onmetal/onmetal-api/api" "common:v1alpha1 compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")" \
  --input-dirs "k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version" \
  --input-dirs "k8s.io/api/core/v1" \
  --input-dirs "k8s.io/apimachinery/pkg/api/resource" \
  --output-package "github.com/onmetal/onmetal-api/client-go/openapi" \
  -O zz_generated.openapi \
  --report-filename "$SCRIPT_DIR/../client-go/openapi/api_violations.report"

echo "Generating ${BLUE}applyconfiguration${NORMAL}"
applyconfigurationgen_external_apis+=("k8s.io/apimachinery/pkg/apis/meta/v1")
applyconfigurationgen_external_apis+=("$(qualify-gvs "github.com/onmetal/onmetal-api/api" "common:v1alpha1 compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")")
applyconfigurationgen_external_apis_csv=$(IFS=,; echo "${applyconfigurationgen_external_apis[*]}")
"$APPLYCONFIGURATION_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "${applyconfigurationgen_external_apis_csv}" \
  --openapi-schema <("$MODELS_SCHEMA" --openapi-package "github.com/onmetal/onmetal-api/client-go/openapi" --openapi-title "onmetal-api") \
  --output-package "github.com/onmetal/onmetal-api/client-go/applyconfigurations"

echo "Generating ${BLUE}client${NORMAL}"
"$CLIENT_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "$(qualify-gvs "github.com/onmetal/onmetal-api/api" "common:v1alpha1 compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")" \
  --input "$(qualify-gvs "github.com/onmetal/onmetal-api/api" "compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")" \
  --output-package "github.com/onmetal/onmetal-api/client-go" \
  --apply-configuration-package "github.com/onmetal/onmetal-api/client-go/applyconfigurations" \
  --clientset-name "onmetalapi" \
  --input-base ""

echo "Generating ${BLUE}lister${NORMAL}"
"$LISTER_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "$(qualify-gvs "github.com/onmetal/onmetal-api/api" "compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")" \
  --output-package "github.com/onmetal/onmetal-api/client-go/listers"

echo "Generating ${BLUE}informer${NORMAL}"
"$INFORMER_GEN" \
  --output-base "$GOPATH/src" \
  --go-header-file "$SCRIPT_DIR/boilerplate.go.txt" \
  --input-dirs "$(qualify-gvs "github.com/onmetal/onmetal-api/api" "compute:v1alpha1 ipam:v1alpha1 networking:v1alpha1 storage:v1alpha1")" \
  --versioned-clientset-package "github.com/onmetal/onmetal-api/client-go/onmetalapi" \
  --listers-package "github.com/onmetal/onmetal-api/client-go/listers" \
  --output-package "github.com/onmetal/onmetal-api/client-go/informers" \
  --single-directory
