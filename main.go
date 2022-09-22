package main

import (
	"app/db"
	"app/server"
)

func main() {
	db.Open()
	server.Start()
}
