package tests

import (
	"fmt"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
)

// Checks existing record in memcache
func exists(rec string, mc *memcache.Client) bool {
	key := fmt.Sprintf("%v_%v_%v", rec)
	err := mc.Add(&memcache.Item{
		Key:        key,
		Value:      []byte(""),
		Expiration: 10, // sec
	})
	if err != nil {
		glog.Error(err)
		if strings.Contains(err.Error(), memcache.ErrNotStored.Error()) {
			return true
		}
	}
	return false

}

//// Create memcache connection
//mc := memcache.New(cfg.MemcacheURL)
