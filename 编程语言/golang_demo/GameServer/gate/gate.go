package gate

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lgynico/gameserver/gate/network/broadcaster"
	"github.com/lgynico/gameserver/gate/network/ws"
	"github.com/lgynico/gameserver/logger"
)

var (
	upgrader = &websocket.Upgrader{
		HandshakeTimeout: 3 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	sessionID atomic.Int32
)

func Run() {
	logger.Config("gateway")

	http.HandleFunc("/websocket", webscoketHandshake)
	if err := http.ListenAndServe("127.0.0.1:12345", nil); err != nil {
		panic(err)
	}
}

func webscoketHandshake(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}

	cliConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("websocket upgrade error: %v", err)
		return
	}

	logger.Info("Client incoming connection...")

	svrConn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:54321/websocket", nil)
	if err != nil {
		logger.Error("dial server error: %v", err)
		return
	}

	ctx := ws.NewWebsocketContext(cliConn, svrConn)
	ctx.SetSessionID(sessionID.Add(1))

	broadcaster.Add(&ctx)
	defer broadcaster.Remove(ctx.SessionID())

	ctx.Start()
}
