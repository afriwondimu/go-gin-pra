package main

import (
	"github.com/afriwondimu/go-gin-pra/controllers"
	"github.com/afriwondimu/go-gin-pra/initializers"
	"github.com/afriwondimu/go-gin-pra/middleware"
	"github.com/gin-gonic/gin"
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	// User Auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate",middleware.RequireAuth, controllers.Validate)
	// CRUD
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}