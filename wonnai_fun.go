package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sDec, _ := b64.StdEncoding.DecodeString(data)
	fmt.Println(string(sDec))
	fmt.Println()
	fmt.Println( Reverse(string(sDec)))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
