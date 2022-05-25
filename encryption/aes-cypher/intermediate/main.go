package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	msg := " this is fun"
	password := "ilovedogs"

	// GenerateFromPassword returns the bcrypt hash of the password at the given
	// cost.If the cost is less than the MinCost , then the cost will ve set to
	// DefaultCost. Use CompareHashAndPassword to compare the retured hashed password
	// with its cleartext version.
	// func GenerateFromPassword (password []byte, cost int)([]byte, error)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	CheckError("GenerateFromPassword", err)
	// hash=hash[:16]

	rslt, err := encodeDecode(hash, msg)
	CheckError("encode", err)
  
	fmt.Println("before base64", string(rslt))


  rslt2,err:=encodeDecode(hash,string(rslt))

  CheckError("decode",err)
  fmt.Println(rslt2)
}


func CheckError( funcName string,err error) {
	if err != nil {
		log.Fatal(fmt.Errorf("Error in %s : %w", funcName, err))
	}
}


func encodeDecode(key []byte, input string)([]byte,error){
	cipherBlock,err:=aes.NewCipher(key)

	CheckError("encodeDecode",err)

	//making memory for the cipher or initialization vector
	iv:=make([]byte,aes.BlockSize)

	
// 	// NewCTR returns a Stream which encrypts/decrypts using the given Block in
// 	// counter mode. The length of iv must be the same as the Block's block size.
// 	//func NewCTR(block Block, iv []byte) Stream {
//


	s:=cipher.NewCTR(cipherBlock ,iv)
	buff:=&bytes.Buffer{}

	sw:=cipher.StreamWriter{
		S:s,
		W:buff,
	}
	_,err=sw.Write([]byte(input))

	CheckError("StreamWriter",err)

	return buff.Bytes(),nil 
}
