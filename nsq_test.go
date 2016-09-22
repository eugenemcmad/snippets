package tests

import (
	"testing"
	"xr/configuration"
	xn "xr/nsq"
	nc "xr/nsq/config"

	"time"
	"xr/track-server/actions"
	g "xr/xutor/global"

	"net/http"

	"github.com/bitly/go-nsq"
)

const (
	_TEST_PROF_ID int64 = 2852695508788684450 // Test value for profile id
	_TEST_SITE_ID int64 = 1                   // Test value for profile site id
)

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
