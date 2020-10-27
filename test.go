package main

import (
  "bufio"
  "os"
  "os/exec"
  "fmt"
  "log"
)

func main() {
  cmd := exec.Command("ping", "127.0.0.1")
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }
  cmd.Start()

  buf := bufio.NewReader(stdout) // Notice that this is not in a loop
  num := 1
  for {
    line, _, _ := buf.ReadLine()
    if num > 3 {
      os.Exit(0)
    }
    num += 1
    fmt.Println(string(line))
  }
}
