package icmp

import (
    "github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	route.GET("/ping", Health)
}

