package server

import (
	"app/db"
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
	ssCount = 0
	sdCount = 0
)

func Start() {
	ssc := make(chan int)
	sdc := make(chan int)
	quit := make(chan bool)
	handleInterrupt(quit)
	go ListenUDP(ssc, sdc, quit)
	go StartApi()
	<-quit
	close(ssc)
	close(sdc)
	fmt.Println("\nSensorState COUNT =", ssCount)
	fmt.Println("SensorData COUNT =", sdCount)
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

func ListenUDP(ssc chan int, sdc chan int, quit chan bool) {
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
	ssConn.Close()
	sdConn.Close()
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
		internal.FatalError(err)

		internal.FatalError(db.DB.Create(&db.MainComputer{MainComputer: sensorState.MainComputer, CreatedAt: time.Now()}).Error)
		internal.FatalError(db.DB.Create(&db.BrakeManager{BrakeManager: sensorState.BrakeManager, CreatedAt: time.Now()}).Error)

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
		internal.FatalError(err)

		internal.FatalError(db.DB.Create(&db.SensorData{SensorData: &sensorData, CreatedAt: time.Now()}).Error)

		receive <- 1
		count++
		log.Printf("[%s] : COUNT = %d\n", addr, count)
	}
}

func StartApi() {
	fiberApi := fiber.New(fiber.Config{
		Prefork: false,
	})
	fiberApi.Use(cors.New(), etag.New(), logger.New())
	api.RegisterRoutes(fiberApi.Group("/api"))
	log.Fatal(fiberApi.Listen(":8000"))
}
