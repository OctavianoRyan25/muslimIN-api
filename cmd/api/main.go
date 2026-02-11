package main

import (
	"log"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/bootstrap"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/routes"
)

func main() {
	app := bootstrap.NewApp()

	routes.RegisterRoutes(app.Engine, app.Modules.AuthHandler)

	log.Printf("Server running on port %s", "8080")
	app.Engine.Run(":8080")
}
