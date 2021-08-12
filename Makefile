clean:
	rm -rf pb/*

gen-proto:
	mkdir -p pb
	protoc --proto_path=proto proto/*.proto  --go_out=pb --go-grpc_out=pb

build:
	go build -o grpc-example .

run:
	go build -o grpc-example .
	./grpc-example