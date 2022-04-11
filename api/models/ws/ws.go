package ws

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
	"pledge-backend/api/models/kucoin"
	"pledge-backend/log"
	"sync"
	"time"
)

const SuccessCode = 0
const ErrorCode = -1

type Server struct {
	sync.Mutex
	Id       string
	Socket   *websocket.Conn
	Send     chan []byte
	LastTime int64 // last send time
}

type ServerManager struct {
	Servers    sync.Map
	Broadcast  chan []byte
	Register   chan *Server
	Unregister chan *Server
}

type Message struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

var Manager = ServerManager{}
var UserPingPongDurTime int64 = 20 // seconds
var ServerChain = make(chan *Server)

func (s *Server) SendToClient(data string, code int) {
	s.Lock()
	defer s.Unlock()

	dataBytes, err := json.Marshal(Message{
		Code: code,
		Data: data,
	})
	err = s.Socket.WriteMessage(websocket.TextMessage, dataBytes)
	if err != nil {
		log.Logger.Sugar().Error(s.Id+" SendToClient err ", err)
	}
}

func (s *Server) ReadAndWrite() {

	eg := errgroup.Group{}

	//write
	eg.Go(func() error {
		for {
			select {
			case message, ok := <-s.Send:
				if !ok {
					Manager.Servers.Delete(s)
					return errors.New("write message error")
				}
				s.SendToClient(string(message), SuccessCode)
			}
		}
	})

	//read
	eg.Go(func() error {
		for {
			_, message, err := s.Socket.ReadMessage()
			if err != nil {
				log.Logger.Sugar().Error(s.Id+" ReadMessage err ", err)
				return err
			}
			//update ping time
			if string(message) == "ping" || string(message) == `"ping"` || string(message) == "'ping'" {
				s.LastTime = time.Now().Unix()
				s.Send <- []byte("pong")
			}
			continue
		}
	})

	//check ping pong
	eg.Go(func() error {
		for {
			select {
			case <-time.After(time.Second):
				if time.Now().Unix()-s.LastTime >= UserPingPongDurTime {
					s.SendToClient("ping timeout", ErrorCode)
					log.Logger.Info(s.Id + " timeout")
					Manager.Servers.Delete(s)
					return errors.New("ping timeout")
				}
			}
		}
	})

	if err := eg.Wait(); err != nil {
		log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
		return
	}

}

func StartServer() {
	log.Logger.Info("WsServer start")
	for {
		select {
		case s, ok := <-ServerChain:
			if ok {
				Manager.Servers.Store(s.Id, s)
				go s.ReadAndWrite()
			}
		case price, ok := <-kucoin.PlgrPriceChain:
			if ok {
				Manager.Servers.Range(func(key, value interface{}) bool {
					value.(*Server).SendToClient(price, SuccessCode)
					return true
				})
			}
		}
	}
}
