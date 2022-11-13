package main

import (
	"log"

	"01.alem.school/git/Nurzhas/groupie-tracker/internal/delivery"
)

func main() {
	route := new(delivery.Route)
	server := new(delivery.Server)
	if err := server.Run("8080", route.InitRoute()); err != nil {
		log.Fatal(err)
	}
}
