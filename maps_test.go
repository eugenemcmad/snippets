package tests

import (
	"errors"
	"fmt"
	"sort"
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

func TestInfSum(t *testing.T) {

	var ais addInfSum

	for i := 10; i > 0; i-- {
		ais.put(fmt.Sprintf("%v Aol %v", i, i))
	}
	fmt.Println(ais)
	ais.save()
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

// Added decz summary
type addInfSum struct {
	sl []string
}

// Put part of decz to summary
// args:
//      -s: stat part description
func (o *addInfSum) put(s string) {
	if o.sl == nil {
		o.sl = []string{}
	}
	o.sl = append(o.sl, s)
}

// Save to log
func (o *addInfSum) save() {

	sort.Strings(o.sl)

	if len(o.sl) > 0 {
		var sum string
		for _, s := range o.sl {
			sum += fmt.Sprintf("\n\t %s", s)
		}
		fmt.Printf("Add decz summary:%v\n", sum)
	}
}
