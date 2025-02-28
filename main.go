package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}
func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}
func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.POST("/register", controller.Register)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntry)

	router.Run(":8080")
	fmt.Println("Server 8080 succ running")
}
