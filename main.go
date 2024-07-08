package main

import (
	"fmt"

	_ "go-gin-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@Swagger Example API
//@version 1.0
//@Description this is an example of using swagger with gin framework in golang
//@termsOfService http://swagger.io/terms
// @host localhost:8080
// @BasePath /api/

func main() {
	r := gin.Default()

	r.GET("/api/api/ping", ping)
	r.GET("/api/api/hello", hello)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
	fmt.Println("hello")
}

// Ping example
// @Summary Show a ping pong message
// @Description Get ping pong response
// @Tags ping
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "pong"
// @Router /api/ping [get]
func ping(c *gin.Context) {
	c.String(200, "pong")
}

// Hello example
// @Summary Show a hello message
// @Description Return Hello response
// @Tags hello
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "hello world"
// @Router /api/hello [get]
func hello(c *gin.Context) {
	c.String(200, "hello world")
}
