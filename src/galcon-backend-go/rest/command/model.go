package info

type CommandRequest struct {
	Command   string
	Arguments interface{}
}

type PlayerReadyResponse struct {
	Success   bool
	Arguments interface{}
}

type PlayerKickedResponse struct {
	Success   bool
	Arguments interface{}
}
