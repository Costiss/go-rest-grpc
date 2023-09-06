package rest

import (
	"strconv"

	"github.com/costiss/go-rest-grpc/internal/service"
	"github.com/gin-gonic/gin"
)

func MakeRestServer() *gin.Engine {

	server := gin.Default()

	server.GET("/:first/:second", func(c *gin.Context) {
		first := c.Param("first")
		second := c.Param("second")

		firstInt, _ := strconv.Atoi(first)
		secondInt, _ := strconv.Atoi(second)

		c.JSON(200, gin.H{
			"result": service.Add(firstInt, secondInt),
		})
	})

	return server

}
