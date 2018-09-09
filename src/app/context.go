package app

import (
	"config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// App has router and db instances
type Context struct {
	Router *mux.Router
	//DB     *gorm.DB
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
	Handler func(*Context, *http.ResponseWriter, *http.Request)
}

type WSEndpoint struct {
	URL     string
	Handler func(*Context, *http.ResponseWriter, *http.Request)
}

// Initialize initializes the app with predefined configuration
func (ctx *Context) Initialize(config *config.Config) {
	//dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
	//    config.DB.Username,
	//    config.DB.Password,
	//    config.DB.Name,
	//    config.DB.Charset)

	//db, err := gorm.Open(config.DB.Dialect, dbURI)
	//if err != nil {
	//    log.Fatal("Could not connect database")
	//}

	//a.DB = model.DBMigrate(db)
	ctx.Router = mux.NewRouter()
}

func (ctx *Context) SetRestAPI(routes *[]*RestEndpoint) {
	for _, r := range *routes {
		ctx.Router.HandleFunc(r.URL, func(wr http.ResponseWriter, req *http.Request) {
			r.Handler(ctx, &wr, req)
		}).Methods(string(r.Method))
	}
}

func (ctx *Context) SetSocketAPI(routes *[]*WSEndpoint) {
	for _, r := range *routes {
		ctx.Router.HandleFunc(r.URL, func(wr http.ResponseWriter, req *http.Request) {
			r.Handler(ctx, &wr, req)
		})
	}
}

func (ctx *Context) Run(port string) {
	log.SetFlags(0)
	log.Fatal(http.ListenAndServe(port, ctx.Router))
}
