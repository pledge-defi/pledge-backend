package middlewares

import (
	"github.com/gin-gonic/gin"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/response"
	"pledge-backend/config"
	"pledge-backend/db"
	"pledge-backend/utils"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := response.Gin{Res: c}
		token := c.Request.Header.Get("authCode")

		username, err := utils.ParseToken(token, config.Config.Jwt.SecretKey)
		if err != nil {
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}

		if username != config.Config.DefaultAdmin.Username {
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}

		// Judge whether the user logout
		resByteArr, err := db.RedisGet(username)
		if string(resByteArr) != `"login_ok"` {
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}

		c.Set("username", username)

		c.Next()
	}
}
