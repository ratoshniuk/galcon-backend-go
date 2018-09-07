package utils

import (
    "fmt"
    "os/exec"
    "testing"
)

func TestSysEnvs(t *testing.T) {
    res := exec.Command("echo $HEROKU_API_SECRET")
    fmt.Printf("%+v", res)
}
