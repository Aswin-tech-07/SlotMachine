package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron/v3"

	"slot-machine-api/config"
	"slot-machine-api/handlers"
	"slot-machine-api/jobs"
	"slot-machine-api/repository"
	"slot-machine-api/services"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	cfg := config.LoadConfig()

	mongoClient := repository.InitMongoDB(cfg.MongoURI)
	redisClient := repository.InitRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(logger.New())

	playerRepo := repository.NewPlayerRepository(mongoClient)
	gameService := services.NewGameService(playerRepo, redisClient)

	handlers.RegisterPlayerHandlers(app, playerRepo)
	handlers.RegisterGameHandlers(app, gameService)
	handlers.RegisterHealthHandlers(app, mongoClient, redisClient)

	c := cron.New()
	c.AddFunc("@startup", func() { jobs.EnsureIndexes(mongoClient) })
	c.Start()

	log.Fatal(app.Listen(":3000"))
}
