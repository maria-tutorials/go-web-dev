package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	s64Standard := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
	fmt.Println(s64Standard)

	bs, _ := base64.StdEncoding.DecodeString(s64)

	fmt.Println("decoded", string(bs))
}
