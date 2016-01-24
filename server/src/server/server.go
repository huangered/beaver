package main

import (
	"log"
	"model"
)

func main() {
	log.Println("Start the server...")

	var (
		s *model.Store
		c *model.Config
	)
	c = model.NewConfig("server.conf")
	s = model.NewStore(c)
	StartApi("localhost:8000", s)
	StartSignal()
}
