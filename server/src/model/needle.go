package model

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"parser"
)

const (
	// Needle size
	magicSize        = 4
	cookieSize       = 8
	keySize          = 8
	alternateKeySize = 4
	flagsSize        = 1
	sizeSize         = 4
	dataChecksumSize = 4
	HeaderSize       = magicSize + cookieSize + keySize + alternateKeySize + flagsSize + sizeSize
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

func (n *Needle) ParseHeader(buf []byte) (err int) {
	var bn int
	n.Header = buf[:magicSize]
	if !bytes.Equal(n.Header, headerMagic) {
		err = 1
		return
	}
	bn += magicSize
	// cookie
	n.Cookie = parser.Int64(buf[bn:])
	bn += cookieSize
	// key
	n.Key = parser.Int64(buf[bn:])
	bn += keySize
	// alter key
	n.AlternateKey = parser.Int32(buf[bn:])
	bn += alternateKeySize
	n.Flags = buf[bn]
	if n.Flags != FlagOk && n.Flags != FlagDel {
		err = 2
		return
	}
	bn += flagsSize
	// size
	n.Size = parser.Int64(buf[bn:])
	return
}

// the buf contains data and footer
func (n *Needle) ParseFooter(buf []byte) (err int) {
	var (
		bn       int64
		checksum uint32
	)
	n.Data = buf[:n.Size]
	bn += n.Size
	n.Footer = buf[bn : bn+magicSize]
	if !bytes.Equal(n.Footer, footerMagic) {
		err = 1
		return
	}
	bn += magicSize
	checksum = crc32.Update(0, crc32Table, n.Data)
	n.DataChecksum = parser.Uint32(buf[bn : bn+dataChecksumSize])
	if n.DataChecksum != checksum {
		err = 1
		return
	}
	bn += dataChecksumSize
	n.Padding = buf[bn : bn+1]
	return
}

func (n *Needle) WriterHeader(buf []byte) {
	var bn int
	// header
	// magic
	bn += copy(buf, n.Header)
	// cookie
	parser.PutInt64(buf[bn:], n.Cookie)
	// key
	bn += cookieSize
	parser.PutInt64(buf[bn:], n.Key)
	// alterkey
	bn += keySize
	parser.PutInt32(buf[bn:], n.AlternateKey)
	buf[bn] = FlagOk
	bn += flagsSize
	// size
	parser.PutInt64(buf[bn:], n.Size)
}

func (n *Needle) WriterFooter(buf []byte) {

}

func (n *Needle) Write(buf []byte) {

}
