package service

import (
	"Geeklanet/models"
	"Geeklanet/repository"
	"math/rand"
)

type recommendService struct {
	r repository.RecommendRepository
}



func (s recommendService) GetRecommendWeight(userID string) func(post models.Post) float64 {
	return func(post models.Post) float64 {
		return float64(rand.Intn(100) / 100.0)
	}
}