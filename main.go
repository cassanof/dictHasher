package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	plaintext := "hello"

	// MD5 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	fmt.Println(plaintext)
	fmt.Printf("%x\n", hash(plaintext))
}

func hash(plaintext string) string {
	h := md5.New()

	h.Write([]byte(plaintext))

	ciphertext := h.Sum(nil)
	return string(ciphertext)
}
