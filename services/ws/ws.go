package services

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type client struct {
	ws      *websocket.Conn
	context echo.Context
	message chan *Message
}

// Message .
type Message struct {
	UserID  string
	Content interface{}
}

type hub struct {
	register    chan *client
	sendMessage chan *Message
	clients     map[string][]*client
}

var currentHub *hub

// Init .
func init() {
	currentHub = &hub{
		register:    make(chan *client),
		sendMessage: make(chan *Message),
	}
	go runHub()
}

func runHub() {
	select {
	case register := <-currentHub.register:
		fmt.Println(register)
	case message := <-currentHub.sendMessage:
		if clients, ok := currentHub.clients[message.UserID]; ok {
			for _, client := range clients {
				client.message <- message
			}
		}
	}
}

func writeData(client *client) {
	for {
		select {
		case message := <-client.message:
			b, err := json.Marshal(message.Content)
			if err == nil {
				client.ws.Write(b)
			}
		}
	}
}

func readData(client *client) {
	ws := client.ws
	logger := client.context.Logger()
	defer ws.Close()
	for {
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			logger.Error(err)
			break
		}
	}
}

// SendMessage .
func SendMessage(message *Message) {
	currentHub.sendMessage <- message
}

// Upgrade .
func Upgrade(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		client := &client{
			ws:      ws,
			context: c,
		}
		defer readData(client)
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
