package bootstrap

import (
	"context"
	"log"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/config"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/handler"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Engine  *gin.Engine
	DB      *gorm.DB
	Modules *Modules
}

type Modules struct {
	AuthHandler *handler.UserHandler
	DoaHandler  *handler.DoaHandler
}

// Init Module
func InitModules(db *gorm.DB, redis *redis.Client) *Modules {
	// Auth Module
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	// DoaModule
	doaRepo := repository.NewDoaRepository(db, redis)
	doaUseCase := usecase.NewDoaUsecase(doaRepo)
	doaHandler := handler.NewDoaUsecase(doaUseCase)

	return &Modules{
		AuthHandler: userHandler,
		DoaHandler:  doaHandler,
	}
}

func NewApp(ctx context.Context) *App {
	cfg := config.LoadConfig()
	cfgRedis := config.LoadRedis()

	db := InitDatabase(cfg)
	redis, err := InitRedis(cfgRedis)

	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	if err := seedDoa(db, redis, ctx); err != nil {
		log.Printf("Seed error: %v", err)
	}

	engine := gin.Default()

	return &App{
		Engine:  engine,
		DB:      db,
		Modules: InitModules(db, redis),
	}
}
