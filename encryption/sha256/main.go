package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)
func main(){
  f,err:=os.Open("simple-file.txt")

  if err!=nil{
    log.Fatal(err)
  }
  defer f.Close()


// New returns a new hash.Hash computing the SHA256 checksum. The Hash
// also implements encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler to marshal and unmarshal the internal
// state of the hash.
// func New() hash.Hash {
// 	d := new(digest)
// 	d.Reset()
// 	return d
// }


  // hash
  h:=sha256.New()

  _,err=io.Copy(h,f)

  if err!=nil{
    log.Fatalln("couldn't io.copy: ",err)
  }

  xb:=h.Sum(nil)

  fmt.Printf("%x\n",xb)

}
