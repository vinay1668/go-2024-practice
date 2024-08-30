package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-2024/database"
	"github.com/go-2024/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// server, err := internal.NewServer()
	// if err != nil {
	// 	panic("")
	// }

	// routes.PostRoutes(server)
	var err error
	DB, err = database.Init()
	if err != nil {
		log.Fatal("eroor")
	}
	router := gin.Default()

	router.GET("/getposts", GetPosts)
	router.POST("/createpost", CreatePost)
	router.GET("/getpost/:id", GetPostByID)
	router.DELETE("/deletepost/:id", DeletePostByID)
	router.PUT("/updatepost/:id", UpdatePostByID)

	router.Run(":8080")

}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if result := DB.Find(&posts); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": result.Error.Error(),
		})
		return
	}
	c.JSON(200, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if result := DB.Create(&post); result.Error != nil {
		c.JSON(500, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(201, gin.H{"msg": "Post Created Successfullly"})

}
func GetPostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	if result := DB.First(&post, ID); result.Error != nil {
		c.JSON(500, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(200, post)
}

func DeletePostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	if result := DB.Delete(&post, ID); result.Error != nil {
		c.JSON(500, gin.H{"err": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "Post deleted successfully"})
}
func UpdatePostByID(c *gin.Context) {
	ID := c.Param("id")
	var post models.Post
	if result := DB.First(&post, ID); result.Error != nil {
		c.JSON(404, gin.H{"err": result.Error.Error()})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.JSON(500, gin.H{"err": err})
		return
	}
	DB.Save(&post)
	c.JSON(200, &post)

}
