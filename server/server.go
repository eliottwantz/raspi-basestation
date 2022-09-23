package server

import (
	"app/db"
	"app/db/sqlc"
	"app/internal"
	"app/pb"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/protobuf/proto"
)

const (
	MAXLINE = 1024
	WORKERS = 100
)

var (
	ssCount  = 0
	ssWCount = 0
	sdCount  = 0
	sdWCount = 0

	ssc   = make(chan *pb.SensorState)
	sdc   = make(chan *pb.SensorData)
	Wsss  = make(chan *db.SensorState)
	Wssd  = make(chan *sqlc.SensorData)
	sscc  = make(chan int)
	sswcc = make(chan int)
	sdcc  = make(chan int)
	sdwcc = make(chan int)
	quit  = make(chan bool)
)

func Start() {
	handleInterrupt(quit)
	go listenUDP()
	go writeToDb()
	<-quit
	close(ssc)
	close(sdc)
	close(sscc)
	close(sswcc)
	close(sdcc)
	close(sdwcc)
	close(quit)
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

func listenUDP() {
	ssAddr, err := net.ResolveUDPAddr("udp", ":8080")
	internal.FatalError(err)
	ssConn, err := net.ListenUDP("udp", ssAddr)
	internal.FatalError(err)

	sdAddr, err := net.ResolveUDPAddr("udp", ":8081")
	internal.FatalError(err)
	sdConn, err := net.ListenUDP("udp", sdAddr)
	internal.FatalError(err)

	fmt.Println("server listening \n", ssConn.LocalAddr(), sdConn.LocalAddr())

	for i := 0; i < WORKERS; i++ {
		go handleSensorState(ssConn)
		go handleSensorData(sdConn)
	}
	go func() {
		for {
			select {
			case c := <-sscc:
				ssCount += c
			case c := <-sswcc:
				ssWCount += c
				fmt.Println("SensorStateWritten COUNT =", ssCount)
			case c := <-sdcc:
				sdCount += c
			case c := <-sdwcc:
				sdWCount += c
				fmt.Println("SensorDataWritten COUNT =", sdWCount)
			}
		}
	}()
	<-quit
	ssConn.Close()
	sdConn.Close()
}

func handleSensorState(conn *net.UDPConn) {
	count := 0
	for {
		message := make([]byte, MAXLINE)
		size, _, err := conn.ReadFrom(message)
		if err != nil {
			return
		}
		var sensorState pb.SensorState
		err = proto.Unmarshal(message[:size], &sensorState)
		internal.FatalError(err)

		ssc <- &sensorState

		sscc <- 1
		count++
		// log.Printf("[%s] : COUNT = %d\n", addr, count)
	}
}

func handleSensorData(conn *net.UDPConn) {
	count := 0
	for {
		message := make([]byte, MAXLINE)
		size, _, err := conn.ReadFrom(message)
		if err != nil {
			return
		}
		var sensorData pb.SensorData
		err = proto.Unmarshal(message[:size], &sensorData)
		internal.FatalError(err)

		sdc <- &sensorData

		sdcc <- 1
		count++
		// log.Printf("[%s] : COUNT = %d\n", addr, count)
	}
}

func writeToDb() {
	for {
		select {
		case ss := <-ssc:
			mc, err := db.Queries.CreateMainComputer(db.Ctx, sqlc.CreateMainComputerParams{
				CreatedAt: time.Now(),
				State:     int64(ss.MainComputer.State),
			})
			internal.NotFatalError(err)
			bm, err := db.Queries.CreateBrakeManager(db.Ctx, sqlc.CreateBrakeManagerParams{
				CreatedAt:                            time.Now(),
				State:                                int64(ss.BrakeManager.State),
				HydrolicPressureLoss:                 int64(ss.BrakeManager.HydrolicPressureLoss),
				CriticalPodAccelerationMesureTimeout: int64(ss.BrakeManager.CriticalPodAccelerationMesureTimeout),
				CriticalPodDecelerationInstructionTimeout: int64(ss.BrakeManager.CriticalEmergencyBrakesWithoutDeceleration),
				VerinBlocked: int64(ss.BrakeManager.VerinBlocked),
				EmergencyValveOpenWithoutHydrolicPressorDiminution: int64(ss.BrakeManager.EmergencyValveOpenWithoutHydrolicPressorDiminution),
				CriticalEmergencyBrakesWithoutDeceleration:         int64(ss.BrakeManager.CriticalEmergencyBrakesWithoutDeceleration),
				MesuredDistanceLessThanDesired:                     int64(ss.BrakeManager.MesuredDistanceLessThanDesired),
				MesuredDistanceGreaterAsDesired:                    int64(ss.BrakeManager.MesuredDistanceGreaterAsDesired),
			})
			internal.NotFatalError(err)

			sswcc <- 1
			Wsss <- &db.SensorState{
				BrakeManager: &bm,
				MainComputer: &mc,
			}
		case sd := <-sdc:
			dbsd, err := db.Queries.CreateSensorData(db.Ctx, sqlc.CreateSensorDataParams{
				CreatedAt: time.Now(),
				SensorID:  int64(sd.SensorId),
				Value:     float64(sd.Value),
			})
			internal.NotFatalError(err)
			sdwcc <- 1
			Wssd <- &dbsd
		}
	}
}
