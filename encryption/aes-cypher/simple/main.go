package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// cipher key that you need to encrypt and decrypt the message

	var key = "thisis32bitlongpassphraseimusing"

	// plain text

	var pt = " This is a secret"

	var c = EncryptAES([]byte(key), pt)

	fmt.Println("the encrypted message: ", c)

	var d =DecryptAES([]byte(key),c)
	fmt.Println("The decrypted message: ",d)

}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher

	c, err := aes.NewCipher(key)

	CheckError(err, "EncryptAES")

	// allocate space for ciphered message 
	out := make([]byte, len(plaintext))

	// Encrypt
	c.Encrypt(out, []byte(plaintext))

	// return hex string
	return hex.EncodeToString(out)

}


func DecryptAES(key []byte, ct string) string{

	ciphertext,_:=hex.DecodeString(ct)

	c,err:=aes.NewCipher(key)
  CheckError(err, "DecryptAES")

  pt:= make([]byte,len(ciphertext))
  c.Decrypt(pt,ciphertext)

  s:=string(pt[:])
  
  return s 


}


func CheckError(err error, funcName string) {
	if err != nil {
		log.Fatal(fmt.Errorf("Error in %s : %w", funcName, err))
	}
}
