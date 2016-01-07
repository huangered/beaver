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
	//start()
	StartApi("localhost:8000")
	StartSignal()
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
	for {
		header := make([]byte, model.TagLen+model.LenLen)
		_, err := c.Read(header)
		if err != nil {
			log.Fatal("Read conn error: " + err.Error())
		}

		message := model.Message{}
		message.Decode(header)

		header = make([]byte, message.Len)
		_, err = c.Read(header)
		if err != nil {
			log.Fatal("Read conn error: " + err.Error())
		}
		message.Value = header

		switch message.Tag {
		case model.UserTag:
			handler := controller.UserController{}
			handler.Handle(message.Value)
		case model.FileTag:
			handler := controller.FileController{}
			handler.Handle(message.Value)
		}
	}
}
