package main

import (
	"app-user/configs"
	"app-user/controllers"
	docs "app-user/docs"
	"context"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title App API
// @version 1.0
// @description API for management user and login App Project

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(CORSMiddleware())
	docs.SwaggerInfo.BasePath = "/"
	mongoCollection, err := configs.MongoConfig()
	if err != nil {
		panic(err)
	}
	defer mongoCollection.Disconnect(context.TODO())
	userController := controllers.InitUserController(configs.MongoCollection)

	mainGroup := router.Group("/api/v1")
	{
		account := mainGroup.Group("/user")
		{
			account.POST("", userController.CreateUserController)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":6003")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
