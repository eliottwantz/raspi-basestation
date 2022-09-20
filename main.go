package main

import (
	"app/database"
)

const (
	MAXLINE = 1024
)

func main() {
	_, err := database.Open()
	if err != nil {
		panic(err)
	}

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
