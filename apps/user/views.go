package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"test.com/example/apps/base"
	"test.com/example/common/util"
	"test.com/example/config"
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
	req := UserSignupRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, base.NewResponse(base.ERR_PARAMS, util.WrapValidationError(err)))
		return
	}

	count := -1
	config.DB.Model(&UserModel{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		ctx.JSON(200, base.NewResponse(base.ERR_PARAMS, "", "用户已存在"))
		return
	}

	user := UserModel{
		Username: req.Username,
		Nickname: req.Nickname,
	}
	user.SetPassword(req.Password)

	err = config.DB.Create(user).Error
	if err != nil {
		ctx.JSON(200, base.NewResponse(base.ERR_DB, err.Error()))
		return
	}

	ctx.JSON(200, base.NewResponse(base.OK))
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
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
	req := UserLoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, base.NewResponse(base.ERR_PARAMS, util.WrapValidationError(err)))
		return
	}

	user := UserModel{}
	err = config.DB.First(&user, "username=?", req.Username).Error
	if err != nil {
		ctx.JSON(200, base.NewResponse(base.ERR_USER, err.Error(), "用户不存在"))
		return
	}

	if !user.CheckPassword(req.Password) {
		ctx.JSON(200, base.NewResponse(base.ERR_USER, "", "密码错误"))
		return
	}

	session := sessions.Default(ctx)
	session.Set("username", user.Username)
	_ = session.Save()

	ctx.JSON(200, UserLoginResponse{
		Response: base.NewResponse(base.OK),
		Result:   user,
	})

}

// @Summary 用户登出
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} base.Response
// @Router /user/v1/logout [post]
func UserLogoutView(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	_ = session.Save()

	ctx.JSON(200, base.NewResponse(base.OK))

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
	username := ctx.GetString("username")
	user := UserModel{}
	err := config.DB.First(&user, "username=?", username).Error
	if err != nil {
		ctx.JSON(200, base.NewResponse(base.ERR_DB, ""))
		return
	}
	ctx.JSON(200, UserProfileResponse{
		Response: base.NewResponse(base.OK),
		Result:   user,
	})
}
