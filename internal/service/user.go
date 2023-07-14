package service

type LoginRequest struct {
	Username string `form:"username" binding:"required,min=5"`
	Password string `form:"password" binding:"required,min=6"`
}

func (s *Service) Login(param *LoginRequest) {
}
