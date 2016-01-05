package core

const (
	Login  = 1
	Logout = 2
)

//    0    1     2    3    4    5      6
// [tag][tag],[len][len][len][len],[data]
type Message struct {
	Tag   uint8
	Len   uint32
	Value []byte
}

func (t *Message) encode() []byte {
	data := make([]byte, t.Len)
	// tag
	data[0] = 0
	data[1] = t.Tag
	// len
	data[2] = byte(t.Len)
	data[3] = byte(t.Len >> 8)
	data[4] = byte(t.Len >> 16)
	data[5] = byte(t.Len >> 24)
	// data
	append(data, Value)
	return data
}

func (t *Message) decode(data []byte) {
	// tag
	t.Tag = data[1]
	// len
	t.Len = uint32(data[2]) + uint32(data[3])<<8 + uint32(data[4])<<16 + uint32(data[5])<<24
	// data
	t.Value = data[6:]
}
