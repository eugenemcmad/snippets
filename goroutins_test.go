package tests

import (
	"fmt"
	"testing"
	"time"
)

var ()

func TestGoroutine(t *testing.T) {
	f1()

	time.Sleep(time.Second * 5)
	fmt.Println("\n TestGoroutine() exit")
}

func f1() {

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
