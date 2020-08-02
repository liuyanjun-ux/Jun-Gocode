package router

import (
	"github.com/garylailai/cateraway-api/handlers/api/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/cateraway", apis.FoodAllQuery)

	//router.POST("/user", Store)
	//
	//router.PUT("/user/:id", Update)
	//
	//router.DELETE("/user/:id", Destroy)

	return router
}
