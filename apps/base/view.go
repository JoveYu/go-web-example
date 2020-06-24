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
	if session.Get("userid") == nil {
		ctx.JSON(200, Response{
			ErrorCode: -1,
			Msg:       "login required",
		})
		return
	}
	// 下一个handler处理
	ctx.Next()
}
