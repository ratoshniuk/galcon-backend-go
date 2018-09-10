package info

import (
	"app"
	"bytes"
	"galcon-backend-go/rest/common"
	"math"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var startTime = time.Now()

func GetInfoHandler(ctx *app.Context, rw http.ResponseWriter, req *http.Request) {

	common.RespondJSON(rw, http.StatusOK, &Info{
		Version:         "1.0.0-SNAPSHOT",
		Revision:        fetchLastCommit(),
		Owner:           "ratoshniuk",
		UptimeInSeconds: math.Trunc(time.Since(startTime).Seconds()),
	})
}

func fetchLastCommit() string {
	cmd := exec.Command("bash", "-c", "git log | head -1")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		return "UNKNOWN"
	}
	out := string(cmdOutput.Bytes())
	version := strings.Replace(out, "commit ", "", 1)
	return version
}
