FROM golang:1.20.0-alpine

# make work directory
RUN mkdir -p /go/src/github.com/masahiro331/sample
WORKDIR /go/src/github.com/masahiro331/sample/


# initial workspace
ADD main.go  /go/src/github.com/masahiro331/sample/main.go
RUN go mod init && go mod tidy

# install trivy
RUN apk add --update curl
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.42.1

CMD ["trivy", "fs", "--skip-dirs", "/usr/local/bin/", "--list-all-pkgs", "-f", "cyclonedx", "/go" ]