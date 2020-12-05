REVISION = $(shell git rev-parse HEAD)
VERSION = $(shell git describe --tags --abbrev=0)

all: build

build:
	CGO_ENABLED=0 \
		go build -ldflags "-s -w -X 'main.Version=$(VERSION)' -X 'main.Revision=$(REVISION)'"

image:
	docker build --build-arg VERSION=$(VERSION) --build-arg REVISION=$(REVISION) -t ecr-get-login-password .