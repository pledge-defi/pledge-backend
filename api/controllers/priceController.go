package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models"
	"pledge-backend/api/models/request"
	"pledge-backend/api/models/response"
	"pledge-backend/api/models/ws"
	"pledge-backend/api/services"
	"pledge-backend/api/validate"
	"pledge-backend/log"
	"time"
)

type PriceController struct {
}

func (c *PriceController) NewPrice(ctx *gin.Context) {

	defer func() {
		recover := recover()
		if recover != nil {
			log.Logger.Sugar().Error("new price recover ", recover)
		}
	}()

	conn, err := (&websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool { //Cross domain
			return true
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Logger.Sugar().Error("websocket request err:", err)
		return
	}

	server := &ws.Server{
		Id:       ctx.Request.RemoteAddr,
		Socket:   conn,
		Send:     make(chan []byte, 800),
		LastTime: time.Now().Unix(),
	}

	for true {
		select {
		case ws.ServerChain <- server:
			log.Logger.Info(server.Id + " connect success !")
		case <-time.After(time.Second * 3):
			server.SendToClient([]byte("server busy, please try later"))
			log.Logger.Error("server busy, please try later")
			return
		}
	}
}

func (c *PriceController) PoolDataInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.PoolDataInfo{}
	var result []models.PoolDataInfoRes

	errCode := validate.NewPoolDataInfo().PoolDataInfo(ctx, &req)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode = services.NewPool().PoolDataInfo(req.ChainId, &result)
	if errCode != statecode.COMMON_SUCCESS {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.COMMON_SUCCESS, result)
	return
}
