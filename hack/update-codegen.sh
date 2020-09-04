#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

rm -rf vendor # just in case this didn't get cleaned up
go mod vendor # required by codegenerators

# Change these
PKG="github.com/pwittrock/apiserver-runtime-sample"
GROUP="sample"

# update versions
APIS="$GROUP:v1alpha1"
VERSIONED="$PKG/pkg/apis/$GROUP/v1alpha1"

GENERATED="$PKG/pkg/generated"

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# clean up generated code
find . -name "zz_generated.*" | xargs rm
rm -rf pkg/generated/clientset
rm -rf pkg/generated/informers
rm -rf pkg/generated/listers

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash "${CODEGEN_PKG}/generate-groups.sh" all  $PKG/pkg/generated $PKG/pkg/apis \
  $APIS  --output-base "${SCRIPT_ROOT}/../../.." --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt"

openapi-gen --input-dirs k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version,$VERSIONED \
  --output-package $GENERATED/openapi -O zz_generated.openapi --output-base ../../.. --go-header-file ./hack/boilerplate.go.txt

# To use your own boilerplate text append:
#   --go-header-file "${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt"

rm -rf vendor # get rid of this