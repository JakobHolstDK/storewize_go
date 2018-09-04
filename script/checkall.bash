#!/bin/bash -e
cd storwize
diff -u <(echo -n) <(gofmt -d -s .)
go vet ./...
golint -set_exit_status ./...
go test ./...
