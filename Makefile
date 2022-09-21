proto:
	protoc --proto_path=. --go_out=. pb/*.proto

sql:
	sqlc generate

sqlc:
	sqlc compile

run:
	rm -f db/polyloop.sqlite3
	go run .

.PHONY:
	proto
	sqlc
	sqlcheck