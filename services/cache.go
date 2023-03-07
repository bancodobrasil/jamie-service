package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/bancodobrasil/jamie-service/dtos"
)

// Cache ...
type Cache interface {
	Get(ctx context.Context, uuid string, version string) (*dtos.Menu, error)
	Put(ctx context.Context, uuid string, version string, value *dtos.Menu, ttl time.Duration) error
	Close() error
}

type cache struct {
	eMap *expiringMapSync
}

// NewCache ...
func NewCache() Cache {
	return &cache{
		eMap: newExpiringMapSync(),
	}
}

func (c *cache) buildkey(uuid string, version string) string {
	return fmt.Sprintf("%s-%s", uuid, version)
}

func (c *cache) Get(ctx context.Context, uuid string, version string) (*dtos.Menu, error) {
	v, ok := c.eMap.Get(c.buildkey(uuid, version))
	if !ok {
		return nil, nil
	}
	return v.(*dtos.Menu), nil
}

func (c *cache) Put(ctx context.Context, uuid string, version string, value *dtos.Menu, ttl time.Duration) error {
	c.eMap.Put(c.buildkey(uuid, version), value, ttl)
	return nil
}

func (c *cache) Close() error {
	c.eMap.StopCleanup()
	return nil
}

// Expiring Map NON SYNC

// type expiringMap struct {
// 	entries  map[string]*entry
// 	stopChan chan struct{}
// }

// type entry struct {
// 	value     *dtos.Menu
// 	expiresAt time.Time
// }

// func newExpiringMap() *expiringMap {
// 	m := &expiringMap{
// 		entries:  make(map[string]*entry),
// 		stopChan: make(chan struct{}),
// 	}
// 	go m.startCleanup()
// 	return m
// }

// func (m *expiringMap) Put(key string, value *dtos.Menu, expiration time.Duration) {
// 	e := &entry{
// 		value:     value,
// 		expiresAt: time.Now().Add(expiration),
// 	}
// 	m.entries[key] = e
// }

// func (m *expiringMap) Get(key string) (*dtos.Menu, bool) {
// 	e, ok := m.entries[key]
// 	if !ok {
// 		return nil, false
// 	}
// 	if time.Now().After(e.expiresAt) {
// 		delete(m.entries, key)
// 		return nil, false
// 	}
// 	return e.value, true
// }

// func (m *expiringMap) startCleanup() {
// 	ticker := time.NewTicker(time.Minute)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			for key, e := range m.entries {
// 				if time.Now().After(e.expiresAt) {
// 					delete(m.entries, key)
// 				}
// 			}
// 		case <-m.stopChan:
// 			return
// 		}
// 	}
// }

// func (m *expiringMap) StopCleanup() {
// 	close(m.stopChan)
// }

// Expiring Map SYNC

type expiringMapSync struct {
	m        sync.Map
	stopChan chan struct{}
}

type entry struct {
	value     interface{}
	expiresAt time.Time
}

func newExpiringMapSync() *expiringMapSync {
	m := &expiringMapSync{
		stopChan: make(chan struct{}),
	}
	go m.startCleanup()
	return m
}

// Put ...
func (m *expiringMapSync) Put(key, value interface{}, expiration time.Duration) {
	e := &entry{
		value:     value,
		expiresAt: time.Now().Add(expiration),
	}
	m.m.Store(key, e)
}

// Get ...
func (m *expiringMapSync) Get(key interface{}) (interface{}, bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return nil, false
	}
	e := v.(*entry)
	if time.Now().After(e.expiresAt) {
		m.m.Delete(key)
		return nil, false
	}
	return e.value, true
}

func (m *expiringMapSync) startCleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			m.m.Range(func(key, value interface{}) bool {
				e := value.(*entry)
				if time.Now().After(e.expiresAt) {
					m.m.Delete(key)
				}
				return true
			})
		case <-m.stopChan:
			return
		}
	}
}

// StopCleanup ...
func (m *expiringMapSync) StopCleanup() {
	close(m.stopChan)
}
