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

// Test gnome sort algorithm
func TestGnomeSortAsc(t *testing.T) {
	ii := []int{0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 50, 100, -1, -10, -100, -1000}

	rand.Seed(time.Now().UnixNano())
	a := rndInt(ii)
	fmt.Println(a)
	a = sortAsc(a)
	fmt.Println(a)
}

// Test gnome sort algorithm
func TestGnomeSortDsc(t *testing.T) {
	ii := []int{0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 50, 100, -1, -10, -100, -1000}

	rand.Seed(time.Now().UnixNano())
	a := rndInt(ii)
	fmt.Println(a)
	a = sortDsc(a)
	fmt.Println(a)
}

func sortAsc(a []int) []int {
	i := 1
	for i < len(a) {
		if i == 0 || a[i-1] <= a[i] {
			i++
		} else {
			temp := a[i]
			a[i] = a[i-1]
			a[i-1] = temp
			i--
		}
	}
	return a
}

func sortDsc(a []int) []int {
	i := 1
	for i < len(a) {
		if i == 0 || a[i-1] >= a[i] {
			i++
		} else {
			temp := a[i]
			a[i] = a[i-1]
			a[i-1] = temp
			i--
		}
	}
	return a
}

func rndInt(ii []int) []int {
	jj := make([]int, len(ii))
	xx := rand.Perm(len(ii))
	for i, x := range xx {
		jj[i] = ii[x]
	}
	return jj
}
