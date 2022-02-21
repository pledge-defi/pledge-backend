package services

import (
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/log"
	"pledge-backend/utils"
)

type UserService struct{}

func NewUser() *UserService {
	return &UserService{}
}

func (s *UserService) Login(req *request.Login, result *response.Login) int {
	log.Logger.Sugar().Info("contractService", req)
	if req.Name == "admin" && req.Password == "password" {
		token, err := utils.CreateToken(req.Name)
		if err != nil {
			log.Logger.Error("CreateToken" + err.Error())
			return statecode.COMMON_ERR_SERVER_ERR
		}
		result.TokenId = token
		//save to redis
		db.RedisSet(req.Name, "login_ok", config.Config.Jwt.ExpireTime)
		return statecode.COMMON_SUCCESS
	} else {
		return statecode.NAME_OR_PASSWORD_ERR
	}
}
