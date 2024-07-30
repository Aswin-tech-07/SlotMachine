package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"slot-machine-api/models"
	"slot-machine-api/repository"
)

const (
	RTPThreshold   = 0.975
	PlayCost       = 10
	SmallWinAmount = 20
	BigWinAmount   = 100
	SmallWinChance = 0.1
	BigWinChance   = 0.01
)

type GameService struct {
	playerRepo  *repository.PlayerRepository
	redisClient *redis.Client
}

func NewGameService(playerRepo *repository.PlayerRepository, redisClient *redis.Client) *GameService {
	return &GameService{playerRepo, redisClient}
}

func (s *GameService) Play(playerID primitive.ObjectID) (*models.PlayResult, error) {
	player, err := s.playerRepo.GetPlayerByID(playerID)
	if err != nil {
		return nil, err
	}

	if player.Status != "active" {
		return nil, fiber.NewError(fiber.StatusForbidden, "player is not active")
	}

	if player.Credits < PlayCost {
		return nil, fiber.NewError(fiber.StatusForbidden, "not enough credits")
	}

	player.Credits -= PlayCost

	rand.Seed(time.Now().UnixNano())
	roll := rand.Float64()
	var result models.PlayResult

	if roll < BigWinChance {
		player.Credits += BigWinAmount
		result = models.PlayResult{Result: "big_win", Payout: BigWinAmount}
		s.redisClient.Incr(context.Background(), "big_wins")
	} else if roll < SmallWinChance {
		player.Credits += SmallWinAmount
		result = models.PlayResult{Result: "small_win", Payout: SmallWinAmount}
		s.redisClient.Incr(context.Background(), "small_wins")
	} else {
		result = models.PlayResult{Result: "lose", Payout: 0}
		s.redisClient.Incr(context.Background(), "loses")
	}

	err = s.playerRepo.UpdatePlayerCredits(playerID, player.Credits)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
