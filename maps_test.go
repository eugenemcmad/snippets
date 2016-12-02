package tests

import (
	"fmt"
	"testing"
)

func TestIsMapIsReference(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	m2 := m1
	m2[2] = 20
	fmt.Println(m1)
}
