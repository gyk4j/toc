package main

import (
  "fmt"
  "log"
  "os/exec"
  "strings"
  "time"
)

func main() {
  Exec("db")
  Exec("file")
  Exec("mail")
  Exec("web")
  time.Sleep(60 * time.Minute)
}

func Exec(service string) {
  cmd := exec.Command("docker", "compose", "exec", service, "cat", "/etc/hostname")
  var out strings.Builder
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  // Shell output has a new line at the end.
  // So excluding our own \n new line in Printf.
  fmt.Printf("%s: %s", service, out.String())
}
