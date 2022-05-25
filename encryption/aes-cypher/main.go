package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cast"
)

func main() {

	msg := "this is totlaly fun gettijng hand on and learning it from the ground up."
	// 	encoded_msg := encode(msg)
	// 	fmt.Println("\n\nEncoded message: ", encoded_msg)
	//
	// 	decoded_msg, err := decode(encoded_msg)
	// 	if err != nil {
	// 		log.Println(fmt.Errorf("Error in decoding value: %w", err))
	// 	}
	// 	fmt.Println("\nDecoded message: ", decoded_msg)

	password := "myPass"
}
func enDecode(key []byte) ([]byte, error) {
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't newCipher: %w", err)
	}

}
