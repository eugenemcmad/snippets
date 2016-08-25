package tests

import (
	"fmt"
	"testing"
	"time"
)

var (
	_test_stop_signal bool
)

func TestAfter(t *testing.T) {

	max := 300
	x := make(chan bool, 10)
	d := 10 * time.Second
	go breaker(x, d)
	for i := 0; i < max; i++ {
		time.Sleep(30 * time.Millisecond)
		fmt.Print(".")
		if _test_stop_signal {
			fmt.Println("\n TestAfter() got '_test_stop_signal'")
			return
		}
		if i == max/2 {
			x <- true
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Println("\n TestAfter() exit")
}

func breaker(x chan bool, d time.Duration) {

	select {

	case <-time.After(d):
		_test_stop_signal = true
		fmt.Println("\n breaker() timed out")
		return

	case <-x:
		fmt.Println("\n breaker() exit ok")
		return
	}
}

func f2() {
	c := time.After(5 * time.Second)
	go sWatch1(c)
	for i := 0; i < 30; i++ {
		time.Sleep(30 * time.Millisecond)
		fmt.Print(",")
		if _test_stop_signal {
			fmt.Println("\n _test_stop_signal")
			return
		}
	}
}

func sWatch1(c <-chan time.Time) {

	select {
	case <-c:
		_test_stop_signal = true
		fmt.Println("\n timed out")
	}
}
