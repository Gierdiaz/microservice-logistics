package services

import (
	"github.com/Gierdiaz/microservice-logistics/internal/domain/entities"
	"github.com/Gierdiaz/microservice-logistics/internal/domain/repositories"
)

type ClientService struct {
	repo *repositories.ClientRepository
}

func NewClientService(repo *repositories.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Create(client *entities.Client) (*entities.Client, error) {
	return s.repo.Create(client)
}
