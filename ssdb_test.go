package tests

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
	"xr/xutor/ssdb"
	"xr/xutor/utils"
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

func TestSsdbKeys(t *testing.T) {
	host := "104.155.41.252"
	conn, err := ssdb.GetConnector(host, 8888, 0, 0)
	if err != nil {
		fmt.Println("Connect Error:", err)
		return
	}
	defer conn.Close()
	rand.Seed(time.Now().Unix())
	resIncrOk, resInt := ssdb.IncrBy(conn, "key_"+strconv.Itoa(utils.RandInt(10))+"_"+strconv.Itoa(utils.RandInt(10)), "2")
	fmt.Printf("err:%v, incr:%v,%v\n", err, resIncrOk, resInt)

	fmt.Println(".................................")

	res := conn.Cmd("keys", "", "", 100)
	fmt.Printf("%+v\n", *res)
	for _, d := range res.Data {
		fmt.Println(d)
	}
	fmt.Println(".................................")

	keys, err := ssdb.Keys(conn, "", "", 100)
	fmt.Println(keys, err)
	fmt.Println(".................................")

	keys, err = ssdb.RKeys(conn, "", "", 100)
	fmt.Println(keys, err)
	fmt.Println(".................................")

}
