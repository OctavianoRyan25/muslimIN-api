package bootstrap

import (
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/config"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/handler"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Engine  *gin.Engine
	DB      *gorm.DB
	Modules *Modules
}

type Modules struct {
	AuthHandler *handler.UserHandler
}

func InitModules(db *gorm.DB) *Modules {
	// Init Module
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	return &Modules{
		AuthHandler: userHandler,
	}
}

func NewApp() *App {
	cfg := config.LoadConfig()

	db := InitDatabase(cfg)

	engine := gin.Default()

	return &App{
		Engine:  engine,
		DB:      db,
		Modules: InitModules(db),
	}
}
