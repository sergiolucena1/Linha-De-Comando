package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicação := app.Gerar()
	if err := aplicação.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
