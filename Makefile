.PHONY: all goclean gofmt goget gotest loader projectcheck

GOPATH := ${PWD}/.workspace
export GOPATH

all: goget loader

goclean:
	@rm -rf coverage.txt loader.go profile.out .workspace/

gofmt:
	@go fmt ./...

goget:
	@# Install project dependencies.
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@go get -d -v ./...
	@# Install dev dependencies.
	@go get github.com/client9/misspell/cmd/misspell
	@go get github.com/fzipp/gocyclo
	@go get github.com/golang/lint/golint

gotest:
	@./go.test.sh \

loader:
	@go build \
		-o .workspace/bin/loader \
		-ldflags "-X main.version=$(shell git rev-parse --short HEAD)" \
		github.com/xh3b4sd/loader

projectcheck:
	@./project.check.sh
