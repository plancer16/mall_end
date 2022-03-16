package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	gin.SetMode(viper.GetString("model"))

	user := r.Group("/api/user")
	{
		user.POST("/add",UserHandler.AddUserHandler)
	}
	port := viper.GetString("port")
	r.Run(port)
}