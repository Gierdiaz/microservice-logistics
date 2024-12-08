package contracts

import "github.com/Gierdiaz/microservice-logistics/internal/domain/entities"

type ClientRepository interface {
	Create(client *entities.Client) (*entities.Client, error)
}