package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

// go get -u https://github.com/easierway/concurrent_map.git
// https://github.com/golang/dep
// brew install glide
// glide init
// glide install
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))

}
