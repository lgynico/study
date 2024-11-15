package ws

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/lgynico/gameserver/logger"
	"google.golang.org/protobuf/proto"
)

const (
	rateLimteDuration = time.Second
	rateLimteCount    = 16
)

type WebsocketContext struct {
	userID    int64
	sessionID int32
	cliConn   *websocket.Conn
	svrConn   *websocket.Conn
	sendC     chan proto.Message
	exitC     chan struct{}
}

func NewWebsocketContext(cliConn, svrConn *websocket.Conn) WebsocketContext {
	return WebsocketContext{
		cliConn: cliConn,
		svrConn: svrConn,
		exitC:   make(chan struct{}),
	}
}

func (p *WebsocketContext) SetUserID(userID int64) {
	p.userID = userID
}

func (p *WebsocketContext) UserID() int64 {
	return p.userID
}

func (p *WebsocketContext) SetSessionID(sessionID int32) {
	p.sessionID = sessionID
}

func (p *WebsocketContext) SessionID() int32 {
	return p.sessionID
}

func (p *WebsocketContext) SendMessage(msg proto.Message) {
	if msg == nil {
		return
	}

	// innerMsg := pb.InnerMsg{
	// 	SessionID: p.sessionID,
	// 	UserID:    p.userID,
	// 	Data:      msg,
	// }

	p.sendC <- msg
}

func (p *WebsocketContext) SendCode(code int32) {
}

func (p *WebsocketContext) Start() {
	go p.writeLoop()
	p.readLoop()
}

func (p *WebsocketContext) Stop() {
}

func (p *WebsocketContext) writeLoop() {
	for {
		_, data, err := p.svrConn.ReadMessage()
		if err != nil {
			logger.Error("websocket read error: %v", err)
			break
		}

		if err := p.cliConn.WriteMessage(websocket.BinaryMessage, data); err != nil {
			logger.Error("websocket write error: %v", err)
			break
		}
	}
}

func (p *WebsocketContext) readLoop() {
	defer func() {
		p.cliConn.Close()
		p.svrConn.Close()
	}()

	p.cliConn.SetReadLimit(64 * 1024)

	rateCount := 0
	t := time.Now()

	for {
		_, data, err := p.cliConn.ReadMessage()
		if err != nil {
			logger.Error("websocket read error: %v", err)
			break
		}

		if time.Since(t) >= rateLimteDuration {
			t = time.Now()
			rateCount = 0
		}

		rateCount++
		if rateCount > rateLimteCount {
			logger.Error("rate limit exceeded")
			continue
		}

		if err = p.svrConn.WriteMessage(websocket.BinaryMessage, data); err != nil {
			logger.Error("websocket write error: %v", err)
			break
		}
	}

}
