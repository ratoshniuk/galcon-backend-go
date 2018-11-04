package rest

import (
	command "galcon-backend-go/rest/command"
	"galcon-backend-go/rest/info"
)

var Routes = append(info.Router, command.Router...)
