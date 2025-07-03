package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/net"
	"go.bug.st/serial"
)

func main() {
	portName := flag.String("port", "", "Bluetooth COM port (COM5, /dev/rfcomm0, …)")
	flag.Parse()
	if *portName == "" {
		log.Fatalln("Usage: traffic_bt -port=COM5")
	}

	mode := &serial.Mode{BaudRate: 115200}
	port, err := serial.Open(*portName, mode)
	if err != nil {
		log.Fatalf("Opening %s failed: %v\n", *portName, err)
	}
	defer port.Close()

	prev, _ := net.IOCounters(false)
	prevTime := time.Now()

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		nowStats, err := net.IOCounters(false)
		if err != nil || len(nowStats) == 0 {
			continue
		}
		elapsed := time.Since(prevTime).Seconds()

		up   := float64(nowStats[0].BytesSent-prev[0].BytesSent) / 1024 / elapsed
		down := float64(nowStats[0].BytesRecv-prev[0].BytesRecv) / 1024 / elapsed

		fmt.Fprintf(port, "%.1f,%.1f\n", up, down)

		prev = nowStats
		prevTime = time.Now()
	}
}
