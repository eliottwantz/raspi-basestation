package main

import (
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
		var unmarshalled pb.SensorPacketData
		err = proto.Unmarshal(message[:rlen], &unmarshalled)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[%s]: %v", remote, &unmarshalled)
	}
}
