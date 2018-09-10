package chat

import (
    "app"
    "log"
    "net/http"
    "galcon-backend-go/wsctx"
)

func ServeHome(ctx *app.Context, w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL)

    if r.URL.Path != "/" {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    http.ServeFile(w, r, "src/github.com/gorilla/websocket/examples/chat/home.html")
}

func ServeWs(ctx *app.Context, w http.ResponseWriter, r *http.Request) {
    conn, err := wsctx.Upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    client := &wsctx.Client{Huv: ctx.Hub, Conn: conn, Send: make(chan []byte, 256)}
    client.Huv.Register <- client

    // Allow collection of memory referenced by the caller by doing all work in
    // new goroutines.
    go client.WritePump()
    go client.ReadPump()
}
