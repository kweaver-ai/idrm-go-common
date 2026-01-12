package data_lineage

import "time"

type CacheModel struct {
	ID         string
	CreateTime int64
	Models     []any
}

func (c *LineageCallback) needRefresh() bool {
	return time.Now().Unix()-c.lastRefreshTime > LineageCacheRefreshInterval
}

func (c *LineageCallback) refresh() {
	now := time.Now().Unix()
	for _, s := range c.stage {
		if now-s.CreateTime > LineageCacheRefreshInterval {
			c.delete(s.ID)
		}
	}
}

func (c *LineageCallback) delete(id string) {
	c.lock.Lock()
	delete(c.stage, id)
	c.lock.Unlock()
}

func (c *LineageCallback) Cut(id string) []any {
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
