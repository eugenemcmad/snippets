package tests

import (
	"fmt"
	"testing"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

func TestGoCache(t *testing.T) {
	var err error
	var lim = 10
	var key = "a"
	var val = "A"
	var ttl, cln = time.Second * 3, time.Second * 1

	var c = gocache.New(ttl, cln)
	err = c.Add(key, val, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[_DBG] ok")

	err = c.Add(key, val, 0)
	fmt.Println("[_DBG] ok, expected error:", err)
	if err == nil {
		t.Fatal()
	}

	for i := 0; i < lim; i++ {
		time.Sleep(ttl / 5)
		v, ok := c.Get(key)
		fmt.Println("[_DBG]", v, ok)

		switch {
		case i == 0 && !ok:
			t.Fatal()
		case i == 1:
			err = c.Add(key, val+val, 0)
			if err == nil {
				t.Fatal()
			}
			fmt.Println("[_DBG] ok, expected error:", err)
		case i == lim-1 && ok:
			t.Fatal()
		}
	}
	fmt.Println("[_DBG] ok")

	err = c.Add(key, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[_DBG] ok")
}
