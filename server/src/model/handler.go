package model

type Handler interface {
	Handle(data []byte) error
}
