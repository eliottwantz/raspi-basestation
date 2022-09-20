package main

import (
	"app/database"
	"app/pb"
	"fmt"
	"log"
	"net"

	// "github.com/bvinc/go-sqlite-lite/sqlite3"
	"google.golang.org/protobuf/proto"
)

const (
	MAXLINE = 1024
)

func main() {
	// _, err := database.Open()
	// if err != nil {
	// 	panic(err)
	// }
	db := database.Open()

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	count := 0

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
		// _, err = db.CreateMainComputer(context.Background(), int64(sensorState.MainComputer.State))
		// PanicError(err)
		// _, err = db.CreateBrakeManager(context.Background(), sqlc.CreateBrakeManagerParams{
		// 	State:                                int64(sensorState.BrakeManager.State),
		// 	HydrolicPressureLoss:                 sensorState.BrakeManager.HydrolicPressureLoss,
		// 	CriticalPodAccelerationMesureTimeout: sensorState.BrakeManager.CriticalPodAccelerationMesureTimeout,
		// 	CriticalPodDecelerationInstructionTimeout: sensorState.BrakeManager.CriticalEmergencyBrakesWithoutDeceleration,
		// 	VerinBlocked: sensorState.BrakeManager.VerinBlocked,
		// 	EmergencyValveOpenWithoutHydrolicPressorDiminution: sensorState.BrakeManager.EmergencyValveOpenWithoutHydrolicPressorDiminution,
		// 	CriticalEmergencyBrakesWithoutDeceleration:         sensorState.BrakeManager.CriticalEmergencyBrakesWithoutDeceleration,
		// 	MesuredDistanceLessThanDesired:                     sensorState.BrakeManager.MesuredDistanceLessThanDesired,
		// 	MesuredDistanceGreaterAsDesired:                    sensorState.BrakeManager.MesuredDistanceGreaterAsDesired,
		// })
		// PanicError(err)

		// count, err := db.GetMainComputerCount(context.Background())
		// PanicError(err)
		count++
		log.Printf("[%s] : Count = %d\n", remote, count)
	}
}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
