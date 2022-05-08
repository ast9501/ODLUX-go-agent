package src

import (
	"agent/service"
	"github.com/gin-gonic/gin"
)

func AddDefaultRouter(r *gin.RouterGroup) {
	logs := r.Group("/")

	// Get device connection logs
	logs.GET("netconflogs", service.GetDeviceLog)
	// Receive device connection logs from onos
	logs.POST("netconglogs", service.PostDeviceLog)
	// Return api status
	logs.GET("status", service.GetStatus)
}