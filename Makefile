proto:
	protoc --proto_path=. --go_out=. \
	--plugin=./web/node_modules/.bin/protoc-gen-ts_proto \
	--ts_proto_out=./web/ \
	pb/*.proto

run:
	rm -f db/polyloop.sqlite3
	go run .

.PHONY:
	proto
	run