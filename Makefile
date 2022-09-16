all:
	protoc --proto_path=. --go_out=. sensor.proto
	mkdir -p ./client/pb
	protoc --proto_path=. --cpp_out=./client/pb sensor.proto

go:
	protoc --proto_path=. --go_out=. sensor.proto

cpp:
	mkdir -p ./client/pb
	protoc --proto_path=. --cpp_out=./client/pb sensor.proto