#!/bin/bash

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

PKG_DIR="github.com/alejandroEsc/k8s-provisioner-juju-example/pkg"
PKG_REL="./pkg"
PKG_CLIENT="${PKG_DIR}/client"
PKG_APIS="${PKG_DIR}/apis"

# Deep gen, manually calling this works.
echo "generationg deepcopy"
${GOPATH}/bin/deepcopy-gen \
--input-dirs "${PKG_APIS}/controller/v1alpha1" \
-O zz_generated.deepcopy

# client
echo "generating client group"
${GOPATH}/bin/client-gen \
--clientset-name versioned \
--input-base "${PKG_APIS}" \
--input "controller/v1alpha1" \
--output-package "${PKG_CLIENT}/clientset"

# listers
echo "generating listers group"
${GOPATH}/bin/lister-gen \
--input-dirs "${PKG_APIS}/controller/v1alpha1" \
--output-package "${PKG_CLIENT}/listers"

# informer
echo "generating informers group"
${GOPATH}/bin/informer-gen \
--input-dirs "${PKG_APIS}/controller/v1alpha1" \
--versioned-clientset-package ${PKG_CLIENT}/clientset/versioned \
--listers-package "${PKG_CLIENT}/listers" \
--output-package "${PKG_CLIENT}/informers"
