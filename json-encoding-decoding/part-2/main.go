package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string `json:"first"`
}

func main() {
	var bs []byte
	xp := []person{}

	err := json.Unmarshal(bs, &xp)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Print("\n\n", string(bs))

}
