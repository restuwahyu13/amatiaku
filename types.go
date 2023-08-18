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

	IntervalConfig struct {
		IntervalEveryRequest int
		IntervalNextRequest  int
		IntervalUnitType     string
	}

	IntervalCounterConfig struct {
		IntervalEveryRequest int
		IntervalNextRequest  int
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
		IntervalTarget     int
		IntervalUnitType   string
		Ticker             *time.Ticker
		Release            chan bool
		NowTriggerRequest  int
		NextTriggerRequest int
		StopTriggerRequest int
	}
)
