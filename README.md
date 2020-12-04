# ecr-get-login-password

The lightweight utility for log in to an AWS Elastic Container Registry.

This command compatible with the official CLI of `aws ecr get-login-password`.

Please see AWS CLI official docs.

[ecr-login-password](https://docs.aws.amazon.com/cli/latest/reference/ecr/get-login-password.html)

## Usage

```
Usage of ecr-get-login-password:
  -json
        output json
  -login
        get-login with --no-include-email compatible mode
  -region string
        region
  -version
        show version
```

Get login password and login to ECR.
```shell script
$ ecr-get-login-password \
  | docker login --username AWS \
                 --password-stdin <aws_account_id>.dkr.ecr.<region>.amazonaws.com
```

## Installation

### Install from releases
Please see [Releases](https://github.com/yacchi/ecr-get-login-password/releases/latest).

### Install from github
```shell script
$ go get github.com/yacchi/ecr-get-login-password
```

### Install into your docker application

```dockerfile
COPY --from=ecr-get-login-password:v1.0.0 /ecr-get-login-password /usr/bin/
```

### Build from source
```shell script
$ make
```
