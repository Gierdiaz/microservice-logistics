package repositories

import (
	"github.com/Gierdiaz/microservice-logistics/internal/domain/entities"
	"github.com/jmoiron/sqlx"
)

type ClientRepository struct {
    db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) *ClientRepository{
    return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(client *entities.Client) (*entities.Client, error) {
	query := "INSERT INTO clients (id, name, email) VALUES ($1, $2, $3) RETURNING id, name, email"
	err := r.db.QueryRow(query, client.ID, client.Name, client.Email).Scan(&client.ID, &client.Name, &client.Email)
	if err != nil {
		return nil, err
	}
	return client, nil
}