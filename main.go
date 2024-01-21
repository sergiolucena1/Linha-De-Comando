package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")
	aplicação := app.Gerar()
	if err := aplicação.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
