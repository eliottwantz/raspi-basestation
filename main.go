package main

import (
	"app/database"
	"app/pb"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

const (
	MAXLINE = 1024
	WORKERS = 100
)

var (
	db *gorm.DB
)

func main() {
	db = database.Open()
	ListenUDP()
}

func ListenUDP() {
	conn, err := net.ListenPacket("udp", ":8080")
	handleError(err)
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())
	var wg sync.WaitGroup
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			HandlePacket(conn, db)
		}()
	}
	wg.Wait()
}

func HandlePacket(conn net.PacketConn, db *gorm.DB) {
	defer conn.Close()
	count := 0
	for {
		message := make([]byte, MAXLINE)
		size, addr, err := conn.ReadFrom(message)
		handleError(err)
		var sensorState pb.SensorState
		err = proto.Unmarshal(message[:size], &sensorState)
		handleError(err)
		// handleError(db.Create(&sensorState.MainComputer).Error)
		// handleError(db.Create(&sensorState.BrakeManager).Error)
		count++
		log.Printf("[%s] : COUNT = %d\n", addr, count)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
