all:
	protoc --proto_path=. --go_out=. pb/sensor.proto

go:
	protoc --proto_path=. --go_out=. pb/sensor.proto
