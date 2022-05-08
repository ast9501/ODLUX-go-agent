package src

import (
	"agent/service"
	"github.com/gin-gonic/gin"
)

func AddOperationRouter(r *gin.RouterGroup){
	reader := r.Group("/operations")

	reader.POST("/data-provider:operation", service.Handler)
	/*
	// Read netconf devices connection list
	reader.POST("/data-provider:read-network-element-connection-list", service.GetConnectDevice)
	// Mount new device
	reader.POST("/data-provider:create-network-element-connection", service.CreateDevice)
	// Get device connection logs
	reader.POST("/data-provider:read-connectionlog-list", service.ReadDeviceLog)
	*/
}