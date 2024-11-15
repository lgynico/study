package ws

import (
	"encoding/binary"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lgynico/gameserver/game/handler"
	"github.com/lgynico/gameserver/logger"
	"github.com/lgynico/gameserver/pb"
	"github.com/lgynico/gameserver/thread"
	"google.golang.org/protobuf/proto"
)

const (
	rateLimteDuration = time.Second
	rateLimteCount    = 16
)

type WebsocketContext struct {
	userID    int64
	sessionID int32
	conn      *websocket.Conn
	sendC     chan proto.Message
	exitC     chan struct{}
}

func NewWebsocketContext(conn *websocket.Conn) WebsocketContext {
	return WebsocketContext{
		conn:  conn,
		exitC: make(chan struct{}),
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
	p.sendC = make(chan proto.Message, 64)

	for {
		msg, ok := <-p.sendC
		if !ok {
			continue
		}

		data, err := pb.Encode(msg)
		if err != nil {
			logger.Error("encode message error: cmd=%s, err=%v", msg.ProtoReflect().Descriptor().Name(), err)
		}

		if err := p.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
			logger.Error("websocket write error: %v", err)
		}
	}
}

func (p *WebsocketContext) readLoop() {
	defer handler.OnUserQuit(p)

	p.conn.SetReadLimit(64 * 1024)

	rateCount := 0
	t := time.Now()

	for {
		_, data, err := p.conn.ReadMessage()
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

		if len(data) < 4 {
			logger.Error("message too short")
			continue
		}

		code := binary.BigEndian.Uint16(data[2:4])
		cmd, err := pb.Decode(data[4:], int32(code))
		if err != nil {
			logger.Error("decode message error: code=%d, err=%v", code, err)
			continue
		}

		logger.Info("Receive message: %+v", cmd.ProtoReflect().Descriptor().Name())

		handleFunc, ok := handler.GetHandler(int32(code))
		if !ok {
			logger.Error("handler not found: code=%d", code)
			continue
		}

		thread.Logic.Submit(func() {
			if err := handleFunc(p, cmd); err != nil {
				logger.Error("handler error: code=%d, err=%v", code, err)
			}
		})
	}

}
