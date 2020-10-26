package main

import(
	"os/exec"
	"os"
	"log"
	"bytes"
	"fmt"
	"strings"
	)
func main (){
	//WHY THE FUCK IS THIS NOT WORKING
	var (
	buf bytes.Buffer
	output bytes.Buffer
	ipaddress string = os.Getenv("IP")
	logger = log.New(&buf, "Dante's Phone: ", log.Ltime)
	)
	
	cmd := exec.Command("ping", "-i 5", ipaddress)
	cmd.Stdout = &output
	err := cmd.Start()

	if err != nil {
		fmt.Println(err)
	}
	
	if strings.Contains(output.String(), "Destination Host Unreachable"){
		logger.Print("Disconnected")
		fmt.Print("connected")
	}
	logger.Print("Connected")
	fmt.Print("connected")
cmd.Wait()
}
