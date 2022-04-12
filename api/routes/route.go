package routes

import (
	"github.com/gin-gonic/gin"
	"pledge-backend/api/controllers"
	"pledge-backend/api/middlewares"
	"pledge-backend/config"
)

func InitRoute(e *gin.Engine) *gin.Engine {

	// version group
	v2Group := e.Group("/api/v" + config.Config.Env.Version)

	// pledge-defi backend
	poolController := controllers.PoolController{}
	v2Group.GET("/poolBaseInfo", poolController.PoolBaseInfo)                                   //pool base information
	v2Group.GET("/poolDataInfo", poolController.PoolDataInfo)                                   //pool data information
	v2Group.GET("/token", poolController.TokenList)                                             //pool token information
	v2Group.POST("/pool/debtTokenList", middlewares.CheckToken(), poolController.DebtTokenList) //pool debtTokenList
	v2Group.POST("/pool/search", middlewares.CheckToken(), poolController.Search)               //pool search

	// plgr-usdt price
	priceController := controllers.PriceController{}
	v2Group.GET("/price", priceController.NewPrice) //new price on ku-coin-exchange

	// pledge-defi admin backend
	multiSignPoolController := controllers.MultiSignPoolController{}
	v2Group.POST("/pool/setMultiSign", middlewares.CheckToken(), multiSignPoolController.SetMultiSign) //multi-sign set
	v2Group.POST("/pool/getMultiSign", middlewares.CheckToken(), multiSignPoolController.GetMultiSign) //multi-sign get

	userController := controllers.UserController{}
	v2Group.POST("/user/login", userController.Login)                             // login
	v2Group.POST("/user/logout", middlewares.CheckToken(), userController.Logout) // logout

	return e
}
