package tests

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	"encoding/json"

	_ "github.com/lib/pq"
)

const _CLEANED_SQL_TMPL string = `
SET max_statement_mem='3GB'; SET statement_mem='2GB';

SELECT day,domain_group,partner_id,source_id,ready_for_act,in_act,not_act,acted
FROM cleaning_stat
WHERE day >= %d AND day <= %d
ORDER BY day;
`

var (
	pgDB *sql.DB
)

func TestPGQuery(t *testing.T) {
	setup()
	now := time.Now()
	ss := strconv.FormatInt(now.Add(-time.Hour*24 /*yesterday .. now*/).Unix(), 10)
	se := strconv.FormatInt(now.Unix(), 10)
	figak(ss, se)
}

func figak(start, end string) {

	sdt, err := strconv.ParseInt(start, 10, 0)
	if err != nil {
		panic(err)
	}

	// use dates discrete by day (time.Unix create 'local' time)
	sd := time.Unix(sdt, 0).Truncate(time.Hour * 24)
	pd := sd.Add(-time.Hour * 24).Unix()

	ed, err := strconv.ParseInt(end, 10, 0)
	if err != nil {
		panic(err)
	}

	sqlReq := fmt.Sprintf(_CLEANED_SQL_TMPL, pd, ed)
	rows, err := pgDB.Query(sqlReq)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	m := map[CleanedCounterKey]CleaningStat{}
	for rows.Next() {
		k := CleanedCounterKey{}
		v := CleaningStat{}
		err := rows.Scan(
			&k.Timestamp, &k.Edg, &k.PartnerId, &k.SourceId,
			&v.ReadyForActivation, &v.InActivation, &v.NotActivated, &v.Activated)
		//fmt.Printf("%+v: %+v\n", k, v)
		if err != nil {
			panic(err)
		}
		m[k] = v
	}
	sum := comp(m, sd.Unix())
	fmt.Printf("sum:\n%v\n", sum)
	bytes, err := json.Marshal(sum)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bytes:\n%s\n", bytes)
}

func comp(m map[CleanedCounterKey]CleaningStat, start int64) []CleanedCounterSummary {
	fmt.Printf("start: %d\n", start)
	sum := []CleanedCounterSummary{}
	for k, v := range m {
		r := CleanedCounterSummary{}
		if k.Timestamp < start {
			continue
		}
		r.CleanedCounterKey = k
		r.ReadyForActivation = v.ReadyForActivation
		r.InActivation = v.InActivation
		r.NotActivated = v.NotActivated
		r.Activated = v.Activated

		k0 := CleanedCounterKey{
			Timestamp: k.Timestamp - 60*60*24,
			Edg:       k.Edg,
			PartnerId: k.PartnerId,
			SourceId:  k.SourceId,
		}
		v0, ok := m[k0]
		if ok {
			r.DReadyForActivation = v.ReadyForActivation - v0.ReadyForActivation
			r.DNotActivated = v.NotActivated - v0.NotActivated
			r.DActivated = v.Activated - v0.Activated
			r.DInActivation = v.InActivation - v0.InActivation
		} else {
			r.DReadyForActivation = v.ReadyForActivation
			r.DNotActivated = v.NotActivated
			r.DActivated = v.Activated
			r.DInActivation = v.InActivation
		}

		sum = append(sum, r)
		fmt.Printf("r: %+v\n", r)

	}

	return sum
}

func setup() {
	var err error
	// Create database connection
	dbInfo := "host=138.201.30.69 port=5432 user=gpadmin dbname=xr-stats sslmode=disable"
	pgDB, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	err = pgDB.Ping()
	if err != nil {
		panic(err)
	}
}

// Cleaning statistic multi key
type CleanedCounterKey struct {
	Timestamp int64  `json:"ts"`  // process timestamp
	Edg       string `json:"edg"` // email domain group
	PartnerId int    `json:"prt"` // first partner id
	SourceId  int    `json:"src"` // first source id
}

// Cleaning statistic struct
type CleaningStat struct {
	ReadyForActivation int `json:"rdy"`
	InActivation       int `json:"in"`
	NotActivated       int `json:"not"`
	Activated          int `json:"act"`
}

// Cleaning statistics summary
type CleanedCounterSummary struct {
	CleanedCounterKey
	CleaningStat
	DReadyForActivation int `json:"drdy"`
	DNotActivated       int `json:"dnot"`
	DActivated          int `json:"dact"`
	DInActivation       int `json:"din"`
}
