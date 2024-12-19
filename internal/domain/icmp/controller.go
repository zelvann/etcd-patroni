package icmp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zelvann/etcd-patroni/internal/utils"
)

func Health(ctx *gin.Context) {
	res := utils.NewSucessResponse("Pong")
	ctx.JSON(http.StatusOK, res)
}
