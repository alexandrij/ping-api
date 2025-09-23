package main

import (
	"log"
	"ping-api/cmd/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
