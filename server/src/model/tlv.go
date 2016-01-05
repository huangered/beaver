package model

const (
	TagLen = 2
	LenLen = 4
)

const (
	UserTag = 1
)

//    0    1     2    3    4    5      6
// [tag][tag],[len][len][len][len],[data]
type Message struct {
	Tag   uint8
	Len   uint32
	Value []byte
}

func (t *Message) Encode() []byte {
	data := make([]byte, t.Len+TagLen+LenLen)
	// tag
	data[0] = 0
	data[1] = t.Tag
	// len
	data[2] = byte(t.Len)
	data[3] = byte(t.Len >> 8)
	data[4] = byte(t.Len >> 16)
	data[5] = byte(t.Len >> 24)
	// data
	copy(data[6:], t.Value)
	return data
}

func (t *Message) Decode(data []byte) {
	// tag
	t.Tag = data[1]
	// len
	t.Len = uint32(data[2]) + uint32(data[3])<<8 + uint32(data[4])<<16 + uint32(data[5])<<24
}
