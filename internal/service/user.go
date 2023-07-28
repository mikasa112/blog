package service

import (
	"v1/pkg"

	"go.uber.org/zap"
	e "v1/pkg/err"
	"v1/pkg/util"
)

type LoginRequest struct {
	Username string `form:"username" binding:"required,min=5"`
	Password string `form:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Username string `form:"username" binding:"required,min=5"`
	Password string `form:"password" binding:"required,min=6"`
}

func (s *Service) Login(param LoginRequest) (string, *e.Error) {
	u, err := s.dao.QueryUserByName(param.Username)
	if err != nil {
		return "", e.UsernameOrPasswordError
	}
	pwd := util.MD5Encode(param.Password)
	if pwd != u.Password {
		return "", e.UsernameOrPasswordError
	}
	return "", nil
}

func (s *Service) GenerateUser(username, password string) *e.Error {
	pkg.Log.Info("生成User", zap.String("username", username), zap.String("password", password))
	count, err1 := s.dao.QueryUsersCount()
	if count == 1 || err1 != nil {
		pkg.Log.Sugar().Debug("用户列表长度已不为1")
		return e.UserIsExist
	}
	u, err := s.dao.QueryUserByName(username)
	if u != nil {
		return e.UserIsExist
	}
	err = s.dao.CreateUser(username, util.MD5Encode(password), "", "", "")
	if err != nil {
		pkg.Log.Error("Create User Error", zap.String("error", err.Error()))
		return e.ServerError
	}
	return nil
}
