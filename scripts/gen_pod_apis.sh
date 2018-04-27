#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

source ${SCRIPT_ROOT}/scripts/common.sh

PROTO_COMMON="protoc ${SCRIPT_ROOT}/api/pod.proto -I /usr/local/include -I ${SCRIPT_ROOT}/api -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway"

inf "generating api stubs..."
${PROTO_COMMON} --go_out=plugins=grpc:${SCRIPT_ROOT}/api

inf "generating REST gateway stubs..."
${PROTO_COMMON} --grpc-gateway_out=logtostderr=true:${SCRIPT_ROOT}/api

inf "generating swagger docs..."
${PROTO_COMMON} --swagger_out=logtostderr=true:${SCRIPT_ROOT}/swagger

inf "generating api docs..."
${PROTO_COMMON} --doc_out docs --doc_opt=markdown,api.md