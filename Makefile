all: build

build:
	CGO_ENABLED=0 go build .

image:
	docker build -t ecr-get-login-password .