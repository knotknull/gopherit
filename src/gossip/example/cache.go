package main

import (
	"sync"
	"time"
)

type Value struct {
	Value       string    `json:"value"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]Value
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]Value),
	}
}

func (c *Cache) Get(key string) Value {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Cache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = Value{Value: value, LastUpdated: time.Now()}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) Clone() map[string]string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[string]string)
	for k, v := range c.data {
		m[k] = v.Value
	}
	return m
}
