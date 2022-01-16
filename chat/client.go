package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	// send is a channel where message is sent to
	send chan []byte
	// room is a chatroom where this client participates in
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
