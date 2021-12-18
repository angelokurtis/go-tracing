package tracing

import "sync"

var cache = struct {
	sync.RWMutex
	data map[string]string
}{data: make(map[string]string)}

func set(key, value string) {
	cache.Lock()
	cache.data[key] = value
	cache.Unlock()
}
