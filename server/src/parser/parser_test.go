package parser

import (
	"testing"
)

func TestInt64(t *testing.T) {
	var num int64 = 1234567890
	data := make([]byte, 8)
	PutInt64(num, data)
	if Int64(data) != num {
		t.Errorf("int64 method is failed, expected %v, but it is %v", num, Int64(data))
		t.FailNow()
	}
}

func TestInt32(t *testing.T) {
	var num int32 = 123456
	data := make([]byte, 4)
	PutInt32(num, data)
	if Int32(data) != num {
		t.Errorf("int32 method is failed, expected %v, but it is %v", num, Int32(data))
		t.FailNow()
	}
}

func TestUint32(t *testing.T) {
	var num uint32 = 123456
	data := make([]byte, 4)
	PutUint32(num, data)
	if Uint32(data) != num {
		t.Errorf("uint32 method is failed, expected %v, but it is %v", num, Uint32(data))
		t.FailNow()
	}
}
