package main

import (
	"encoding/hex"
	"fmt"
	"password_cracker/internal"
	"password_cracker/internal/cracker"
)

func main() {
	internal.Md5Func()
	internal.MyMd5Func()
	hashBytes, _ := hex.DecodeString("b3af409bb8423187c75e6c7f5b683908")
	var hashArray [16]byte
	copy(hashArray[:], hashBytes)
	fmt.Printf("%v\n", cracker.Crack(hashArray,
		cracker.Config{Min: 3, Max: 3, SpecialChars: false, Digits: false, Capitals: false}))
}
