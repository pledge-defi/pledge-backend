package ws

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"os/signal"
	"pledge-backend/log"
	"sync"
	"time"
)

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

var Manager = ServerManager{}
var UserPingPongDurTime int64 = 20 // seconds
var ServerChain = make(chan *Server)

func (s *Server) SendToClient(msg []byte) {
	s.Lock()
	defer s.Unlock()
	err := s.Socket.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Logger.Sugar().Error(s.Id+" SendToClient err ", err)
	}
}

func (s *Server) ReadAndWrite() {
	defer func() {
		recover := recover()
		if recover != nil {
			fmt.Println(s.Id, "read recover")
		}
		s.Socket.Close()
		Manager.Servers.Delete(s)
	}()
	//write
	go func() {
		for {
			select {
			case message, ok := <-s.Send:
				if !ok {
					Manager.Servers.Delete(s)
					return
				}
				_, has := Manager.Servers.Load(s)
				if has {
					s.SendToClient(message)
				} else {
					Manager.Servers.Delete(s)
					return
				}
			case <-time.After(time.Second):
				_, has := Manager.Servers.Load(s)
				if !has {
					Manager.Servers.Delete(s)
					return
				}
			}
		}
	}()
	//read
	go func() {
		for {
			_, message, err := s.Socket.ReadMessage()
			if err != nil {
				return
			}
			//update ping time
			if string(message) == "ping" || string(message) == `"ping"` || string(message) == "'ping'" {
				s.LastTime = time.Now().Unix()
				//写不进跳过
				_, has := Manager.Servers.Load(s)
				if has {
					s.Send <- []byte("pong")
				} else {
					Manager.Servers.Delete(s)
					return
				}
			}
			continue
		}
	}()
	//check ping pong
	for {
		select {
		case <-time.After(time.Second):
			if time.Now().Unix()-s.LastTime >= UserPingPongDurTime {
				log.Logger.Info(s.Id + " timeout")
				return
			}
			_, has := Manager.Servers.Load(s)
			if !has {
				return
			}
		}
	}
}

func StartServer() {
	log.Logger.Info("WsServer start")
	for {
		select {
		case conn, ok := <-ServerChain:
			if ok {
				go conn.ReadAndWrite()
			}
		}
	}
}

var addr = flag.String("addr", "https://openapi-sandbox.kucoin.com/api/v1/bullet-public", "http service address")

func GetExchangeData() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Logger.Sugar().Infof("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Logger.Sugar().Error("dial:", err)
		return
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Logger.Sugar().Info("read:", err)
				return
			}
			log.Logger.Sugar().Info("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Logger.Sugar().Error("write:", err)
				return
			}
		case <-interrupt:
			log.Logger.Info("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Logger.Sugar().Error("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
