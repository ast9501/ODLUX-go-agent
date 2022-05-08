package main

import (
	. "agent/src"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	rests := router.Group("/rests")
	AddOperationRouter(rests)

	v1 := router.Group("/v1")
	AddDefaultRouter(v1)

	data := router.Group("/rests/data")
	AddDataRouter(data)


	router.Run(":8001")
}