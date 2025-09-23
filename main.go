package main

import (
	"github.com/Alexandrij/ping-api/cmd/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
