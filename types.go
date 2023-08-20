package main

import "time"

const (
	SECONDS = "second"
	MINUTES = "minutes"
	HOURS   = "hours"
	DAYS    = "days"
	WEEKS   = "weeks"
	MONTHS  = "months"
)

type (
	AmatiakuHandler interface {
		ByInterval(HandleFunc func())
		ByIntervalCounter(HandleFunc func())
	}

	NextRequest struct {
		Delay int
		Type  string
	}

	IntervalConfig struct {
		IntervalEveryRequest  int
		IntervalTargetRequest int
		IntervalNextRequest   NextRequest
		IntervalUnitType      string
	}

	IntervalCounterConfig struct {
		IntervalEveryRequest int
		IntervalNextRequest  NextRequest
		IntervalUnitType     string
		MaxCounterRequest    int
	}

	AmatiakuConfig struct {
		IntervalConfig        IntervalConfig
		IntervalCounterConfig IntervalCounterConfig
	}

	amatiaku struct {
		IntervalConfig        IntervalConfig
		IntervalCounterConfig IntervalCounterConfig
	}

	amatiakuRequest struct {
		Counter            int
		StopCounter        int
		TotalCounter       int
		Interval           int
		IntervalTarget     int
		IntervalUnitType   string
		Ticker             *time.Ticker
		Release            chan bool
		NowTriggerRequest  int
		NextTriggerRequest int
		StopTriggerRequest int
	}
)
