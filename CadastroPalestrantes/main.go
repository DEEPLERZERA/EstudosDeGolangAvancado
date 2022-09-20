package main

import (
	"ProjetosDeGolang/CadastroPalestrantes/database"
	"ProjetosDeGolang/CadastroPalestrantes/routes"
)


func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}