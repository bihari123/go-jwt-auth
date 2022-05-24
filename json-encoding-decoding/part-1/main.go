package main

import (
	"encoding/json"
	"fmt"
	"log"
)


type person struct{
  First string 
}

func main(){

  p1:= person{
    First: "Jenny",
  }

  p2:= person{
    First: "James",
  }

  xp:=[]person{p1,p2}

  bs,err:=json.Marshal(xp)

  if err!= nil{
    log.Fatal(err)

  }

  fmt.Print("\n\n",string(bs))

}
