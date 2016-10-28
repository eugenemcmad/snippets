package tests

import (
	"net/http"
	"testing"
	"time"
	"xr/configuration"
	xn "xr/nsq"
	nc "xr/nsq/config"
	"xr/track-server/actions"
	g "xr/xutor/global"

	"strconv"

	"github.com/bitly/go-nsq"
)

const (
	_TEST_PROF_ID int64 = 2852695508788684450 // Test value for profile id
	_TEST_SITE_ID int64 = 1                   // Test value for profile site id
)

func TestPutProfToNSQ(t *testing.T) {
	var err error
	var npp *nsq.Producer

	nsqCfg := getNsqCfg()
	npp, err = nsq.NewProducer(nsqCfg.DaemonAddr, nsq.NewConfig())
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 100; i++ {
		s := `apikey=os6peow476hmf53t&email=test` + strconv.Itoa(i%10) +
			`%40gmail.com&fname=test&token=68fffdf7dfcac7ce2c84092a3297fbff&ip=94.231.116.43&sourceurl=https%3A%2F%2Fdrivingtests.us&regdate=2016-10-24&utmsource=a&utmmedium=b&status=11`
		err = xn.Publish(npp, xn.GET_PROFILES_TOPIC, []byte(s))
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestFillTrackingEvents(t *testing.T) {
	var err error
	var nsqProdPtr *nsq.Producer
	var fullActions []actions.FullAction

	now := int32(time.Now().UTC().Unix()) // seconds

	for i := int64(0); i < 10; i++ {
		// Create test tracking action
		fa := actions.FullAction{
			Action: actions.Action{
				ActionId:     g.ActionRedirect,
				ContentId:    6,
				DomainGroup:  "Google",
				EmailId:      7,
				EmailRunTime: int64(now) + i,
				LinkType:     1,
				ProfileId:    _TEST_PROF_ID,
				TemplateId:   8,
				Timestamp:    int64(now) + i,
				SiteId:       _TEST_SITE_ID,
				XXHash:       9,
			},
			Headers: http.Header{"X-Forwarded-For": []string{"94.231.116.43"}},
		}
		fullActions = append(fullActions, fa)
	}

	nsqCfg := getNsqCfg()
	nsqProdPtr, err = nsq.NewProducer(nsqCfg.DaemonAddr, nsq.NewConfig())
	if err != nil {
		t.Error(err)
		return
	}

	for _, fa := range fullActions {
		var data []byte
		data, err = actions.SerializeFullAction(fa)
		if err != nil {
			t.Error(err)
			return
		}

		err = xn.Publish(nsqProdPtr, xn.TRACK_TOPIC, data)
		if err != nil {
			t.Error(err)
			return
		}
	}

	// export CONSUL_HTTP_ADDR=23.251.134.57:8500
	// printenv CONSUL_HTTP_ADDR
	// CGO_ENABLED=0 go test -v xr/snippets -run ^TestFillTrackingEvents$
}

func getNsqCfg() nc.NSQConfig {
	var err error
	var cfg nc.NSQConfig

	// Setup configuration
	err = configuration.Setup()
	if err != nil {
		panic(err)
	}

	// Fill configuration
	err = configuration.Fill(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
