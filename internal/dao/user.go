package dao

import "v1/internal/model"

func (d *Dao) CreateUser(username, password, email, avatar, bio string) error {
	u := model.User{
		Username:  username,
		Password:  password,
		Email:     email,
		Avatar:    avatar,
		Biography: bio,
	}
	return u.Create(d.engine)
}

func (d *Dao) QueryUserByName(username string) (*model.User, error) {
	u := model.User{
		Username: username,
	}
	return u.QueryItemByName(d.engine)
}

func (d *Dao) UpdateUserById(id uint, username, password, email, avatar, bio string) error {
	u := model.User{
		Username:  username,
		Password:  password,
		Email:     email,
		Avatar:    avatar,
		Biography: bio,
	}
	u.ID = id
	return u.Update(d.engine)
}
