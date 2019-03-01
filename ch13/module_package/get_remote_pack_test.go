package module

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
	"github.com/easierway/fast_hashmap"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

func TestFastMap(t *testing.T) {
	m := fast_hashmap.NewFasterMap(100)
	t.Log(m)
}
