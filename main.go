package main

import (
	"app/database"
	"app/pb"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const (
	MAXLINE = 1024
	WORKERS = 100
)

var (
	db    *gorm.DB
	count = 0
)

func main() {
	db = database.Open()
	receive := make(chan int)
	quit := make(chan bool)
	handleInterrupt(quit)
	ListenUDP(receive, quit)
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

func ListenUDP(receive chan int, quit chan bool) {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	handleError(err)
	conn, err := net.ListenUDP("udp", addr)
	handleError(err)
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for i := 0; i < WORKERS; i++ {
		go HandlePacket(conn, receive)
	}
	go func() {
		for c := range receive {
			count += c
		}
	}()
	<-quit
	close(receive)
	conn.Close()
	fmt.Println("\nCOUNT =", count)
}

func HandlePacket(conn *net.UDPConn, receive chan int) {
	count := 0
	for {
		message := make([]byte, MAXLINE)
		size, addr, err := conn.ReadFrom(message)
		if err != nil {
			return
		}
		var sensorState pb.SensorState
		err = proto.Unmarshal(message[:size], &sensorState)
		handleError(err)
		handleError(db.Create(&sensorState.MainComputer).Error)
		handleError(db.Create(&sensorState.BrakeManager).Error)
		receive <- 1
		count++
		log.Printf("[%s] : COUNT = %d\n", addr, count)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
