package dbsave

import (
	"time"

	"github.com/lgynico/gameserver/facade"
)

type LazyRecord struct {
	facade.LazySaver
	lastUpdateTime time.Time
}

func NewLazyRecord(data facade.LazySaver) *LazyRecord {
	return &LazyRecord{LazySaver: data, lastUpdateTime: time.Now()}
}
