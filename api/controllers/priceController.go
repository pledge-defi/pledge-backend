package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"pledge-backend/api/models/ws"
	"pledge-backend/log"
	"pledge-backend/utils"
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

	ramdomId := ""
	remoteIP, ok := ctx.RemoteIP()
	if ok {
		ramdomId = remoteIP.String() + "-" + utils.GetRandomString(10)
	} else {
		ramdomId = utils.GetRandomString(30)
	}

	server := &ws.Server{
		Id:       ramdomId,
		Socket:   conn,
		Send:     make(chan []byte, 800),
		LastTime: time.Now().Unix(),
	}

	server.ReadAndWrite()
	//
	//for true {
	//	select {
	//	case ws.ServerChain <- server:
	//		log.Logger.Info(server.Id + " ws connect ... ")
	//	case <-time.After(time.Second * 3):
	//		server.SendToClient([]byte("server busy, please try later"))
	//		log.Logger.Error("server busy, please try later")
	//		return
	//	}
	//}
}
