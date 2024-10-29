package base58

import "github.com/eachain/encoding/basen"

const base58Std = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var base58 = basen.NewEncoding(base58Std)

func EncodeToString(b []byte) string {
	return base58.EncodeToString(b)
}

func DecodeString(b string) ([]byte, error) {
	return base58.DecodeString(b)
}
