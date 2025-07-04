package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/net"
	"go.bug.st/serial"
)

var (
	start = time.Now()
)

func main() {
	com := flag.String("port", "", "Outgoing COM port (e.g. COM7 or /dev/rfcomm0)")
	baud := flag.Int("baud", 115200, "Baud rate (must match ESP32 sketch)")
	retry := flag.Duration("retry", 3*time.Second, "How long to wait before retrying a lost port")
	flag.Parse()

	if *com == "" {
		log.Fatalln("Missing -port flag. Example: traffic_bt -port=COM7")
	}

	log.Printf("Starting sender on %s @ %d baud …", *com, *baud)

	for {
		if err := runSession(*com, *baud); err != nil {
			log.Printf("▸ Session ended: %v", err)
		}
		log.Printf("Re‑attempting connection in %v …", *retry)
		time.Sleep(*retry)
	}
}

func runSession(portName string, baud int) error {
	mode := &serial.Mode{BaudRate: baud}
	port, err := serial.Open(portName, mode)
	if err != nil {
		return fmt.Errorf("opening %s: %w", portName, err)
	}
	defer port.Close()
	log.Printf("Port %s opened", portName)

	prev, _ := net.IOCounters(false)
	prevTime := time.Now()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		nowStats, err := net.IOCounters(false)
		if err != nil || len(nowStats) == 0 {
			log.Printf("warning: could not read net stats: %v", err)
			continue
		}

		elapsed := time.Since(prevTime).Seconds()
		upKBs := float64(nowStats[0].BytesSent-prev[0].BytesSent) / 1024 / elapsed
		dnKBs := float64(nowStats[0].BytesRecv-prev[0].BytesRecv) / 1024 / elapsed
		runtime := int(time.Since(start).Seconds())

		line := fmt.Sprintf("%.1f,%.1f,%d\n", upKBs, dnKBs, runtime)

		if _, err := port.Write([]byte(line)); err != nil {
			return fmt.Errorf("write failed: %w", err)
		}

		prev = nowStats
		prevTime = time.Now()
	}
	return nil
}
