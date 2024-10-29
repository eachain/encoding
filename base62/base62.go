package base62

import "github.com/eachain/encoding/basen"

const base62Std = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var base62 = basen.NewEncoding(base62Std)

func EncodeToString(b []byte) string {
	return base62.EncodeToString(b)
}

func DecodeString(b string) ([]byte, error) {
	return base62.DecodeString(b)
}
