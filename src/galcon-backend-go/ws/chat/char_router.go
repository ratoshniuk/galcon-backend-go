package chat

import (
    "app"
)

var Routes = []*app.WSEndpoint{
    {
        URL:     "/",
        Handler: ServeHome,
    },
    {
        URL: "/ws",
        Handler: ServeWs,
    },
}

