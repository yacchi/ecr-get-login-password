FROM golang:1.15.6 as build

WORKDIR /go/src/main

COPY go.mod go.sum ./
COPY main.go ./

RUN go mod download
RUN CGO_ENABLED=0 go build -o /main .

FROM scratch

COPY --from=golang:1.15.6 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /main /ecr-get-login-password

ENV HOME=/
CMD ["/ecr-get-login-password"]