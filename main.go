package main

import (
	"github.com/L-Bellei/api-gin-rest/database"
	"github.com/L-Bellei/api-gin-rest/routes"
)

func main() {
	database.ConectarBancoDeDados()

	routes.HandleRequests()

}
