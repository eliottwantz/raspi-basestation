package main

import (
	"app/database"
	"app/database/sqlc"
	"app/pb"
	"context"
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
		fmt.Printf("[%s]: %v\n", remote, &sensorState)
		mc, err := db.CreateMainComputerState(context.Background(), pb.MainComputer_States_name[int32(sensorState.MainComputer.State)])
		if err != nil {
			panic(err)
		}
		bm, err := db.CreateBrakeManager(context.Background(), sqlc.CreateBrakeManagerParams{
			State:                                int64(sensorState.BrakeManager.State),
			HydrolicPressureLoss:                 sensorState.BrakeManager.HydrolicPressureLoss,
			CriticalPodAccelerationMesureTimeout: sensorState.BrakeManager.CriticalPodAccelerationMesureTimeout,
			CriticalPodDecelerationInstructionTimeout: sensorState.BrakeManager.CriticalEmergencyBrakesWithoutDeceleration,
			VerinBlocked: sensorState.BrakeManager.VerinBlocked,
			EmergencyValveOpenWithoutHydrolicPressorDiminution: sensorState.BrakeManager.EmergencyValveOpenWithoutHydrolicPressorDiminution,
			CriticalEmergencyBrakesWithoutDeceleration:         sensorState.BrakeManager.CriticalEmergencyBrakesWithoutDeceleration,
			MesuredDistanceLessThanDesired:                     sensorState.BrakeManager.MesuredDistanceLessThanDesired,
			MesuredDistanceGreaterAsDesired:                    sensorState.BrakeManager.MesuredDistanceGreaterAsDesired,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("FROM DB:", mc)
		fmt.Println("FROM DB:", bm)
	}
}
