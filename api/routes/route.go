package routes

import (
	"github.com/gin-gonic/gin"
	"pledge-backend/api/controllers"
	"pledge-backend/api/middlewares"
	"pledge-backend/config"
)

func InitRoute(e *gin.Engine) *gin.Engine {

	// version group
	group := e.Group("/api/v" + config.Config.Env.Version)

	// pledge-defi backend
	poolController := controllers.PoolController{}
	group.GET("/poolBaseInfo", poolController.PoolBaseInfo)                                   //pool base information
	group.GET("/poolDataInfo", poolController.PoolDataInfo)                                   //pool data information
	group.GET("/token", poolController.TokenList)                                             //pool token information
	group.POST("/pool/debtTokenList", middlewares.CheckToken(), poolController.DebtTokenList) //pool debtTokenList
	group.POST("/pool/search", middlewares.CheckToken(), poolController.Search)               //pool search

	// plgr-usdt price
	priceController := controllers.PriceController{}
	group.GET("/price", priceController.NewPrice) //new price on ku-coin-exchange

	// pledge-defi admin backend
	multiSignPoolController := controllers.MultiSignPoolController{}
	group.POST("/pool/setMultiSign", middlewares.CheckToken(), multiSignPoolController.SetMultiSign) //multi-sign set
	group.POST("/pool/getMultiSign", middlewares.CheckToken(), multiSignPoolController.GetMultiSign) //multi-sign get

	userController := controllers.UserController{}
	group.POST("/user/login", userController.Login)                             // login
	group.POST("/user/logout", middlewares.CheckToken(), userController.Logout) // logout

	return e
}
