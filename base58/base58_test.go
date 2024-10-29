package base58

import (
	"bytes"
	"testing"
)

func TestBase58(t *testing.T) {
	b1 := []byte("hello base58")
	s := EncodeToString(b1)
	b2, err := DecodeString(s)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	if !bytes.Equal(b1, b2) {
		t.Fatalf("b1: %s, b2: %s", b1, b2)
	}
}
