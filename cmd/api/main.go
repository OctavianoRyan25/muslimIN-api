package main

import (
	"context"
	"log"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/bootstrap"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/routes"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app := bootstrap.NewApp(ctx)

	routes.RegisterRoutes(app.Engine, app.Modules.AuthHandler, app.Modules.DoaHandler)

	log.Printf("Server running on port %s", "8080")
	app.Engine.Run(":8080")
}
