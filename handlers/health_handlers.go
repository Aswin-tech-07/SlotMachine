package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterHealthHandlers(app *fiber.App, mongoClient *mongo.Client, redisClient *redis.Client) {
	app.Get("/health", healthCheck(mongoClient, redisClient))
	app.Get("/liveness", livenessCheck)
	app.Get("/readiness", readinessCheck(mongoClient, redisClient))
}

func healthCheck(mongoClient *mongo.Client, redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := mongoClient.Ping(context.Background(), nil); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "MongoDB not available"})
		}
		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Redis not available"})
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

func livenessCheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func readinessCheck(mongoClient *mongo.Client, redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := mongoClient.Ping(context.Background(), nil); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "MongoDB not available"})
		}
		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Redis not available"})
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
