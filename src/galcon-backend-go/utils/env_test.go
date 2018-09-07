package utils

import (
    "fmt"
    "log"
    "os/exec"
    "testing"
)

func TestEnv(t *testing.T) {

    out, err := exec.Command("echo $HEROKU_API").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The date is %s\n", out)
}
