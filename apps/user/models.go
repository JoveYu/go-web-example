package user

type UserModel struct {
	Username string `json:"username"` // 用户名
	Password string `json:"-"`        // 密码
	Nickname string `json:"nickname"` // 昵称
}

func (UserModel) TableName() string {
	return "user"
}

func (u *UserModel) SetPassword(pass string) {
	u.Password = EncodePassword(pass)
}

func (u UserModel) CheckPassword(pass string) bool {
	return u.Password == EncodePassword(pass)
}

func EncodePassword(pass string) string {
	// TODO: get password hash
	return pass
}
