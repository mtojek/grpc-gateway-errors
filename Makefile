build: go-get generate-grpc-stub
	go get github.com/mtojek/grpc-gateway-errors/server
	go test ./...

go-get:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

generate-grpc-stub:
	protoc  -I/usr/local/include -I. \
             -I${GOPATH}/src \
             -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
             --go_out=plugins=grpc:. ./api/hello.proto \
             ./api/hello.proto
	protoc  -I/usr/local/include -I. \
             -I${GOPATH}/src \
             -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
             --grpc-gateway_out=logtostderr=true:. \
             ./api/hello.proto

