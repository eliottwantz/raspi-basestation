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
	db            *gorm.DB
	sensor_states []*pb.SensorState = make([]*pb.SensorState, 0)
)

func main() {
	db = database.Open()
	receive := make(chan *pb.SensorState)
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

func ListenUDP(receive chan *pb.SensorState, quit chan bool) {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	handleError(err)
	conn, err := net.ListenUDP("udp", addr)
	handleError(err)
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for i := 0; i < WORKERS; i++ {
		go HandlePacket(conn, receive)
	}
	go func() {
		for s := range receive {
			sensor_states = append(sensor_states, s)
		}
	}()
	<-quit
	close(receive)
	conn.Close()
	fmt.Println("\nsensor_states COUNT =", len(sensor_states))
	fmt.Println("Writing to db")
	writeToDb()
}

func HandlePacket(conn *net.UDPConn, receive chan *pb.SensorState) {
	count := 0
	for {
		message := make([]byte, MAXLINE)
		size, addr, err := conn.ReadFrom(message)
		if err != nil {
			// log.Printf("Connection closed: %v", err)
			return
		}
		var sensorState pb.SensorState
		err = proto.Unmarshal(message[:size], &sensorState)
		handleError(err)
		receive <- &sensorState
		// handleError(db.Create(&sensorState.MainComputer).Error)
		// handleError(db.Create(&sensorState.BrakeManager).Error)
		count++
		log.Printf("[%s] : COUNT = %d\n", addr, count)
	}
}

func writeToDb() {
	for _, ss := range sensor_states {
		handleError(db.Create(ss.MainComputer).Error)
		handleError(db.Create(ss.BrakeManager).Error)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
