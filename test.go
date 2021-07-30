package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "my password"
	pwd := sha256.Sum256([]byte(password))
	fmt.Printf("%x",pwd)
}
