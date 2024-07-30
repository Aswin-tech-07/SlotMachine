package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"slot-machine-api/services"
)

func RegisterGameHandlers(app *fiber.App, service *services.GameService) {
	app.Post("/play", play(service))
}

func play(service *services.GameService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type PlayRequest struct {
			PlayerID string `json:"player_id"`
		}
		var req PlayRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}

		playerID, err := primitive.ObjectIDFromHex(req.PlayerID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid player ID"})
		}

		result, err := service.Play(playerID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(result)
	}
}
