package main

import (
	. "agent/src"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Handle api for odl
	rests := router.Group("/rests")
	AddOperationRouter(rests)

	data := router.Group("/rests/data")
	AddDataRouter(data)

	//Handel api for odlux user login
	oauth := router.Group("/oauth")
	AddOauthRouter(oauth)

	// Handle api for other application
	v1 := router.Group("/v1")
	AddDefaultRouter(v1)



	router.Run(":8001")
}