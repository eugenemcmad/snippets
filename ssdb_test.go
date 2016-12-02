package tests

import (
	"fmt"
	"math/rand"
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
	for i := 0; i < 100; i++ {
		resIncrOk, resInt := ssdb.IncrBy(conn,
			fmt.Sprintf("%d_%04d_%d", utils.RandInt(100), utils.RandInt(100), utils.RandInt(100)),
			2)
		fmt.Print(err, resIncrOk, resInt)
	}
	fmt.Println(".................................")

	res := conn.Cmd("keys", "", "", 1000)
	fmt.Printf("%+v\n", *res)
	for _, d := range res.Data {
		fmt.Println(d)
	}
	fmt.Println(".................................")

	keys, err := ssdb.Keys(conn, "", "", 1000)
	fmt.Println(keys, err)
	fmt.Println(".................................")

	keys, err = ssdb.RKeys(conn, "", "", 1000)
	fmt.Println(keys, err)
	fmt.Println(".................................")

	keys, err = ssdb.Keys(conn, "0_0012_", "0_0013_", 1000)
	fmt.Println(keys, err)
	fmt.Println(".................................")

	keys, err = ssdb.Keys(conn, "0_0012_", "0_0013_", 1)
	fmt.Println(keys, err)
	fmt.Println(".................................")

	keys, err = ssdb.Keys(conn, "0_0012_", "0_0012_", 1000)
	fmt.Println(keys, err)
	fmt.Println(".................................")

}
