package routes

import (
	"github.com/go-2024/controllers"
	"github.com/go-2024/models"
)

func PostRoutes(ir *models.Server) {

	postController := controllers.NewPostController(ir.DB)

	ir.S.GET("/getposts", postController.GetPosts)
	ir.S.POST("/createpost", postController.CreatePost)
	ir.S.GET("/getpost/:id", postController.GetPostByID)
	ir.S.DELETE("/deletepost/:id", postController.DeletePostByID)
	ir.S.PUT("/updatepost/:id", postController.UpdatePostByID)

}
