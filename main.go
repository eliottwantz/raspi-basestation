package main

import (
	"app/pb"
	"fmt"
	"net"

	"github.com/google/uuid"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	// for {
	// 	message := make([]byte, 1024*8)
	// 	rlen, remote, err := conn.ReadFromUDP(message[:])
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	data := strings.TrimSpace(string(message[:rlen]))
	// 	fmt.Printf("received: %s from %s\n", data, remote)
	// }
	sdata := &pb.SensorData{Id: uuid.NewString(), Value: 32}
	fmt.Println(sdata)
	fmt.Println("sdata.ProtoMessage():", sdata.ProtoReflect())
}
