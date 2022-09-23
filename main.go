package main

import (
	"app/db"
	"app/server"
	"app/server/api"
)

func main() {
	db.Open()
	go api.Start()
	server.Start()
}
