package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/go-2024/database"
	"github.com/go-2024/models"
)

func NewServer() (*models.Server, error) {
	db, err := database.Init()
	if err != nil {
		return nil, err
	}
	return &models.Server{
		DB: db,
		S:  gin.Default(),
	}, nil
}
