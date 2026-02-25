package bootstrap

import (
	"context"
	"log"
	"sync"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/config"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/delivery/http/handler"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/infrastructure/messaging"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/seed"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Engine   *gin.Engine
	DB       *gorm.DB
	Modules  *Modules
	RabbitMQ *messaging.RabbitMQ
}

type Modules struct {
	AuthHandler         *handler.UserHandler
	DoaHandler          *handler.DoaHandler
	JadwalSholatHandler *handler.JadwalSholatHandler
	CityHandler         *handler.CityHandler
}

// Init Module
func InitModules(db *gorm.DB, redis *redis.Client, rabbitmq messaging.EmailPublisherHandler) *Modules {
	// Auth Module
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo, rabbitmq)
	userHandler := handler.NewUserHandler(userUseCase)

	// DoaModule
	doaRepo := repository.NewDoaRepository(db, redis)
	doaUseCase := usecase.NewDoaUsecase(doaRepo)
	doaHandler := handler.NewDoaUsecase(doaUseCase)

	// JadwalSholat Module
	jadwalSholatRepo := repository.NewJadwalSholatRepository(db)
	jadwalSholatUseCase := usecase.NewJadwalSholatuseCase(jadwalSholatRepo)
	jadwalSholatHandler := handler.NewJadwalSholatHandler(jadwalSholatUseCase)

	// City Module
	cityRepo := repository.NewCityRepository(db, redis)
	cityUseCase := usecase.NewCityUsecase(cityRepo)
	cityHandler := handler.NewCityHandler(cityUseCase)

	return &Modules{
		AuthHandler:         userHandler,
		DoaHandler:          doaHandler,
		JadwalSholatHandler: jadwalSholatHandler,
		CityHandler:         cityHandler,
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

	// cron := cron.NewCronJob(db)
	// cron.Start()

	rabbitmq, err := messaging.NewRabbitMQ("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	publisher := messaging.NewEmailPublisher(rabbitmq.Channel)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := seedDoa(db, redis, ctx); err != nil {
			log.Printf("Seed Doa error: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := seed.SeedCities(db); err != nil {
			log.Printf("Seed Cities error: %v", err)
		}
	}()

	wg.Wait()

	engine := gin.Default()

	return &App{
		Engine:   engine,
		DB:       db,
		Modules:  InitModules(db, redis, publisher),
		RabbitMQ: rabbitmq,
	}
}
