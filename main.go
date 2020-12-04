package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"os"
	"path/filepath"
	"strings"
)

func errorExit(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

var (
	Version  = "v0.0.0"
	Revision = ""
)

func main() {
	var (
		region  string
		login   bool
		version bool
		rawJSON bool
	)

	flag.StringVar(&region, "region", "", "region")
	flag.BoolVar(&login, "login", false, "get-login with --no-include-email compatible mode")
	flag.BoolVar(&version, "version", false, "show version")
	flag.BoolVar(&rawJSON, "json", false, "output json")
	flag.Parse()

	if version {
		fmt.Printf("%s version %s(%s)\n", filepath.Base(os.Args[0]), Version, Revision)
		os.Exit(0)
	}

	conf := aws.NewConfig()
	if region != "" {
		conf = conf.WithRegion(region)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            *conf,
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := ecr.New(sess)
	ctx := context.Background()

	res, err := client.GetAuthorizationTokenWithContext(ctx, &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		errorExit("Can not get authorization token: %s\n", err)
	} else if len(res.AuthorizationData) != 1 {
		errorExit("Invalid GetAuthToken response")
	}

	authToken := res.AuthorizationData[0]
	data, err := base64.StdEncoding.DecodeString(*authToken.AuthorizationToken)
	if err != nil {
		errorExit("Can not decode authorization token: %s\n", err)
	}

	token := strings.SplitN(string(data), ":", 2)
	if login {
		fmt.Println(fmt.Sprintf("docker login -u %s -p %s %s", token[0], token[1], *authToken.ProxyEndpoint))
	} else if rawJSON {
		enc := json.NewEncoder(os.Stdout)
		if err := enc.Encode(authToken); err != nil {
			errorExit("Can not encode to JSON: %s\n", err)
		}
	} else {
		fmt.Println(token[1])
	}

	os.Exit(0)
}
