package endpoints

import (
	"github.com/Gierdiaz/microservice-logistics/internal/interfaces/setup"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()

	clientHandler := setup.SetupService(db)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/clients", clientHandler.CreateClient)
	}

	return r
}
