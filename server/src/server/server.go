package main

import (
	"fmt"
	"log"
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

}
