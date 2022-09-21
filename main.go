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
	db      *gorm.DB
	ssCount = 0
	sdCount = 0
)

func main() {
	db = database.Open()
	ssc := make(chan int)
	sdc := make(chan int)
	quit := make(chan bool)
	handleInterrupt(quit)
	ListenUDP(ssc, sdc, quit)
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

func ListenUDP(ssc chan int, sdc chan int, quit chan bool) {
	ssAddr, err := net.ResolveUDPAddr("udp", ":8080")
	handleError(err)
	ssConn, err := net.ListenUDP("udp", ssAddr)
	handleError(err)

	sdAddr, err := net.ResolveUDPAddr("udp", ":8081")
	handleError(err)
	sdConn, err := net.ListenUDP("udp", sdAddr)
	handleError(err)

	fmt.Println("server listening \n", ssConn.LocalAddr(), sdConn.LocalAddr())

	for i := 0; i < WORKERS; i++ {
		go HandleSensorState(ssConn, ssc)
		go HandleSensorData(sdConn, sdc)
	}
	go func() {
		for c := range ssc {
			ssCount += c
		}
	}()
	go func() {
		for c := range sdc {
			sdCount += c
		}
	}()
	<-quit
	close(ssc)
	close(sdc)
	ssConn.Close()
	fmt.Println("\nSensorState COUNT =", ssCount)
	fmt.Println("SensorData COUNT =", sdCount)

	// var mc database.MainComputer
	// if err := db.Last(&mc).Error; err != nil {
	// 	handleError(err)
	// }
	// var bm database.BrakeManager
	// if err := db.Last(&bm).Error; err != nil {
	// 	handleError(err)
	// }
}

func HandleSensorState(conn *net.UDPConn, receive chan int) {
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

func HandleSensorData(conn *net.UDPConn, receive chan int) {
	count := 0
	for {
		message := make([]byte, MAXLINE)
		size, addr, err := conn.ReadFrom(message)
		if err != nil {
			return
		}
		var sensorData pb.SensorData
		err = proto.Unmarshal(message[:size], &sensorData)
		handleError(err)

		handleError(db.Create(&sensorData).Error)

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
