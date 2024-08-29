package main

import (
  "fmt"
  "log"
  "os/exec"
  "strings"
  "time"
)

func main() {
  Exec("toc-db-1")
  Exec("toc-file-1")
  Exec("toc-mail-1")
  Exec("toc-web-1")
  time.Sleep(2 * time.Minute)
}

func Exec(container string) {
  cmd := exec.Command("docker", "exec", container, "cat", "/etc/hostname")
  var out strings.Builder
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  // Shell output has a new line at the end.
  // So excluding our own \n new line in Printf.
  fmt.Printf("%s: %s", container, out.String())
}
