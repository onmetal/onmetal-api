#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

"$SCRIPT_DIR"/update-codegen-api.sh
"$SCRIPT_DIR"/update-codegen-client-go.sh
"$SCRIPT_DIR"/update-codegen-onmetal-apiserver.sh
