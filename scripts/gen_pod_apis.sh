#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

source ${SCRIPT_ROOT}/scripts/common.sh

echo
inf "generating api stubs..."
inf "protoc -I ${SCRIPT_ROOT}/api/ ${SCRIPT_ROOT}/api/pod.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --go_out=plugins=grpc:${SCRIPT_ROOT}/api"
protoc -I ${SCRIPT_ROOT}/api/ ${SCRIPT_ROOT}/api/pod.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --go_out=plugins=grpc:${SCRIPT_ROOT}/api

echo
inf "generating REST gateway stubs..."
inf "protoc -I /usr/local/include/ -I ${SCRIPT_ROOT}/api/ ${SCRIPT_ROOT}/api/pod.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${SCRIPT_ROOT}/api"
protoc -I /usr/local/include/ -I ${SCRIPT_ROOT}/api/ ${SCRIPT_ROOT}/api/pod.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:${SCRIPT_ROOT}/api

echo
inf "generating swagger docs..."
inf "protoc -I /usr/local/include/ -I ${SCRIPT_ROOT}/api/ ${SCRIPT_ROOT}/api/pod.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${SCRIPT_ROOT}/swagger"
protoc -I /usr/local/include/ -I ${SCRIPT_ROOT}/api/ ${SCRIPT_ROOT}/api/pod.proto -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:${SCRIPT_ROOT}/swagger

echo
inf "generating api docs..."
inf "protoc -I api/ -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway api/pod.proto --doc_out docs --doc_opt=markdown,api.md"
protoc -I api/ -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/grpc-ecosystem/grpc-gateway api/pod.proto --doc_out docs --doc_opt=markdown,api.md