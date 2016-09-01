package tests

import (
	"fmt"
	"testing"
	"xr/xutor/aerospikedb"
	g "xr/xutor/global"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/golang/glog"
)

var (
	pid      = int64(10101010)
	dayStart = int64(1234567890)
	slots    = map[int64]int{1000000000: 10, 2000000000: 20}
)

func TestCalendarUpsert_PutObject(t *testing.T) {
	acp, err := aerospikedb.GetClient(g.AsHostsTestX)
	if err != nil {
		t.Errorf("aerospikedb.GetClient(%v) %v \n", g.AsHostsTestX, err)
	}

	cal, err := Get(acp, pid, dayStart)
	if err != nil {
		fmt.Printf("Get() err: %v \n", err)
	}

	fmt.Printf("[_DBG] Get() result: %+v \n", cal)

	cal.slots = slots

	err = cal.Upsert(acp)
	if err != nil {
		fmt.Printf("cal.Upsert() err: %v \n", err)
	}
}

// Return Calendar structure from DB or new instance of Calendar (with dayStart set to today), if there is not such record.
func Get(client *as.Client, pid int64, startOfDay int64) (*Calendar, error) {
	sk := fmt.Sprintf("%d_%d", pid, startOfDay)
	key, err := as.NewKey(g.ASPNameSpace, g.ASPCalendarsSet, sk)
	if err != nil {
		return nil, err
	}

	ok, err := client.Exists(nil, key)
	if err != nil {
		return nil, err
	}

	if !ok {
		fmt.Println("[_DBG] cal not exist - create new")
		return &Calendar{
			pid:      pid,
			dayStart: startOfDay,
			slots:    make(map[int64]int),
		}, nil
	} else {
		fmt.Println("[_DBG] cal in Get() exist ok")

	}

	rec, err := client.Get(nil, key)
	if err != nil {
		fmt.Printf("client.Get(%s.%s.%v) err: %v \n", g.ASPNameSpace, g.ASPCalendarsSet, sk, err)
	}
	fmt.Printf("client.Get(%s.%s.%v) rec: %v \n", g.ASPNameSpace, g.ASPCalendarsSet, sk, rec)

	c, err := TryNewCalendar(rec.Bins)
	if err != nil {
		fmt.Printf("TryNewCalendar(%v) err: %v \n", rec.Bins, err)
	}

	return c, err
}

func TryNewCalendar(bins as.BinMap) (*Calendar, error) {
	c := Calendar{slots: make(map[int64]int)}

	switch n := bins[g.PidKeyNamePtrn].(type) {
	case int64:
		c.pid = n
	case int:
		c.pid = int64(n)
	default:
		return nil, fmt.Errorf("'%s' unexpected type %T", g.PidKeyNamePtrn, n)
	}

	switch n := bins[g.DayStartKeyNamePtrn].(type) {
	case int64:
		c.dayStart = n
	case int:
		c.dayStart = int64(n)
	default:
		return nil, fmt.Errorf("'%s' unexpected type %T", g.DayStartKeyNamePtrn, n)
	}

	m, ok := bins[g.SlotsKeyNamePtrn].(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("'%s' unexpected type", g.SlotsKeyNamePtrn)
	}
	for k, v := range m {
		var ts int64
		var eid int
		switch n := k.(type) {
		case int64:
			ts = n
		case int:
			ts = int64(n)
		default:
			return nil, fmt.Errorf("'%s' unexpected key type %T", g.SlotsKeyNamePtrn, n)
		}

		switch n := v.(type) {
		case int:
			eid = n
		default:
			return nil, fmt.Errorf("'%s' unexpected value type %T", g.SlotsKeyNamePtrn, n)
		}

		c.slots[ts] = eid
	}

	return &c, nil
}

// Update record in DB or insert if there is not such record.
func (c *Calendar) Upsert(client *as.Client) error {
	sk := fmt.Sprintf("%d_%d", c.pid, c.dayStart)
	glog.V(10).Infof("try Upsert(%s.%s.%s): %+v \n", g.ASPNameSpace, g.ASPCalendarsSet, sk, *c)
	key, err := as.NewKey(g.ASPNameSpace, g.ASPCalendarsSet, sk)
	if err != nil {
		return fmt.Errorf("as.NewKey(%s.%s.%s) err: %v", g.ASPNameSpace, g.ASPCalendarsSet, sk, err)
	}

	ok, err := client.Exists(nil, key)
	if err != nil {
		return fmt.Errorf("client.Exists() err: %v", err)
	}
	fmt.Printf("[_DBG] exist: %v \n", ok)

	err = client.PutBins(nil, key,
		as.NewBin("pid", c.pid),
		as.NewBin("day_start", c.dayStart),
		as.NewBin("slots", c.slots))
	if err != nil {
		return fmt.Errorf("client.PutBins(%+v) err: %v \n", *c, err)
	}

	/*err = client.PutObject(nil, key, c)
	if err != nil {
		return fmt.Errorf("client.PutObject(%+v) err: %v \n", *c, err)
	}*/

	return nil
}

// Сигнатуры типа Calendar и его методов.
type Calendar struct {
	pid      int64         `as:pid`       // Profile ID.
	dayStart int64         `as:day_start` // Start of a day by profile's timezone in UTC.
	slots    map[int64]int `as:slots`     // 15-minutes send slot for a day of a calendar.
}

func TestExecuteCleanHBUDF(t *testing.T) {
	hosts := []*as.Host{as.NewHost("127.0.0.1", g.AS_DEF_PORT)}
	cli, err := aerospikedb.GetClient(hosts)
	if err != nil {
		t.Errorf("aerospikedb.GetClient(%v) %v \n", hosts, err)
		return
	}

	pids := []int64{
		2773078823355393185,
		4247475075215243666,
		4032530127856820284,
		5652457131391091506,
	}

	for _, pid := range pids {

		asKey, err := as.NewKey(g.ASPNameSpace, g.ASPProfilesSet, pid)
		if err != nil {
			fmt.Println(err)
			return
		}

		execRes, err := cli.Execute(nil, asKey, "profiles", "cln_hb")
		fmt.Printf("pid: %d, %v\n", pid, execRes)
		if err != nil {
			fmt.Printf("Execute result=%v, error=%v\n", execRes, err)
			return
		}
	}
}
