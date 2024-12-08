package main

import (
	"log"

	"github.com/Gierdiaz/microservice-logistics/config"

	"github.com/Gierdiaz/microservice-logistics/internal/infrastructure/endpoints"
	"github.com/Gierdiaz/microservice-logistics/internal/infrastructure/persistence"
	"github.com/Gierdiaz/microservice-logistics/pkg/logger"
)

func main() {

	logger.InitLogger()
	defer logger.Sync()

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

    db, err := persistence.InitDatabase(config)
    if err != nil {
        log.Fatalf("Erro ao inicializar o banco de dados: %s", err)
    }

	router := endpoints.NewRouter(db)
    router.Run(":" + config.Server.APP_PORT)
}
