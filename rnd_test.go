package tests

import (
	"fmt"
	"testing"

	"xr/xutor/utils"
)

func TestRndByRanges(t *testing.T) {
	conns := []uint32{1, 2, 7}
	var m = map[uint32]int{}

	var totalSumOfParts uint32
	for _, dc := range conns {
		totalSumOfParts += dc
	}

	for i := 0; i < 100; i++ {
		var pcTotal uint32
		var rnd = uint32(utils.RandInt(int(totalSumOfParts)))
		for _, dc := range conns {
			pcTotal += dc
			if rnd < pcTotal {
				m[dc] += 1
				break
			}

		}
	}

	fmt.Printf("%v \n", m)
}
