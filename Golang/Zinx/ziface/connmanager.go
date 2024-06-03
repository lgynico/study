package ziface

type ConnManager interface {
	Add(conn Connection)
	Remove(conn Connection)
	Get(connID int32) (Connection, bool)
	Len() int
	CleanConn()
}
