package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando, pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	app.Commands = []cli.Command{ //comandos e seus valores
		{
			Name:  "ip",
			Usage: "Busca IPS de endereços na Internet",
			Flags: []cli.Flag{
				cli.StringFlag{ // flag e seu valor
					Name:  "host",
					Value: "amazon.com",
				},
			},
			Action: buscarIps,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host) // procura o host passado e retorna o IP
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
