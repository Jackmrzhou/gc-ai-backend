#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o gc_ai .
swagger generate spec -o swaggerui/swagger.json
rm gc_ai.tar
docker build --no-cache -t gc-ai .
docker save gc-ai -o gc_ai.tar