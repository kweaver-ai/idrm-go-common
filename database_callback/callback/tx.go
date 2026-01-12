package callback

import "sync"

type txGroups map[string]txGroup

// txLock 同一个事务
type txGroup struct {
	chs  []chan any
	lock *sync.RWMutex
}

func NewTxGroups() txGroups {
	return make(map[string]txGroup)
}

func (d *DatabaseCallback) InGroup(id string, c chan any) *sync.RWMutex {
	ch, ok := d.txGroups[id]
	if !ok {
		ch = txGroup{
			chs:  make([]chan any, 0),
			lock: &sync.RWMutex{},
		}
	}
	ch.chs = append(ch.chs, c)
	d.txGroups[id] = ch
	return ch.lock
}

// CloseGroup  关闭组
func (d *DatabaseCallback) CloseGroup(id string) {
	if _, ok := d.txGroups[id]; !ok {
		return
	}
	delete(d.txGroups, id)
}

// ReleaseGroup 小组内有ch有结果了，其他的也跟着有结果了
func (d *DatabaseCallback) ReleaseGroup(id string) {
	ch, ok := d.txGroups[id]
	if !ok {
		return
	}
	for _, c := range ch.chs {
		c <- 1
	}
}
