package main

import (
	"app/database"
	"app/server"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	database.Open()
	ssc := make(chan int)
	sdc := make(chan int)
	quit := make(chan bool)
	handleInterrupt(quit)
	server.ListenUDP(ssc, sdc, quit)
}

func handleInterrupt(quit chan bool) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		quit <- true
	}()
}
