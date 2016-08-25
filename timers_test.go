package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {

	c := time.After(5 * time.Second)
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}

	fmt.Println("\n end")

	select {
	case <-c:
		fmt.Println("\n timed out")
	}

	fmt.Println("\n exit")
}
