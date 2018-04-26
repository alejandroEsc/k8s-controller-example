#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
PKG_DIR="github.com/alejandroEsc/k8s-controller-example/pkg"
PKG_REL="./pkg"
PKG_CLIENT="${PKG_DIR}/client"
PKG_APIS="${PKG_DIR}/apis"

# Deep gen, manually calling this works.
echo "generationg deepcopy"
if [ ! -e ${GOPATH}/bin/deepcopy-gen ]; then
    echo "installing deepcopy-gen tool"
    cd ${SCRIPT_ROOT}
    go install ${SCRIPT_ROOT}/vendor/k8s.io/code-generator/cmd/deepcopy-gen
fi
${GOPATH}/bin/deepcopy-gen \
--input-dirs "${PKG_APIS}/controller/v1alpha1" \
--go-header-file ${SCRIPT_ROOT}/scripts/boilerplate.go.txt \
-O zz_generated.deepcopy

# client
echo "generating client group"
if [ ! -e ${GOPATH}/bin/client-gen ]; then
    echo "installing client-gen tool"
    cd ${SCRIPT_ROOT}
    go install ${SCRIPT_ROOT}/vendor/k8s.io/code-generator/cmd/client-gen
fi
${GOPATH}/bin/client-gen \
--clientset-name versioned \
--input-base "${PKG_APIS}" \
--input "controller/v1alpha1" \
--go-header-file ${SCRIPT_ROOT}/scripts/boilerplate.go.txt \
--output-package "${PKG_CLIENT}/clientset"

# listers
echo "generating listers group"
if [ ! -e ${GOPATH}/bin/lister-gen ]; then
    echo "installing lister-gen tool"
    cd ${SCRIPT_ROOT}
    go install ${SCRIPT_ROOT}/vendor/k8s.io/code-generator/cmd/lister-gen
fi
${GOPATH}/bin/lister-gen \
--input-dirs "${PKG_APIS}/controller/v1alpha1" \
--go-header-file ${SCRIPT_ROOT}/scripts/boilerplate.go.txt \
--output-package "${PKG_CLIENT}/listers"

# informer
echo "generating informers group"
if [ ! -e ${GOPATH}/bin/informer-gen ]; then
    echo "installing informer-gen tool"
    cd ${SCRIPT_ROOT}
    go install ${SCRIPT_ROOT}/vendor/k8s.io/code-generator/cmd/informer-gen
fi
${GOPATH}/bin/informer-gen \
--input-dirs "${PKG_APIS}/controller/v1alpha1" \
--versioned-clientset-package ${PKG_CLIENT}/clientset/versioned \
--listers-package "${PKG_CLIENT}/listers" \
--go-header-file ${SCRIPT_ROOT}/scripts/boilerplate.go.txt \
--output-package "${PKG_CLIENT}/informers"
