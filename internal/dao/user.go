package dao

import "v1/internal/model"

func (d *Dao) CreateUser(username, password, email, avatar, biography string) error {
	u := &model.User{
		Username:  username,
		Password:  password,
		Email:     email,
		Avatar:    avatar,
		Biography: biography,
	}
	return d.engine.Create(u).Error
}
