package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"slot-machine-api/models"
	"slot-machine-api/repository"
)

func RegisterPlayerHandlers(app *fiber.App, repo *repository.PlayerRepository) {
	app.Post("/players", createPlayer(repo))
	app.Get("/players/:id", getPlayer(repo))
	app.Put("/players/:id/suspend", suspendPlayer(repo))
	app.Put("/players/:id/activate", activatePlayer(repo))
}

func createPlayer(repo *repository.PlayerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		player := new(models.Player)
		if err := c.BodyParser(player); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}

		err := repo.CreatePlayer(player)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot insert player"})
		}

		return c.Status(fiber.StatusCreated).JSON(player)
	}
}

func getPlayer(repo *repository.PlayerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		playerID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid player ID"})
		}

		player, err := repo.GetPlayerByID(playerID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "player not found"})
		}

		return c.JSON(player)
	}
}

func suspendPlayer(repo *repository.PlayerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		playerID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid player ID"})
		}

		err = repo.SuspendPlayer(playerID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update player status"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
func activatePlayer(repo *repository.PlayerRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		playerID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid player ID"})
		}

		err = repo.ActivatePlayer(playerID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update player status"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
