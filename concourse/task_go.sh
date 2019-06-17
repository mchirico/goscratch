#!/bin/bash

go get github.com/axw/gocov/gocov
go install github.com/axw/gocov/gocov
go get -u github.com/mchirico/date/parse
go get gopkg.in/yaml.v2
go get golang.org/x/crypto/ssh
go get github.com/spf13/cobra
go get github.com/mitchellh/go-homedir
go get github.com/spf13/viper


# Now test
# mkdir -p $GOPATH/src/github.com/mchirico/
# cp -r ./resource-tutorial $GOPATH/src/github.com/mchirico/goscratch
cd ./resource-tutorial
go test -race -v -coverprofile=c.out ./...





