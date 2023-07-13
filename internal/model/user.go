package model

type User struct {
	Username  string `gorm:"comment:用户名"`
	Password  string `gorm:"comment:密码"`
	Email     string `gorm:"comment:邮箱"`
	Avatar    string `gorm:"comment:头像路径"`
	Biography string `gorm:"comment:bio;type:text"`
	*Model
}
