package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//gin bind go-playground/validate
func BindingValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("IsPassword", IsPassword)                           //判断是否为合法密码
		v.RegisterValidation("IsPhoneNumber", IsPhoneNumber)                     //检查手机号码字段是否合法
		v.RegisterValidation("IsEmail", IsEmail)                                 //检查邮箱字段是否合法
		v.RegisterValidation("CheckUserNicknameLength", CheckUserNicknameLength) //检查用户昵称长度是否合法
		v.RegisterValidation("CheckUserAccount", CheckUserAccount)               //检查用户账号是否合法
		v.RegisterValidation("OnlyOne", OnlyOne)                                 //字段唯一性约束
	}
}
