package info

import (
	"app"
	"galcon-backend-go/rest/common"
	"net/http"
)

func GetInfoHandler(ctx *app.Context, rw *http.ResponseWriter, req *http.Request) {
	common.RespondJSON(*rw, http.StatusOK, &Info{
		Version:  "1.0.0-SNAPSHOT",
		Branch:   "develop",
		Revision: "10rfasra12arra3",
		Owner:    "ratoshniuk",
	})
}
