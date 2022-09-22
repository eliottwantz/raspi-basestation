package server

import (
	"app/db"
	"app/db/sqlc"
	"app/internal"
	"app/pb"
	"app/server/api"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	sscc  = make(chan int)
	sswcc = make(chan int)
	sdcc  = make(chan int)
	sdwcc = make(chan int)
	quit  = make(chan bool)
)

func Start() {
	handleInterrupt(quit)
	go listenUDP()
	go startApi()
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
		for c := range sscc {
			ssCount += c
		}
	}()
	go func() {
		for c := range sswcc {
			ssWCount += c
			fmt.Println("SensorStateWritten COUNT =", ssCount)
		}
	}()
	go func() {
		for c := range sdcc {
			sdCount += c
		}
	}()
	go func() {
		for c := range sdwcc {
			sdWCount += c
			fmt.Println("SensorDataWritten COUNT =", sdWCount)
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

func startApi() {
	fiberApi := fiber.New(fiber.Config{
		Prefork: false,
	})
	fiberApi.Use(cors.New(), etag.New(), logger.New())
	api.RegisterRoutes(fiberApi.Group("/api"))
	log.Fatal(fiberApi.Listen(":8000"))
}

func writeToDb() {
	for {
		select {
		case ss := <-ssc:
			_, err := db.Queries.CreateMainComputer(db.Ctx, sqlc.CreateMainComputerParams{
				CreatedAt: time.Now(),
				State:     int64(ss.MainComputer.State),
			})
			internal.NotFatalError(err)
			_, err = db.Queries.CreateBrakeManager(db.Ctx, sqlc.CreateBrakeManagerParams{
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
		case sd := <-sdc:
			_, err := db.Queries.CreateSensorData(db.Ctx, sqlc.CreateSensorDataParams{
				CreatedAt: time.Now(),
				SensorID:  int64(sd.SensorId),
				Value:     float64(sd.Value),
			})
			internal.NotFatalError(err)
			sdwcc <- 1
		}
	}
}
