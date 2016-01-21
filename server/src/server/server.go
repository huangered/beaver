package main

import (
	"log"
	"model"
)

func main() {
	log.Println("Start the server...")

	var (
		s *model.Store
	)
	s = model.NewStore()
	StartApi("localhost:8000", s)
	StartSignal()
}
