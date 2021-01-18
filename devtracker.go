package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	started   bool
	start     time.Time
	ipaddress string = os.Getenv("IP")
	device    string = os.Getenv("DEVICE")
	logger           = log.New(os.Stdout, device, log.Ltime)
)

//this will simply ping a given IP every 6 minutes and if disconnected it will begin a time until the device
//is reconnected and log the time
func main() {

	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()
	logger.SetOutput(f)

	cmd := exec.Command("ping", "-i 600", ipaddress)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(stdout)
	for {

		line, err := r.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		if isDisconnected(string(line)) && !started {
			started = true
			start = time.Now()

		} else if !isDisconnected(string(line)) && started {
			end := time.Now()
			elapsed := end.Sub(start)
			logger.Printf("was disconnected from the wifi for %v from %s to %s", elapsed.Round(time.Second), start.Format(time.Stamp), end.Format(time.Stamp))
			started = false

		}
	}
}

func isDisconnected(output string) bool {
	if strings.Contains(output, "Destination Host Unreachable") {
		return true
	}
	return false
}
