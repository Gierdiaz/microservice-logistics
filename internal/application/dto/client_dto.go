package dto

import "github.com/Gierdiaz/microservice-logistics/internal/domain/entities"

type ClientDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (dto *ClientDTO) ToEntity() *entities.Client {
	return &entities.Client{
		Name:  dto.Name,
		Email: dto.Email,
	}
}
