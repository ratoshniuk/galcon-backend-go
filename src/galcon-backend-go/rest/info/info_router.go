package info

import (
	"app"
	rest "galcon-backend-go/rest/common"
)

var Router = []*app.RestEndpoint{
	rest.GET("/info", GetInfoHandler),
}
