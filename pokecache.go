package main

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type pokeCache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) pokeCache {
	pc := &pokeCache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	pc.reapLoop(interval)
	return *pc
}

func (p *pokeCache) add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.cache[key] = entry
}

func (p *pokeCache) Get(key string) (val []byte, success bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	res, ok := p.cache[key]
	if ok {
		return res.val, true
	} else {
		return nil, false
	}
}

func (p *pokeCache) Renew(key string) {
	res, _ := p.cache[key]
	res.createdAt = time.Now()
	p.cache[key] = res
}

func (p *pokeCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				p.cleanup(interval)
			}
		}
	}()
}

func (p *pokeCache) cleanup(expiry time.Duration) {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now()
	for key, val := range p.cache {
		if now.Sub(val.createdAt) > expiry {
			delete(p.cache, key)
			// fmt.Println("Deleted " + key)
		}
	}
}
