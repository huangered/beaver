package main

import (
	"controller"
	"fmt"
	"log"
	"model"
	"net"
)

func main() {
	fmt.Println("Start the server...")
	start()
}

// start the tcp reciever.
func start() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	header := make([]byte, model.TagLen+model.LenLen)
	_, err := c.Read(header)
	if err != nil {
		log.Println("Exit")
		log.Fatal(err)
	}

	message := model.Message{}
	message.Decode(header)

	header = make([]byte, message.Len)
	_, err = c.Read(header)
	if err != nil {
		log.Fatal(err)
	}
	message.Value = header

	switch message.Tag {
	case model.UserTag:
		handler := controller.UserController{}
		handler.Handle(message.Value)
	}
}
