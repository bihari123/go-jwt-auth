package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cast"
)

func main() {

	msg := "this is totlaly fun gettijng hand on and learning it from the ground up."
	encoded_msg := encode(msg)
	fmt.Println("\n\nEncoded message: ", encoded_msg)

	decoded_msg, err := decode(encoded_msg)
	if err != nil {
		log.Println(fmt.Errorf("Error in decoding value: %w", err))
	}
	fmt.Println("\nDecoded message: ", decoded_msg)
}
func encode(msg string) string {

	return base64.URLEncoding.EncodeToString([]byte(msg))

}

func decode(encoded_msg string) (string, error) {
	s, err := base64.URLEncoding.DecodeString(encoded_msg)

	if err != nil {
		return "", fmt.Errorf("couldn't decode string: %w", err)
	}

	return cast.ToString(s), err
}
