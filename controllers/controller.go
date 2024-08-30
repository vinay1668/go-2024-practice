package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-2024/models"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		DB: db,
	}
}

func (pc *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	if result := pc.DB.Find(&posts); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": result.Error.Error(),
		})
		return
	}
	c.JSON(200, posts)
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if result := pc.DB.Create(&post); result.Error != nil {
		c.JSON(500, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(201, gin.H{"msg": "Post Created Successfullly"})

}
func (pc *PostController) GetPostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	if result := pc.DB.First(&post, ID); result.Error != nil {
		c.JSON(500, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(200, post)
}

func (pc *PostController) DeletePostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	if result := pc.DB.Delete(&post, ID); result.Error != nil {
		c.JSON(500, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "Post deleted successfully"})
}
func (pc *PostController) UpdatePostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	if result := pc.DB.First(&post, ID); result.Error != nil {
		c.JSON(404, gin.H{"err": result.Error.Error()})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.JSON(500, gin.H{"err": err})
		return
	}
	pc.DB.Save(&post)
	c.JSON(200, &post)

}
