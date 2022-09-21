package server

import (
	"app/db"
	"app/internal"
	"app/pb"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

const (
	MAXLINE = 1024
	WORKERS = 100
)

var (
	ssCount = 0
	sdCount = 0
)

func ListenUDP(ssc chan int, sdc chan int, quit chan bool) {
	ssAddr, err := net.ResolveUDPAddr("udp", ":8080")
	internal.HandleError(err)
	ssConn, err := net.ListenUDP("udp", ssAddr)
	internal.HandleError(err)

	sdAddr, err := net.ResolveUDPAddr("udp", ":8081")
	internal.HandleError(err)
	sdConn, err := net.ListenUDP("udp", sdAddr)
	internal.HandleError(err)

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
		internal.HandleError(err)

		internal.HandleError(db.DB.Create(&db.MainComputer{MainComputer: sensorState.MainComputer, CreatedAt: time.Now()}).Error)
		internal.HandleError(db.DB.Create(&db.BrakeManager{BrakeManager: sensorState.BrakeManager, CreatedAt: time.Now()}).Error)

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
		internal.HandleError(err)

		internal.HandleError(db.DB.Create(&db.SensorData{SensorData: &sensorData, CreatedAt: time.Now()}).Error)

		receive <- 1
		count++
		log.Printf("[%s] : COUNT = %d\n", addr, count)
	}
}
