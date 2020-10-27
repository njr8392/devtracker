package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var (
	//	buf       bytes.Buffer
		ipaddress string = os.Getenv("IP")
		//switch logger from buffer to os.Stdout
		logger           = log.New(os.Stdout, "Dante's Phone: ", log.Ltime)
	)

	cmd := exec.Command("ping", "-i 25", ipaddress)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(stdout)
	for {
		line, _, _ := r.ReadLine()
		if strings.Contains(string(line), "Destination Host Unreachable") {
			logger.Printf("Disconnected")
		}
		logger.Printf("Connected")
	}
}
