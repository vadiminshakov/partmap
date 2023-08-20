package partmap

import "sync"

type partition struct {
	stor map[string]any
	sync.RWMutex
}

func (p *partition) set(key string, value any) {
	p.Lock()
	p.stor[key] = value
	p.Unlock()
}

func (p *partition) get(key string) (any, bool) {
	p.RLock()
	v, ok := p.stor[key]
	if !ok {
		p.RUnlock()
		return nil, false
	}
	p.RUnlock()
	return v, true
}

func (p *partition) del(key string) {
	p.Lock()
	delete(p.stor, key)
	p.Unlock()
}
