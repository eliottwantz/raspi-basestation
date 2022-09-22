proto:
	protoc --proto_path=. --go_out=. \
	--plugin=./web/node_modules/.bin/protoc-gen-ts_proto \
	--ts_proto_opt=onlyTypes=true \
	--ts_proto_opt=useOptionals=none \
	--ts_proto_out=./web/ \
	pb/*.proto

run:
	rm -f polyloop.sqlite3
	go run .

sqlc:
	sqlc compile

sql:
	sqlc generate

.PHONY:
	proto
	run
	sqlc
	sql