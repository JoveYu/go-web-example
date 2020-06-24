package user

import (
	"github.com/gin-gonic/gin"

	"test.com/example/apps/base"
)

type UserSignupRequest struct {
	Username   string `json:"username" binding:"required"`                     // 用户名
	Password   string `json:"password" binding:"required"`                     // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 重复密码
	Nickname   string `json:"nickname" binding:"required"`                     // 昵称
}

// @Summary 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param json body UserSignupRequest true "请求正文"
// @Success 200 {object} base.Response
// @Router /user/v1/signup [post]
func UserSignupView(ctx *gin.Context) {

}

type UserLoginRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type UserLoginResponse struct {
	base.Response
	Result UserModel `json:"result"` // 用户信息
}

// @Summary 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param json body UserLoginRequest true "请求正文"
// @Success 200 {object} UserLoginResponse
// @Router /user/v1/login [post]
func UserLoginView(ctx *gin.Context) {

}

// @Summary 用户登出
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} base.Response
// @Router /user/v1/logout [post]
func UserLogoutView(ctx *gin.Context) {

}

type UserProfileResponse struct {
	base.Response
	Result UserModel `json:"result"`
}

// @Summary 用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} UserProfileResponse
// @Router /user/v1/profile [get]
func UserProfileView(ctx *gin.Context) {

}
