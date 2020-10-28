package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	var (
		//	buf       bytes.Buffer
		ipaddress string = os.Getenv("IP")
		//switch logger from buffer to os.Stdout
		logger = log.New(os.Stdout, "Dante's Phone: ", log.Ltime)
	)

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
		if err != nil{
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		if strings.Contains(string(line), "Destination Host Unreachable") {
			logger.Printf("Disconnected")
		} else {
				logger.Printf("Connected")
		}
	}
}
