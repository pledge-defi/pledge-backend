package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
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

	errChian := make(chan error)

	Manager.Servers.Store(s.Id, s)

	//write
	go func() {
		for {
			select {
			case message, ok := <-s.Send:
				if !ok {
					Manager.Servers.Delete(s)
					errChian <- errors.New("write message error")
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
				errChian <- err
				return
			}
			//update ping time
			if string(message) == "ping" || string(message) == `"ping"` || string(message) == "'ping'" {
				s.LastTime = time.Now().Unix()
				s.Send <- []byte("pong")
			}
			continue
		}
	}()

	//check ping pong
	for {
		select {
		case <-time.After(time.Second):
			if time.Now().Unix()-s.LastTime >= UserPingPongDurTime {
				s.SendToClient("ping timeout", ErrorCode)
				log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned timeout")
				Manager.Servers.Delete(s)
				s.Socket.Close()
				close(s.Send)
				return
			}
		case err := <-errChian:
			log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
			return

		}
	}
}

func StartServer() {
	log.Logger.Info("WsServer start")
	for {
		select {
		case price, ok := <-kucoin.PlgrPriceChain:
			fmt.Println(price, ok, 999888)
			if ok {
				Manager.Servers.Range(func(key, value interface{}) bool {
					fmt.Println(key.(string))
					value.(*Server).SendToClient(price, SuccessCode)
					return true
				})
			}
		}
	}
}
