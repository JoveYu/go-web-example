package base

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PingView(ctx *gin.Context) {
	ctx.String(200, "OK")
}

func LoginRequiredView(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if session.Get("username") == nil {
		ctx.JSON(200, NewResponse(ERR_LOGIN))
		ctx.Abort()
		return
	}
	// 下一个handler处理
	ctx.Next()
}
