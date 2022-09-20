proto:
	protoc --proto_path=. --go_out=. pb/*.proto

sql:
	sqlc generate

sqlc:
	sqlc compile

run:
	rm -f database/polyloop.db
	go run .

.PHONY:
	proto
	sqlc
	sqlcheck