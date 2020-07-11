package usecase

import (
	"main/internal/models"
	"main/internal/repository"
)

type UseCase interface {
	Register(user *models.User) (*models.User, error)
	DeleteAccount(userID int) error
}

type useCase struct {
	repository repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{
		repository: repo,
	}
}