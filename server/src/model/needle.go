package model

import (
	"fmt"
	"hash/crc32"
)

const (
	// Needle size
	magicSize        = 4
	cookieSize       = 8
	keySize          = 8
	alternateKeySize = 4
	flagsSize        = 1
	dataChecksumSize = 4
	HeaderSize       = magicSize + cookieSize + keySize + alternateKeySize + flagsSize
	FooterSize       = magicSize + dataChecksumSize

	// flags
	FlagOk  = byte(0)
	FlagDel = byte(1)
)

var (
	// crc32 checksun table , gorouting safe
	crc32Table = crc32.MakeTable(crc32.Koopman)
	// magic number
	headerMagic = []byte{0x12, 0x34, 0x56, 0x78}
	footerMagic = []byte{0x87, 0x65, 0x43, 0x21}
)

// Needle is a photo data stored in disk.
type Needle struct {
	Header       []byte
	Cookie       int64
	Key          int64
	AlternateKey int32
	Flags        byte
	Size         int64 // Data size
	Data         []byte
	Footer       []byte
	DataChecksum uint32
	Padding      []byte
}

func (n *Needle) Init(key int64, cookie int64, data []byte) {
	var dataSize = int64(len(data))
	n.Header = headerMagic
	n.Key = key
	n.Cookie = cookie
	n.Size = dataSize
	n.Data = data
	n.Footer = footerMagic
	n.DataChecksum = crc32.Update(0, crc32Table, data)
}

func (n *Needle) String() string {
	return fmt.Sprintf(`
-----------------------
Header:          %v
Cookie:          %d
Key:             %d
AlternativeKey:  %d
Flags:           %d
Size:            %d
Data:            %v
Footer:          %v
Checksum:        %d
Padding:         %v
-----------------------
`, n.Header, n.Cookie, n.Key, n.AlternateKey, n.Flags, n.Size, n.Data, n.Footer, n.DataChecksum, n.Padding)
}
