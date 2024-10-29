# encoding

encoding提供了base58、base62，以及自定义编码。

## 示例

```go
package main

import (
	stdbase64 "encoding/base64"
	"fmt"

	"github.com/eachain/encoding/base58"
	"github.com/eachain/encoding/base62"
	"github.com/eachain/encoding/basen"
)

func main() {
	p := []byte("hello world!")
	base36 := basen.NewEncoding(`0123456789abcdefghijklmnopqrstuvwxyz`)
	base64 := basen.NewEncoding(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`)

	fmt.Println("base36:", base36.EncodeToString(p))
	fmt.Println("base58:", base58.EncodeToString(p))
	fmt.Println("base62:", base62.EncodeToString(p))
	fmt.Println("base64:", base64.EncodeToString(p))
	fmt.Println("std base64:", stdbase64.RawStdEncoding.EncodeToString(p))

	if b, err := base36.DecodeString(base36.EncodeToString(p)); err == nil {
		fmt.Println(string(b))
	}
	if b, err := base58.DecodeString(base58.EncodeToString(p)); err == nil {
		fmt.Println(string(b))
	}
	if b, err := base62.DecodeString(base62.EncodeToString(p)); err == nil {
		fmt.Println(string(b))
	}
	if b, err := base64.DecodeString(base64.EncodeToString(p)); err == nil {
		fmt.Println(string(b))
	}
	if b, err := stdbase64.RawStdEncoding.DecodeString(stdbase64.RawStdEncoding.EncodeToString(p)); err == nil {
		fmt.Println(string(b))
	}
}
```

**注意：** 上述示例中，当且仅当被编码的参数长度是12的倍数时，base64编码结果才和标准库`encoding/base64.RawStdEncoding.EncodeToString`相同。
