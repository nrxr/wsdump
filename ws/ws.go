package ws

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

type Client struct {
	addr string
	ws   *websocket.Conn
}

func New(addr string) (*Client, error) {
	return &Client{addr: addr}, nil
}

func (c *Client) connect() {
	conn, _, err := websocket.DefaultDialer.Dial(c.addr, nil)
	if err != nil {
		panic(err)
	}
	c.ws = conn
}

func (c *Client) readMessages() {
	for {
		_, m, err := c.ws.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Printf("server « %s", string(m))
	}
}

func (c *Client) receiveCommand() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Printf("local  » %s", text)
	c.ws.WriteMessage(websocket.TextMessage, []byte(text))
}

func (c *Client) Run() {
	c.connect()
	fmt.Println("Press Ctrl+C to quit. Write something and hit enter to send a message.")

	go c.readMessages()

	for {
		c.receiveCommand()
	}
}
