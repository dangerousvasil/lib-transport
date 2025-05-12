
clean:
	## proto transport
	-mkdir ptransport
	-mkdir swagger
	-rm ptransport/*
	-rm swagger/*.json

gen:
	-make clean
	protoc --proto_path=proto proto/*.proto  --go_out=:ptransport --go-grpc_out=:ptransport \
 		 --grpc-gateway_out=:ptransport \
 		 --grpc-gateway_opt generate_unbound_methods=true \
 		 --openapiv2_out=:swagger

gen-mocks:
	go generate ./...


server:
	go run cmd/server/main.go

init:
	go install \
       github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
       github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
       google.golang.org/protobuf/cmd/protoc-gen-go \
       google.golang.org/grpc/cmd/protoc-gen-go-grpc



