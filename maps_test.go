package tests

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsMapIsReference(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	m2 := m1
	m2[2] = 20
	fmt.Println(m1)
}

func TestErrSumMaps(t *testing.T) {
	pref := "ABOUT_ME:"
	e1 := errors.New("err1")
	e2 := errors.New("err2")
	e3 := errors.New("err3")

	//var es = make(errInfSum)
	//var es = errInfSum{}
	var es errInfSum

	es.save(pref)

	es.add(e1)
	es = errInfSum{}

	es.add(e1)
	es = newErrInfSum()

	es.add(e1)
	es.add(e2)
	es.add(e2)
	es.add(e3)
	es.add(e3)
	es.add(e3)
	es.save(pref)

	es.add(e2)
	es.add(e3)
	es.add(e3)
	es.save(pref)
}

// Error description stack with counter. No thread safe.
type errInfSum struct {
	m map[string]int
}

// Get new errors info summary
func newErrInfSum() errInfSum {
	return errInfSum{m: make(map[string]int)}
}

// Add error description.
func (es *errInfSum) add(e error) {
	if es.m == nil {
		es.m = make(map[string]int)
	}

	var s = e.Error()
	es.m[s] += 1
}

// Save to log and reset.
func (es *errInfSum) save(pref string) {
	if len(es.m) > 0 {
		var sum string
		for s, n := range es.m {
			sum += fmt.Sprintf("\n\t %d: %s", n, s)
		}
		fmt.Printf("%s errors summary:%v\n", pref, sum)
		es.m = make(map[string]int)
	}
}
