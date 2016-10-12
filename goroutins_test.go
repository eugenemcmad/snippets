package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var ()

func TestMutex(t *testing.T) {
	st := sync_t1{}

	for i := 0; i < 33; i++ {
		go fmt.Println(st.get_sync_f1(i))
	}

	time.Sleep(time.Second)
}

type sync_t1 struct {
	sync.Mutex
	m map[int]int
}

func (sc *sync_t1) get_sync_f1(n int) map[int]int {
	sc.Lock()
	fmt.Printf("%d lock \n", n)
	defer func() {
		sc.Unlock()
		fmt.Printf("%d unlock \n", n)
	}()

	sc.get_sync_f2(n)

	return sc.m
}
func (sc *sync_t1) get_sync_f2(n int) map[int]int {

	fmt.Printf("%d proc... \n", n)
	sc.m = make(map[int]int)
	for i := n; i < n+10; i++ {
		sc.m[i] = i
	}

	fmt.Printf("%d ready. \n", n)

	return sc.m
}

func TestGoroutine(t *testing.T) {
	gr_f1()

	time.Sleep(time.Second * 5)
	fmt.Println("\n TestGoroutine() exit")
}

func gr_f1() {

	go fmtForSleep(300, 10, "-")
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Print(".")
	}

	fmt.Println("\n f1() exit")
}

func fmtForSleep(max, slp int, s string) {
	for i := 0; i < max; i++ {
		time.Sleep(time.Duration(slp) * time.Millisecond)
		fmt.Print(s)
	}
	fmt.Println("\n fmtForSleep() end")
}
