package game

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lgynico/gameserver/db"
	"github.com/lgynico/gameserver/game/dbsave"
	"github.com/lgynico/gameserver/game/network/broadcaster"
	"github.com/lgynico/gameserver/game/network/ws"
	"github.com/lgynico/gameserver/logger"
	"github.com/lgynico/gameserver/pb"
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
	logger.Config("game_server")
	pb.Init()
	db.Init()

	go dbsave.DBSaver.Start()

	http.HandleFunc("/websocket", webscoketHandshake)
	if err := http.ListenAndServe("127.0.0.1:54321", nil); err != nil {
		panic(err)
	}
}

func webscoketHandshake(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("websocket upgrade error: %v", err)
		return
	}

	logger.Info("Client incoming connection...")

	ctx := ws.NewWebsocketContext(conn)
	ctx.SetSessionID(sessionID.Add(1))

	broadcaster.Add(&ctx)
	defer broadcaster.Remove(ctx.SessionID())

	ctx.Start()
}
