package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	S  *gin.Engine
}
