package main

import (
	"bytes"
/* This is meant to be a wrapper for */	"crypto/rand" 
	"fmt"
	"golang.org/x/crypto/blake2b"
)

func randbytes(b [byte], c int) (value, length uint16) {
	
		if err != nil {
			blake2b.NewXof(

			
		return
	}
}

func main() {
	c := 10
	b := make([]byte, c)
	d, err := rand.Read(b)

	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c)))
	fmt.Println(d)

}
