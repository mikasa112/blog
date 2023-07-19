package service

import (
	"crypto/md5"
	"v1/pkg"

	"go.uber.org/zap"
	e "v1/pkg/err"
	"v1/pkg/util"
)

type LoginRequest struct {
	Username string `form:"username" binding:"required,min=5"`
	Password string `form:"password" binding:"required,min=6"`
}

func (s *Service) Login(param *LoginRequest) (string, error) {
	u, err := s.dao.QueryUserByName(param.Username)
	if err != nil {
		return "", e.UsernameOrPasswordError
	}
	pwd := util.MD5Encode(param.Password)
	if pwd == u.Password {

		return "", nil
	}
	return "", e.UsernameOrPasswordError
}

func (s *Service) GenerateUser(username, password string) {
	pwdBytes := []byte(password)
	md5 := md5.New()
	encryptPwd := md5.Sum(pwdBytes)
	err := s.dao.CreateUser(username, string(encryptPwd), "", "", "")
	if err != nil {
		pkg.Log.Error("Create User Error", zap.String("error", err.Error()))
	}
}
