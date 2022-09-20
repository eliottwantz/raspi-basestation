package main

import (
	"app/database"
	"context"
	"fmt"
)

const (
	MAXLINE = 1024
)

func main() {
	db, err := database.Open()
	if err != nil {
		panic(err)
	}
	mc, err := db.GetMainComputerStates(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(mc)
	// conn, err := net.ListenUDP("udp", &net.UDPAddr{
	// 	Port: 8080,
	// 	IP:   net.ParseIP("127.0.0.1"),
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// defer conn.Close()
	// fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	// for {
	// 	message := make([]byte, MAXLINE)
	// 	rlen, remote, err := conn.ReadFromUDP(message[:])
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	var sensorState pb.SensorState
	// 	err = proto.Unmarshal(message[:rlen], &sensorState)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Printf("[%s]: %v\n", remote, &sensorState)
	// }
}
