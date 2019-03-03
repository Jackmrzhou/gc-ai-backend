set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -a -installsuffix cgo -o gc_ai .
go_swagger generate spec -o swaggerui/swagger.json
del gc_ai.tar
docker build --no-cache -t gc-ai .
docker save gc-ai -o gc_ai.tar