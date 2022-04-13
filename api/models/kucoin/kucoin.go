package kucoin

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"pledge-backend/db"
	"pledge-backend/log"
)

// ApiKeyVersionV2 is v2 api key version
const ApiKeyVersionV2 = "2"

var PlgrPrice = "0.0027"
var PlgrPriceChan = make(chan string, 2)

func GetExchangePrice() {

	log.Logger.Sugar().Info("GetExchangePrice ")

	// get plgr price from redis
	price, err := db.RedisGetString("plgr_price")
	if err != nil {
		log.Logger.Sugar().Error("get plgr price from redis err ", err)
	} else {
		PlgrPrice = price
	}

	s := kucoin.NewApiService(
		kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
		kucoin.ApiKeyVersionOption(ApiKeyVersionV2),
	)

	rsp, err := s.WebSocketPublicToken()
	if err != nil {
		log.Logger.Error(err.Error()) // Handle error
		return
	}

	tk := &kucoin.WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		log.Logger.Error(err.Error())
		return
	}

	c := s.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		log.Logger.Sugar().Errorf("Error: %s", err.Error())
		return
	}

	ch := kucoin.NewSubscribeMessage("/market/ticker:PLGR-USDT", false)
	uch := kucoin.NewUnsubscribeMessage("/market/ticker:PLGR-USDT", false)

	if err := c.Subscribe(ch); err != nil {
		log.Logger.Error(err.Error()) // Handle error
		return
	}

	for {
		select {
		case err := <-ec:
			c.Stop() // Stop subscribing the WebSocket feed
			log.Logger.Sugar().Errorf("Error: %s", err.Error())
			_ = c.Unsubscribe(uch)
			return
		case msg := <-mc:
			t := &kucoin.TickerLevel1Model{}
			if err := msg.ReadData(t); err != nil {
				log.Logger.Sugar().Errorf("Failure to read: %s", err.Error())
				return
			}
			PlgrPriceChan <- t.Price
			PlgrPrice = t.Price
			//log.Logger.Sugar().Info("Price ", t.Price)
			_ = db.RedisSetString("plgr_price", PlgrPrice, 0)
		}
	}
}
