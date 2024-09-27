package main

import (
	"homburg/backend/initializers"
	"homburg/backend/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Forum{})
	initializers.DB.AutoMigrate(&models.Thread{})
	initializers.DB.AutoMigrate(&models.Post{})
}
