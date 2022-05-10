package src

import (
	"agent/service"
	"github.com/gin-gonic/gin"
)

func AddOauthRouter(r *gin.RouterGroup) {
	provider := r.Group("/")
	provider.GET("providers", service.GetProvider)
	provider.POST("login", service.VerifyUser)

}