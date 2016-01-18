package parser

func PutInt64(num int64, buf []byte) {
	for i := 0; i < 8; i++ {
		offset := uint(56 - i*8)
		buf[i] = byte(num >> offset)
	}
}

func Int64(buf []byte) int64 {
	var num int64 = 0
	for i := 0; i < 8; i++ {
		offset := uint(i * 8)
		num |= int64(buf[7-i]) << offset
	}
	return num
}

func PutInt32(num int32, buf []byte) {
	buf[0] = byte(num >> 24)
	buf[1] = byte(num >> 16)
	buf[2] = byte(num >> 8)
	buf[3] = byte(num)
}

func Int32(buf []byte) int32 {
	num := int32(buf[3]) | int32(buf[2])<<8 | int32(buf[1])<<16 | int32(buf[0])<<24
	return num
}

func PutUint32(num uint32, buf []byte) {
	buf[0] = byte(num >> 24)
	buf[1] = byte(num >> 16)
	buf[2] = byte(num >> 8)
	buf[3] = byte(num)
}

func Uint32(buf []byte) uint32 {
	num := uint32(buf[3]) | uint32(buf[2])<<8 | uint32(buf[1])<<16 | uint32(buf[0])<<24
	return num
}
