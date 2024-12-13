package main

import (
    "github.com/gorilla/websocket"
)

type client struct {
    // Web socket for this client
    socket *websocket.Conn
    // channel on which messages are sent
    send chan []byte
    // the room client is chatting in
    room *room
}

func (c *client) read() {
    // Infinitely read from the websocket connection and forward it to the room
    for {
        if _, msg, err := c.socket.ReadMessage(); err == nil {
            // upon successful reading; read msg to the room
            c.room.forward <- msg
        } else {
            c.socket.Close()
            break
        }
    }
}

func (c *client) write() {
    // Infinitely check if there is a message from the room 
    for msg := range c.send {
        if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
            c.socket.Close()
            break
        }
    }
}
