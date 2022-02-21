package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (v *User) Login(c *gin.Context, req *request.Login) int {

	err := c.ShouldBind(req)
	if err == io.EOF {
		return statecode.PARAMETER_EMPTY_ERR
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "Name" && e.Tag() == "required" {
				return statecode.P_NAME_EMPTY
			}
			if e.Field() == "Password" && e.Tag() == "required" {
				return statecode.P_NAME_EMPTY
			}
		}
		return statecode.COMMON_ERR_SERVER_ERR
	}

	return statecode.COMMON_SUCCESS
}
