package main

import (
	"fmt"
	"log"

	"github.com/Safwanseban/server-go/models"
)

func main() {

	server := models.Newserver(":3000")
	go func() {
		for v := range server.Msgch {

			fmt.Println(v.From, string(v.Payload))

		}
	}()

	log.Fatal(server.Start())

}
