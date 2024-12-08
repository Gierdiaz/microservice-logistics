package setup

import (
	"github.com/jmoiron/sqlx"

	"github.com/Gierdiaz/microservice-logistics/internal/application/handlers"
	"github.com/Gierdiaz/microservice-logistics/internal/application/services"
	"github.com/Gierdiaz/microservice-logistics/internal/domain/repositories"
)

func SetupService(db *sqlx.DB) *handlers.ClientHandler {
	clientRepo := repositories.NewClientRepository(db)
	clientService := services.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(clientService)
	return clientHandler
}