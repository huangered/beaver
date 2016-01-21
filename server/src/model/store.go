package model

import (
	"log"
)

// Store save volumes
type Store struct {
	Volumes map[int32]*Volume
}

func NewStore() (s *Store) {
	s = &Store{}
	s.Volumes = make(map[int32]*Volume)
	s.init()
	return
}

func (s *Store) init() {
	log.Println("Read store conf")
	log.Println("Read volume")

}

func (s *Store) AddVolume() {

}

func (s *Store) DelVolume() {

}

func (s *Store) Close() {

}

func (s *Store) Clean() {

}
