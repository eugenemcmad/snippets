package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {

	c := time.After(5 * time.Second)
	go sWatch1(c)
	for i := 0; i < 300; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}

	fmt.Println("\n exit")
}

func sWatch1(c <-chan time.Time) {

	select {
	case <-c:
		fmt.Println("\n timed out")
	}
}
