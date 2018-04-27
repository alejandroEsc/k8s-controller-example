GIT_SHA=$(shell git rev-parse --verify HEAD)
GOBUILD_CONTROLLER=go build -o bin/controller -ldflags "-X github.com/alejandroEsc/k8s-controller-example/pkg/version.GitSha=${GIT_SHA}"
GOBUILD_POD_SERVER=go build  -o bin/podapp -ldflags "-X github.com/alejandroEsc/k8s-controller-example/pkg/version.GitSha=${GIT_SHA}"

clean:
	rm -rf bin/*

compile-controller:
	${GOBUILD_CONTROLLER} ./cmd/controller

compile-controller-linux:
	${GOBUILD_CONTROLLER} -o bin/linux-amd64 ./cmd/controller

compile-pod-server:
	${GOBUILD_POD_SERVER} ./cmd/podapp

compile-pod-server-linux:
	${GOBUILD_POD_SERVER} -o bin/linux-amd64 ./cmd/podapp

go-checks:
	./scripts/verify.sh

go-clean:
	./scripts/clean/gofmt-clean.sh
	./scripts/clean/goimports-clean.sh

update-codegen:
	./scripts/update-codegen.sh

generate-pod-apis:
	./scripts/gen_pod_apis.sh


.PHONY: install-tools
install-tools:
	GOIMPORTS_CMD=$(shell command -v goimports 2> /dev/null)
ifndef GOIMPORTS_CMD
	go get golang.org/x/tools/cmd/goimports
endif

	GOLINT_CMD=$(shell command -v golint 2> /dev/null)
ifndef GOLINT_CMD
	go get github.com/golang/lint/golint
endif

	GOCYCLO_CMD=$(shell command -v gocyclo 2> /dev/null)
ifndef GOLINT_CMD
	go get github.com/fzipp/gocyclo
endif


.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
