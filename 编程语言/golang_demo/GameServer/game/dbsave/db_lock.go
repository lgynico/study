package dbsave

import "sync"

var lockItem = struct{}{}

type dbLock struct {
	locks sync.Map
}

var DBLock = &dbLock{}

func (p *dbLock) TryLock(userID int64) bool {
	_, loaded := p.locks.LoadOrStore(userID, lockItem)
	return !loaded
}

func (p *dbLock) Unlock(userID int64) {
	p.locks.Delete(userID)
}

func (p *dbLock) IsLocked(userID int64) bool {
	_, ok := p.locks.Load(userID)
	return ok
}
