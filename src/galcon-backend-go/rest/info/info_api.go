package info

import (
	"app"
	"bufio"
	"fmt"
	"galcon-backend-go/rest/common"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

var startTime = time.Now()

func GetInfoHandler(ctx *app.GlobalContext, rw http.ResponseWriter, req *http.Request) {

	common.RespondJSON(rw, http.StatusOK, &Info{
		Version:         fetchVersion(),
		Revision:        fetchLastCommit(),
		Owner:           "ratoshniuk",
		UptimeInSeconds: math.Trunc(time.Since(startTime).Seconds()),
	})
}

func fetchLastCommit() string {
	f, err := os.Open("revision.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	if e != nil {
		return "UNKNOWN"
	}
	return strings.Replace(s, "commit ", "", 1)
}

func fetchVersion() string {
	f, err := os.Open("version.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	if e != nil {
		return "UNKNOWN"
	}
	return s
}

func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
}
