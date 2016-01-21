package model

type Volume struct {
	Id int32 `json:"id"`
}

func NewVolume(id int32) (v *Volume, err int) {
	v = &Volume{}
	v.Id = id
	return
}
