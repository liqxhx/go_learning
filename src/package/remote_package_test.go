package gopackage

// go get -u github.com/easierway/concurrent_map

import(
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage_ConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(8)
	m.Set(cm.StrKey("key"), 20)
	t.Log(m.Get(cm.StrKey("key")))
}