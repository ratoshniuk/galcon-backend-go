package common

import (
	"app"
	"github.com/hokaccha/go-prettyjson"
	"net/http"
)

// respondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := prettyjson.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

// Get wraps the router for GET method
func GET(path string, f func(ctx *app.Context, w *http.ResponseWriter, r *http.Request)) *app.RestEndpoint {
	return requestBuilder(app.GET, path, f)
}

func POST(path string, f func(ctx *app.Context, w *http.ResponseWriter, r *http.Request)) *app.RestEndpoint {
	return requestBuilder(app.POST, path, f)
}

func PUT(path string, f func(ctx *app.Context, w *http.ResponseWriter, r *http.Request)) *app.RestEndpoint {
	return requestBuilder(app.PUT, path, f)
}

func DELETE(path string, f func(ctx *app.Context, w *http.ResponseWriter, r *http.Request)) *app.RestEndpoint {
	return requestBuilder(app.DELETE, path, f)
}

func requestBuilder(method app.METHOD, path string, f func(ctx *app.Context, w *http.ResponseWriter, r *http.Request)) *app.RestEndpoint {
	return &app.RestEndpoint{
		URL:     path,
		Handler: f,
		Method:  method,
	}
}
