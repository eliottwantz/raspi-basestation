package main

import (
	"app/database"
	"app/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

const (
	MAXLINE = 1024
)

func main() {
	db, err := database.Open()
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, MAXLINE)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}
		var sensorState pb.SensorState
		err = proto.Unmarshal(message[:rlen], &sensorState)
		if err != nil {
			log.Fatalln(err)
		}

		PanicError(db.Create(&sensorState.MainComputer).Error)
		PanicError(db.Create(&sensorState.BrakeManager).Error)

		var mc database.MainComputer
		var bm database.BrakeManager
		PanicError(db.First(&mc).Error)
		PanicError(db.First(&bm).Error)
		// fmt.Println(&mc)
		// fmt.Println(&bm)
		log.Printf("[%s] : Count = %d\n", remote, database.GetCount(db))
	}
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
