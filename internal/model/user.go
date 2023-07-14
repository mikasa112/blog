package model

import "gorm.io/gorm"

type User struct {
	Username  string `gorm:"comment:用户名;not null;"`
	Password  string `gorm:"comment:密码;not null;"`
	Email     string `gorm:"comment:邮箱"`
	Avatar    string `gorm:"comment:头像路径"`
	Biography string `gorm:"comment:bio;type:text"`
	*Model
}

func (u User) QueryItemByName(db *gorm.DB) (*User, error) {
	var user *User
	d := db.Where("name = ?", u.Username).First(user)
	if d.Error != nil {
		return nil, d.Error
	}
	return user, nil
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Update(db *gorm.DB) error {
	return db.Model(&User{}).Where("id = ?", u.ID).Updates(u).Error
}

func (u User) Delete(db *gorm.DB) error {
	return db.Where("id = ?", u.ID).Delete(&u).Error
}
