package main

import (
	"github.com/afriwondimu/go-gin-pra/initializers"
	"github.com/afriwondimu/go-gin-pra/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}