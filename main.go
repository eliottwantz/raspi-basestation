package main

import (
	"app/db"
	"app/server"
	"app/server/api"
	"flag"
)

var (
	ipaddr = flag.String("addr", "127.0.0.1", "Specify the ip addr the server is listening on")
)

func main() {
	db.Open()
	flag.Parse()
	go api.Start(ipaddr)
	server.Start(ipaddr)
}
