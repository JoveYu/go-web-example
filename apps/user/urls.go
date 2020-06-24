package user

import (
	"github.com/gin-gonic/gin"

	"test.com/example/apps/base"
)

func SetRouter(router *gin.Engine) {
	user := router.Group("/user")

	v1 := user.Group("/v1")
	{
		v1.POST("/signup", UserSignupView)
		v1.POST("/login", UserLoginView)
	}

	// 这里使用一个中间件，类似装饰器，其中Next方法执行下一步handler
	v1_login := user.Group("/v1", base.LoginRequiredView)
	{
		v1_login.POST("/logout", UserLogoutView)
		v1_login.GET("/profile", UserProfileView)
	}

}
