package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

// bearer token is very simple. It is like you got a token and you put that into the Header
// HMAC is s cryptographic signing algorithm.

func signMessage(msg []byte)([]byte, error){
  // hmac.New take a new instance of another cryptographic function
  // the most common one to use with hmac is sha512
  // the second parameter that it takes is the key (slice of bytes) - it's the key that you generate yourself and use for all messages with this HMAC 
  // the key size should mach the size of the hash that you are using with it. 
  hash:=hmac.New(sha512.New,GenerateKey())
  // this hash is a io writer, so we can write a message 

  _,err := hash.Write(msg)
  if err!=nil{
      return nil,fmt.Errorf("error while hashing message: %w",err)
  }



	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.

  signature:=hash.Sum(nil) // use nil if you are writing the hashbusing hash.Write
  return signature,nil 


}

// so you got the message as a signature, you send it to the user and the user sends it back to you. In this way, you have both the original message and signature (sent by the user). So, this function compares both
func checkSig(msg,sig []byte)(bool , error){
  // you first sign the message that you got
    newSign,err:=signMessage(msg)
    if err!=nil{
      return false,err
    }
    same:= hmac.Equal(newSign,sig)
    return same,nil 
}



func GenerateKey()[]byte{
  var key=[]byte{}

  for i:=1;i<=63;i++{
    key=append(key, byte(i))
  }

  return key 
}
