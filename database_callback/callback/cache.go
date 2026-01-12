package callback

import (
	"sync"
	"time"
)

const (
	LineageCacheRefreshInterval = int64(3600)
)

type StageCache struct {
	stage           map[string]*CacheModel
	lock            *sync.RWMutex
	lastRefreshTime int64
}

type CacheModel struct {
	ID         string
	CreateTime int64
	Models     []any
}

func NewStageCache() *StageCache {
	return &StageCache{
		stage:           make(map[string]*CacheModel),
		lock:            &sync.RWMutex{},
		lastRefreshTime: time.Now().Unix(),
	}
}

func (c *StageCache) needRefresh() bool {
	return time.Now().Unix()-c.lastRefreshTime > LineageCacheRefreshInterval
}

func (c *StageCache) refresh() {
	now := time.Now().Unix()
	for _, s := range c.stage {
		if now-s.CreateTime > LineageCacheRefreshInterval {
			c.delete(s.ID)
		}
	}
}

func (c *StageCache) delete(id string) {
	c.lock.Lock()
	delete(c.stage, id)
	c.lock.Unlock()
}

func (c *StageCache) Set(id string, models []any) {
	c.lock.Lock()
	c.stage[id] = &CacheModel{
		ID:         id,
		CreateTime: time.Now().Unix(),
		Models:     models,
	}
	c.lock.Unlock()
}

func (c *StageCache) Cut(id string) []any {
	stageMode, ok := c.stage[id]
	if !ok {
		return nil
	}
	c.delete(id)
	if c.needRefresh() {
		go c.refresh()
	}
	return stageMode.Models
}
