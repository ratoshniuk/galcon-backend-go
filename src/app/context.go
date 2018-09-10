package app

import (
	"config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"galcon-backend-go/wsctx"
	"os"
)

// App has router and db instances
type Context struct {
	Router *mux.Router
	Hub *wsctx.Hub
}

type METHOD string

const (
	GET    METHOD = "GET"
	POST   METHOD = "POST"
	PUT    METHOD = "PUT"
	DELETE METHOD = "DELETE"
)

type RestEndpoint struct {
	URL     string
	Method  METHOD
	Handler func(*Context, http.ResponseWriter, *http.Request)
}

type WSEndpoint struct {
	URL     string
	Handler func(*Context, http.ResponseWriter, *http.Request)
}

func (ep *WSEndpoint) AsHandler(ctx *Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ep.Handler(ctx, w, r)
	})
}

// Initialize initializes the app with predefined configuration
func (ctx *Context) Initialize(config *config.Config) {
	ctx.Hub = wsctx.NewHub()
	go ctx.Hub.Run()

	ctx.Router = mux.NewRouter()
}

func (ctx *Context) SetRestAPI(routes *[]*RestEndpoint) {
	for _, r := range *routes {
		ctx.Router.HandleFunc(r.URL, func(wr http.ResponseWriter, req *http.Request) {
			r.Handler(ctx, wr, req)
		}).Methods(string(r.Method))
	}
}

func (ctx *Context) SetSocketAPI(routes *[]*WSEndpoint) {
	for _, r := range *routes {
		ctx.Router.Handle(r.URL, r.AsHandler(ctx))
	}
}

func (ctx *Context) Run(port string) {
	log.SetFlags(0)
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), ctx.Router))
}
