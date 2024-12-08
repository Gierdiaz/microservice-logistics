package handlers

import (
	"net/http"

	"github.com/Gierdiaz/microservice-logistics/internal/application/dto"
	"github.com/Gierdiaz/microservice-logistics/internal/application/services"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientService *services.ClientService
}

func NewClientHandler(clientService *services.ClientService) *ClientHandler {
	return &ClientHandler{clientService: clientService}
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var clientDTO dto.ClientDTO

	if err := c.ShouldBindJSON(&clientDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	client := clientDTO.ToEntity()
	createdClient, err := h.clientService.Create(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create client",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Client created successfully",
		"client":  createdClient,
	})
}
