package tests

import (
	"testing"
	"xr/configuration"
	xn "xr/nsq"
	nc "xr/nsq/config"

	"github.com/bitly/go-nsq"
)

func TestFillTrackingEvents(t *testing.T) {
	var err error
	var nsqProdPtr *nsq.Producer
	var data []byte

	nsqCfg := getNsqCfg()
	nsqProdPtr, err = nsq.NewProducer(nsqCfg.DaemonAddr, nsq.NewConfig())
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
