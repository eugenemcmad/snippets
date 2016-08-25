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

	go f2()
	for i := 0; i < 300; i++ {
		time.Sleep(30 * time.Millisecond)
		fmt.Print(".")
	}

	fmt.Println("\n exit")
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
