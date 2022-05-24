package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string `json:"First"`
}

func encode(w http.ResponseWriter, r *http.Request) {

	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}

  // the diffference between the Marshal and encoder is that encoder encoded the data into json on the go whereas marshal takes the full data and then encodes.
  // Encoder is helpful in writing rest apis as it is memory efficient
  err:=json.NewEncoder(w).Encode(xp)

  if err!=nil{
  	log.Println(fmt.Errorf("encoded bad data: %w",err))
  }

}

func decode(w http.ResponseWriter, r *http.Request) {
	xp := []person{}
	err:=json.NewDecoder(r.Body).Decode(&xp)

	if err!=nil{
		log.Println(fmt.Errorf("Error decoding the data: %w",err))
	} 

	

}

func main(){
	http.HandleFunc("/encode",encode)
	http.HandleFunc("/decode",decode)
	http.ListenAndServe(":8080",nil)
}
