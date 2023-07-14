package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

// 参数绑定并校验
func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		value := c.Value("trans")
		translator := value.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}
		for k, v := range verrs.Translate(translator) {
			errs = append(errs, &ValidError{
				Key:     k,
				Message: v,
			})
		}
		return false, errs
	}
	return true, nil
}
