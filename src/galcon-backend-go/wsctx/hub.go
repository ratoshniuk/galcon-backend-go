package wsctx

// clients.
type Hub struct {
    // Registered clients.
    clients map[*Client]bool

    // Inbound messages from the clients.
    broadcast chan []byte

    // Register requests from the clients.
    Register chan *Client

    // Unregister requests from clients.
    Unregister chan *Client
}

func NewHub() *Hub {
    return &Hub{
        broadcast:  make(chan []byte),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        clients:    make(map[*Client]bool),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.clients[client] = true
        case client := <-h.Unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.Send)
            }
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.Send <- message:
                default:
                    close(client.Send)
                    delete(h.clients, client)
                }
            }
        }
    }
}
