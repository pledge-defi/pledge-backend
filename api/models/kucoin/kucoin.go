package kucoin

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"pledge-backend/log"
)

// ApiKeyVersionV2 is v2 api key version
const ApiKeyVersionV2 = "2"

func GetExchangePrice() {
	log.Logger.Info("2wdasdasq12323455etrdfsdzc")
	s := kucoin.NewApiService(
		// kucoin.ApiBaseURIOption("https://api.kucoin.com"),
		kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
		kucoin.ApiKeyVersionOption(ApiKeyVersionV2),
	)
	log.Logger.Info("6666666")

	rsp, err := s.WebSocketPublicToken()
	if err != nil {
		// Handle error
		log.Logger.Error(err.Error())
		return
	}

	tk := &kucoin.WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		log.Logger.Error(err.Error())
		return
	}
	log.Logger.Info("555555")

	c := s.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	ch := kucoin.NewSubscribeMessage("/market/ticker:PLGR-USDT", false)
	uch := kucoin.NewUnsubscribeMessage("/market/ticker:PLGR-USDT", false)

	//if err := c.Subscribe(ch1, ch2); err != nil {
	if err := c.Subscribe(ch); err != nil {
		// Handle error
		return
	}

	var i = 0
	for {
		select {
		case err := <-ec:
			c.Stop() // Stop subscribing the WebSocket feed
			log.Logger.Sugar().Errorf("Error: %s", err.Error())
			// Handle error
			return
		case msg := <-mc:
			// log.Printf("Received: %s", kucoin.ToJsonString(m))
			t := &kucoin.TickerLevel1Model{}
			if err := msg.ReadData(t); err != nil {
				log.Logger.Sugar().Errorf("Failure to read: %s", err.Error())
				return
			}
			log.Logger.Sugar().Infof("Ticker: %s, %s, %s, %s", msg.Topic, t.Sequence, t.Price, t.Size)
			i++
			if i == 5 {
				log.Logger.Sugar().Info("Unsubscribe PLGR-USDT")
				if err = c.Unsubscribe(uch); err != nil {
					log.Logger.Sugar().Errorf("Error: %s", err.Error())
					// Handle error
					return
				}
			}
			if i == 10 {
				log.Logger.Info("Subscribe PLGR-USDT")
				if err = c.Subscribe(ch); err != nil {
					log.Logger.Sugar().Errorf("Error: %s", err.Error())
					// Handle error
					return
				}
			}
			if i == 15 {
				log.Logger.Info("Exit subscription")
				c.Stop() // Stop subscribing the WebSocket feed
				return
			}
		}
	}
}
