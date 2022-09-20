proto:
	protoc --proto_path=. --go_out=. pb/*.proto

sql:
	sqlc generate

sqlc:
	sqlc compile

.PHONY:
	proto
	sqlc
	sqlcheck