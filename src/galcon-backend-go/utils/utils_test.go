package utils

import (
    "fmt"
    "os/exec"
    "testing"
)

func TestUtils(t *testing.T) {
    res, err := exec.Command("echo $FOO").Output()
    if err != nil {

    }

    fmt.Printf("ssss")
    fmt.Printf(string(res))
}
