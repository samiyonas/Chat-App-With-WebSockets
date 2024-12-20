package main

import (
    "net/http"
    "github.com/gorilla/websocket"
    "log"
)

type room struct {
    // channel holds incoming messages
    forward chan []byte
    // channel for clients wishing to join
    join chan *client
    // channel for clients wishing to leave
    leave chan *client
    // clients in the room
    client map[*client]bool
}

func (r *room) run() {
    for {
        select {
        case client := <-r.join:
            r.client[client] = true
        case client := <-r.leave:
            delete(r.client, client)
            close(client.send)
        case msg := <-r.forward:
            for client := range r.client {
                select {
                case client.send <- msg:
                    // send the message
                default:
                    delete(r.client, client)
                    close(client.send)
                }
            }
        }
    }
}

// Upgrader type to upgrade HTTP to Websocket
var upgrader = websocket.Upgrader{
    // Reads in a chunk of 1024 maximum
    ReadBufferSize: 1024,
    // Writes in a chunk of 1024 maximum
    WriteBufferSize: 1024,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // on success it upgrades and returns the websocket connection representative
    socket, err := upgrader.Upgrade(w, req, nil)
    if err != nil {
        log.Fatal(err)
        return
    }

    // create a client object so we can deal with our connection with the client
    client := &client{
        socket: socket,
        send: make(chan []byte, 1024),
        room: r,
    }

    // add client to the room
    r.join <- client

    // remove the client if any error happend while writing or reading to/from the room
    defer func() { r.leave <- client }()

    // Write to websocket
    go client.write()
    // Read from websocket
    client.read()
}
