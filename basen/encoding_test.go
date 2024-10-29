package basen

import (
	"bytes"
	"testing"
)

func TestEncoding(t *testing.T) {
	base36 := NewEncoding(`0123456789abcdefghijklmnopqrstuvwxyz`)

	b1 := []byte("hello base36")
	s := base36.EncodeToString(b1)
	b2, err := base36.DecodeString(s)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	if !bytes.Equal(b1, b2) {
		t.Fatalf("b1: %s, b2: %s", b1, b2)
	}
}
