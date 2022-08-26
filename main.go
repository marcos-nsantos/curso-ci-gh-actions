package main

import (
	"github.com/marcos-nsantos/curso-ci-gh-actions/database"
	"github.com/marcos-nsantos/curso-ci-gh-actions/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
