package tests

import (
	"fmt"
	"testing"
	"xr/xutor/ssdb"
)

func TestSsdbFlushCmd(t *testing.T) {
	host := "104.155.41.252"
	conn, err := ssdb.GetConnector(host, 8895, 0, 0)
	if err != nil {
		fmt.Println("Connect Error:", err)
		return
	}
	defer conn.Close()

	resIncrOk, resInt := ssdb.IncrBy(conn, "key2", "2")
	fmt.Printf("err:%v, incr:%v,%v\n", err, resIncrOk, resInt)

	fmt.Printf("%v", *conn.Cmd("info"))

	fmt.Println(".................................")
	fr0 := conn.Cmd("flushdb", "yes")
	fmt.Printf("%v\n", *fr0)
	fmt.Println(".................................")

	batch := conn.Batch()
	batch.Cmd("flushdb")
	batch.Cmd("yes")
	bfr, err := batch.Exec()

	fmt.Printf("%v; %v\n", bfr, err)
	for _, r := range bfr {
		fmt.Printf("%v\n", *r)
	}

	fmt.Println(".................................")
}

func TestSsdbFlushCln(t *testing.T) {
	host := "104.155.41.252"
	conn, err := ssdb.GetConnector(host, 8888, 0, 0)
	if err != nil {
		fmt.Println("Connect Error:", err)
		return
	}
	defer conn.Close()

	resIncrOk, resInt := ssdb.IncrBy(conn, "key2", "2")
	fmt.Printf("err:%v, incr:%v,%v\n", err, resIncrOk, resInt)

	//fmt.Printf("%v", *conn.Cmd("info"))

	fmt.Println(".................................")

	cnt, ecr := ssdb.ClearAll(conn)
	fmt.Printf("clearAll() total: %d,  err:%v\n", cnt, ecr)
	fmt.Println(".................................")

}

/*
// ClearAll clear all cached in memcache.
func (rc *Cache) ClearAll() error {
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return err
		}
	}
	keyStart, keyEnd, limit := "", "", 50
	resp, err := rc.Scan(keyStart, keyEnd, limit)
	for err == nil {
		size := len(resp)
		if size == 1 {
			return nil
		}
		keys := []string{}
		for i := 1; i < size; i += 2 {
			keys = append(keys, string(resp[i]))
		}
		_, e := rc.conn.Do("multi_del", keys)
		if e != nil {
			return e
		}
		keyStart = resp[size-2]
		resp, err = rc.Scan(keyStart, keyEnd, limit)
	}
	return err
}*/
