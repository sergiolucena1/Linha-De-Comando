package app

import "github.com/urfave/cli"

//Gerar vai retornar a aplicação de linha de comando, pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"
	return app
}
