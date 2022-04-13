package ws

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"pledge-backend/api/models/kucoin"
	"pledge-backend/config"
	"pledge-backend/log"
	"sync"
	"time"
)

const SuccessCode = 0
const PongCode = 1
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
var UserPingPongDurTime = config.Config.Env.WssTimeoutDuration // seconds

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

	errChan := make(chan error)

	Manager.Servers.Store(s.Id, s)

	defer func() {
		Manager.Servers.Delete(s)
		_ = s.Socket.Close()
		close(s.Send)
	}()

	//write
	go func() {
		for {
			select {
			case message, ok := <-s.Send:
				if !ok {
					errChan <- errors.New("write message error")
					return
				}
				s.SendToClient(string(message), SuccessCode)
			}
		}
	}()

	//read
	go func() {
		for {

			_, message, err := s.Socket.ReadMessage()
			if err != nil {
				log.Logger.Sugar().Error(s.Id+" ReadMessage err ", err)
				errChan <- err
				return
			}

			//update heartbeat time
			if string(message) == "ping" || string(message) == `"ping"` || string(message) == "'ping'" {
				s.LastTime = time.Now().Unix()
				s.SendToClient("pong", PongCode)
			}
			continue

		}
	}()

	//check heartbeat
	for {
		select {
		case <-time.After(time.Second):
			if time.Now().Unix()-s.LastTime >= UserPingPongDurTime {
				s.SendToClient("heartbeat timeout", ErrorCode)
				return
			}
		case err := <-errChan:
			log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
			return
		}
	}
}

func StartServer() {
	log.Logger.Info("WsServer start")
	for {
		select {
		case price, ok := <-kucoin.PlgrPriceChan:
			if ok {
				Manager.Servers.Range(func(key, value interface{}) bool {
					value.(*Server).SendToClient(price, SuccessCode)
					return true
				})
			}
		}
	}
}
