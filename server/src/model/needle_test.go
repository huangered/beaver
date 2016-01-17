package model

import (
	"bytes"
	"testing"
)

func TestInit(t *testing.T) {
	var key int64 = 1234567890
	var cookie int64 = 987654321
	var data []byte = []byte("abcdefghijklmn")
	n := &Needle{}
	n.Init(key, cookie, data)
	if n.Key != key {
		t.Errorf("key is not equal to %d", key)
		t.FailNow()
	}
	if n.Cookie != cookie {
		t.Errorf("cookie is not equal to %d", cookie)
		t.FailNow()
	}
	if bytes.Compare(n.Data, data) != 0 {
		t.Errorf("data is not equal to %v,\n it is %v", data, n.Data)
		t.FailNow()
	}
}
