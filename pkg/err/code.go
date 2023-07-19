package err

var (
	Success                   = NewErr(0, "成功")
	ServerError               = NewErr(10000, "服务内部错误")
	InvalidParams             = NewErr(10001, "入参错误")
	NotFound                  = NewErr(10002, "找不到")
	UnauthorizedAuthNotExist  = NewErr(10003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewErr(10004, "鉴权失败, Token 错误")
	UnauthorizedTokenTimeout  = NewErr(10005, "鉴权失败, Token 超时")
	UnauthorizedTokenGenerate = NewErr(10006, "鉴权失败, Token 生成失败")
	TooManyRequests           = NewErr(10007, "请求过多")

	UsernameOrPasswordError = NewErr(10008, "用户名或密码错误")
)
