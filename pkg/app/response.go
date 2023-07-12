package app

import (
	"net/http"
	"v1/pkg"
	"v1/pkg/err"
	"v1/pkg/util"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ctx *gin.Context
}

func NewResponse(context *gin.Context) *Response {
	return &Response{
		ctx: context,
	}
}

// 响应单条数据
func (r *Response) To(data any) {
	if data == nil {
		data = gin.H{}
	}
	r.ctx.JSON(http.StatusOK, data)
}

// 响应分页数据
func (r *Response) ListTo(list any, total int) {
	if total < 0 {
		total = 0
	}
	r.ctx.JSON(http.StatusOK, gin.H{
		"list":      list,
		"page":      GetPage(r.ctx),
		"page_size": GetPageSize(r.ctx),
		"total":     total,
	})
}

// 响应错误
func (r *Response) ErrTo(err *err.Error) {
	h := gin.H{"code": err.Code(), "msg": err.Msg()}
	s := err.Details()
	if len(s) > 0 {
		h["details"] = s
	}
	r.ctx.JSON(err.StatusCode(), h)
}

// 查询到page参数,例如http://xxx.xxx.x.x/xx?page=3
func GetPage(c *gin.Context) int {
	page := util.Str(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

// 查询到page_size参数
func GetPageSize(c *gin.Context) int {
	pageSize := util.Str(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return pkg.Sc.PageSize
	}
	if pageSize > pkg.Sc.MaxPageSize {
		return pkg.Sc.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
