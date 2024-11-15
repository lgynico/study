package dbsave

import (
	"time"

	"github.com/lgynico/gameserver/facade"
	"github.com/lgynico/gameserver/logger"
)

type dbSaver struct {
	records map[string]*LazyRecord
	saveC   chan facade.LazySaver
	deleteC chan string
}

var DBSaver = dbSaver{
	records: make(map[string]*LazyRecord, 1024),
	saveC:   make(chan facade.LazySaver, 1024),
	deleteC: make(chan string, 16),
}

func (p *dbSaver) Start() {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			p.save()
		case saver := <-p.saveC:
			if record, ok := p.records[saver.Key()]; ok {
				record.lastUpdateTime = time.Now()
			} else {
				p.records[saver.Key()] = NewLazyRecord(saver)
			}
		case key := <-p.deleteC:
			delete(p.records, key)
		}

	}
}

func (p *dbSaver) Save(saver facade.LazySaver) {
	if saver == nil {
		return
	}

	p.saveC <- saver
}

func (p *dbSaver) GetSaver(key string) (facade.LazySaver, bool) {
	record, ok := p.records[key]
	if !ok {
		return nil, false
	}

	return record.LazySaver, true
}

func (p *dbSaver) Discard(key string) {
	p.deleteC <- key
}

func (p *dbSaver) save() {
	var (
		now     = time.Now()
		removes = make([]string, 0, 64)
	)

	for key, record := range p.records {
		if now.Sub(record.lastUpdateTime) >= 30*time.Second {
			record.Save()
			logger.Info("lazy save: %s", record.Key())
			removes = append(removes, key)
		}
	}

	for _, key := range removes {
		delete(p.records, key)
	}
}
