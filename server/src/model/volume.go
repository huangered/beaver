package model

import ()

type Volume struct {
	Id int32 `json:"id"`
}

func NewVolume(id int32, blockFile string, indexFile string) (v *Volume, err error) {
	v = &Volume{}
	v.Id = id
	return
}
