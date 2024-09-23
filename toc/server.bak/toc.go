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
  fmt.Printf("Started sleeping\n")
  time.Sleep(60 * time.Minute)
  fmt.Printf("Stopped sleeping\n")
}

func Exec(service string) {
  cmd := exec.Command("/usr/bin/docker", "compose", "exec", service, "cat", "/etc/hostname")
  cmd.Dir = "/opt/toc"
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
