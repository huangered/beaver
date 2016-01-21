package index

import (
	"fmt"
	"parser"
)

const (
	keySize      = 8
	alterKeySize = 8
	flatsSize    = 1
	offsetSize   = 8
	sizeSize     = 8
)

type needle struct {
	Key      int64 //64-key
	AlterKey int64 // 32-alternate key
	Flags    byte  // currently unused
	Offset   int64 //needle offset in the store
	Size     int64 // needle data size
}

func (n *needle) parse(buf []byte) {
	n.Key = parser.Int64(buf)
	n.AlterKey = parser.Int64(buf[keySize:])
	n.Flags = buf[keySize+alterKeySize]
	n.Offset = parser.Int64(buf[keySize+alterKeySize+flatsSize:])
	n.Size = parser.Int64(buf[keySize+alterKeySize+flatsSize+offsetSize:])
}

func (n *needle) write() (buf []byte) {
	return
}

func (n *needle) string() string {
	return fmt.Sprintf(`
-----------------------
Key:             %d
AlternativeKey:  %d
Flags:           %d
Offset:          %d
Size:            %d
-----------------------
`, n.Key, n.AlterKey, n.Flags, n.Offset, n.Size)
}

type Index struct {
	Superblock superblock
	needles    []needle
}

func (index *Index) Parse(buf []byte) {

}

func (index *Index) Write() (buf []byte) {
	return
}
