FROM golang:1.15.6 as build

ARG VERSION=v1.0.1
ARG REVISION

WORKDIR /go/src/main

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X 'main.Version=$VERSION' -X 'main.Revision=$REVISION'" -o /main .

FROM scratch

COPY --from=golang:1.15.6 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /main /ecr-get-login-password

ENV HOME=/
ENTRYPOINT ["/ecr-get-login-password"]