package tests

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRnd1(t *testing.T) {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {

		jj := rndInt(ii)
		fmt.Println(jj)
	}

}

func rndInt(ii []int) []int {
	jj := make([]int, len(ii))
	xx := rand.Perm(len(ii))
	for i, x := range xx {
		jj[i] = ii[x]
	}
	return jj
}
