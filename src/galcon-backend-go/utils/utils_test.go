package utils

import (
    "fmt"
    "log"
    "os/exec"
    "testing"
)

func TestUtils(t *testing.T) {
    cmd := exec.Command("echo $GOPATH")
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v", stdout)
}
