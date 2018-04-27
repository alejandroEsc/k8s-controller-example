GIT_SHA=$(shell git rev-parse --verify HEAD)
GOBUILD_CONTROLLER=go build -o bin/controller -ldflags "-X github.com/alejandroEsc/k8s-controller-example/pkg/version.GitSha=${GIT_SHA}"
GOBUILD_POD_SERVER=go build -o bin/podapp -ldflags "-X github.com/alejandroEsc/k8s-controller-example/pkg/version.GitSha=${GIT_SHA}"

clean: ## clean build output
	rm -rf bin/*

compile-controller: # build controller and place in local bin directory
	${GOBUILD_CONTROLLER} ./cmd/controller

compile-controller-linux: ## build linux version of the controller
	GOOS=linux GOARCH=amd64 ${GOBUILD_CONTROLLER} -o bin/linux-amd64 ./cmd/controller

compile-pod-server: ## build pod server and place in local bin directory
	${GOBUILD_POD_SERVER} ./cmd/podapp

compile-pod-server-linux: ## build linux version of the pod server
	GOOS=linux GOARCH=amd64 ${GOBUILD_POD_SERVER} -o bin/linux-amd64 ./cmd/podapp

go-lint-checks: ## run linting checks against golang code
	./scripts/verify.sh

go-clean: ## have gofmt and goimports clean up go code
	./scripts/clean/gofmt-clean.sh
	./scripts/clean/goimports-clean.sh

update-codegen: ## generate clientsets, informers,  and listers from versioned apis and types
	./scripts/update-codegen.sh

generate-pod-apis: ## generate golang stub from pod proto file
	./scripts/gen_pod_apis.sh

.PHONY: install-tools
install-tools: ## install tools needed by go-link-checks
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
