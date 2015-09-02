package metric

import (
	"errors"
	"time"

	"github.com/bigdatadev/goryman"
	"github.com/pivotal-golang/lager"
)

type eventEmission struct {
	event  goryman.Event
	logger lager.Logger
}

var riemannClient *goryman.GorymanClient
var eventHost string
var eventTags []string

var clientConnected bool
var emissions = make(chan eventEmission, 1000)

var errQueueFull = errors.New("event queue is full")

func Initialize(riemannAddr string, host string, tags []string) {
	client := goryman.NewGorymanClient(riemannAddr)

	riemannClient = client
	eventHost = host
	eventTags = tags

	go emitLoop()
}

func emit(emission eventEmission) {
	emission.logger.Debug("emit")

	if riemannClient == nil {
		return
	}

	emission.event.Host = eventHost
	emission.event.Time = time.Now().Unix()
	emission.event.Tags = append(emission.event.Tags, eventTags...)

	select {
	case emissions <- emission:
	default:
		emission.logger.Error("queue-full", nil)
	}
}

func emitLoop() {
	for emission := range emissions {
		if !clientConnected {
			err := riemannClient.Connect()
			if err != nil {
				emission.logger.Error("connection-failed", err)
				continue
			}

			clientConnected = true
		}

		err := riemannClient.SendEvent(&emission.event)
		if err != nil {
			emission.logger.Error("failed-to-emit", err)

			if err := riemannClient.Close(); err != nil {
				emission.logger.Error("failed-to-close", err)
			}

			clientConnected = false
		}
	}
}
