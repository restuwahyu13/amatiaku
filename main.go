package main

import (
	"fmt"
	"log"
	"time"

	"github.com/restuwahyu13/amatiaku/helpers"
)

func main() {
	amatiaku := NewAmatiAku(&AmatiakuConfig{
		IntervalCounterConfig: IntervalCounterConfig{
			IntervalEveryRequest: 5,
			IntervalUnitType:     SECONDS,
			IntervalNextRequest: NextRequest{
				Delay: 1,
				Type:  MINUTES,
			},
			MaxCounterRequest: 30,
		},
	})

	amatiaku.ByIntervalCounter(func() {
		fmt.Println("Hello world")
	})
}

func NewAmatiAku(config *AmatiakuConfig) AmatiakuHandler {
	return &amatiaku{IntervalConfig: config.IntervalConfig, IntervalCounterConfig: config.IntervalCounterConfig}
}

/**
================================================================
==  WATCHDOG BY INTERVAL COUNTER
================================================================
**/

func (h *amatiaku) ByIntervalCounter(HandleFunc func()) {
	var req amatiakuRequest = amatiakuRequest{}

	req.Ticker = time.NewTicker(time.Second * 1)
	req.Release = make(chan bool, 1)

	req.IntervalTarget = h.IntervalCounterConfig.IntervalEveryRequest
	req.IntervalUnitType = h.IntervalCounterConfig.IntervalUnitType

	req.NowTriggerRequest = req.IntervalTarget
	req.NextTriggerRequest = (req.NowTriggerRequest + req.IntervalTarget)

	switch h.IntervalCounterConfig.IntervalNextRequest.Type {
	case SECONDS:
	case MINUTES:
	case DAYS:
	case WEEKS:
	case MONTHS:
		return

	default:
		panic("Unit converter not supported")
	}

	req.StopCounter = h.IntervalCounterConfig.MaxCounterRequest
	req.StopTriggerRequest = helpers.ConvertToSeconds(h.IntervalCounterConfig.IntervalNextRequest.Delay, h.IntervalCounterConfig.IntervalNextRequest.Type)
	helpers.StartScreenTime(req.StopTriggerRequest)

	for range req.Ticker.C {
		req.Counter++

		if req.StopCounter == 0 {
			panic("Max counter request not to be zero")
		}

		switch req.IntervalUnitType {

		case SECONDS:
			req.IntervalTarget = (helpers.ConvertToSeconds(req.IntervalTarget, req.IntervalUnitType))

			if req.Counter == req.NowTriggerRequest {
				req.TotalCounter++

				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				log.Printf("=============================== TOTAL COUNTER REQUEST: %d ==============================\n", req.TotalCounter)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		case MINUTES:
			req.IntervalTarget = helpers.ConvertToSeconds(req.IntervalTarget, req.IntervalUnitType)

			if req.Counter == req.NowTriggerRequest {
				req.TotalCounter++

				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				log.Printf("=============================== TOTAL COUNTER REQUEST: %d ==============================\n", req.TotalCounter)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		case HOURS:
			req.IntervalTarget = helpers.ConvertToSeconds(req.IntervalTarget, req.IntervalUnitType)

			if req.Counter == req.NowTriggerRequest {
				req.TotalCounter++

				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				log.Printf("=============================== TOTAL COUNTER REQUEST: %d ==============================\n", req.TotalCounter)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")
				go HandleFunc()
			}

		default:
			panic("Unit converter not supported")
		}

		if req.Counter == req.IntervalTarget {
			req.NextTriggerRequest = 0
			req.NowTriggerRequest += req.IntervalTarget
			req.NextTriggerRequest += (req.NowTriggerRequest + req.IntervalTarget)

		} else if req.Counter == req.NowTriggerRequest {
			req.NowTriggerRequest += req.IntervalTarget
			req.NextTriggerRequest += req.IntervalTarget

		}

		if req.StopTriggerRequest != 0 && int64(req.TotalCounter) >= int64(req.StopCounter) {
			req.Counter = 0
			req.TotalCounter = 0
			req.NowTriggerRequest = req.IntervalTarget
			req.NextTriggerRequest = (req.NowTriggerRequest + req.IntervalTarget)

			time.AfterFunc(time.Duration(time.Second*time.Duration(req.StopTriggerRequest)), func() {
				req.Release <- true
			})

			helpers.ClearScreen()
			helpers.EndScreenTime(req.StopTriggerRequest)

			<-req.Release
		} else if req.StopTriggerRequest == 0 && int64(req.TotalCounter) >= int64(req.StopCounter) {
			req.Counter = 0
			req.TotalCounter = 0
			req.NowTriggerRequest = req.IntervalTarget
			req.NextTriggerRequest = (req.NowTriggerRequest + req.IntervalTarget)

			time.AfterFunc(time.Duration(time.Second*1), func() {
				req.Release <- true
			})

			helpers.ClearScreen()
			helpers.EndScreenTime(req.StopTriggerRequest)

			<-req.Release
		}

		continue
	}
}

/**
================================================================
==  WATCHDOG BY INTERVAL
================================================================
**/

func (h *amatiaku) ByInterval(HandleFunc func()) {
	var req amatiakuRequest = amatiakuRequest{}

	req.Ticker = time.NewTicker(time.Second * 1)
	req.Release = make(chan bool, 1)

	req.Interval = h.IntervalConfig.IntervalEveryRequest
	req.IntervalTarget = h.IntervalConfig.IntervalTargetRequest
	req.IntervalUnitType = h.IntervalConfig.IntervalUnitType

	req.NowTriggerRequest = req.Interval
	req.NextTriggerRequest = (req.NowTriggerRequest + req.Interval)

	switch h.IntervalConfig.IntervalNextRequest.Type {
	case SECONDS:
	case MINUTES:
	case DAYS:
	case WEEKS:
	case MONTHS:
		return

	default:
		panic("Unit converter not supported")
	}

	req.StopTriggerRequest = helpers.ConvertToSeconds(h.IntervalConfig.IntervalNextRequest.Delay, h.IntervalConfig.IntervalNextRequest.Type)
	helpers.StartScreenTime(req.StopTriggerRequest)

	for range req.Ticker.C {

		req.Counter++

		switch req.IntervalUnitType {

		case SECONDS:
			req.Interval = (helpers.ConvertToSeconds(req.Interval, req.IntervalUnitType))

			if req.Counter == req.NowTriggerRequest {
				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		case MINUTES:
			req.Interval = helpers.ConvertToSeconds(req.Interval, req.IntervalUnitType)

			if req.Counter == req.NowTriggerRequest {
				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		case HOURS:
			req.Interval = helpers.ConvertToSeconds(req.Interval, req.IntervalUnitType)

			if req.Counter == req.NowTriggerRequest {
				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		case DAYS:
			req.Interval = helpers.ConvertToSeconds(req.Interval, req.IntervalUnitType)

			if req.Counter == req.NowTriggerRequest {
				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		case WEEKS:
			req.Interval = helpers.ConvertToSeconds(req.Interval, req.IntervalUnitType)

			if req.Counter == req.NowTriggerRequest {
				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		case MONTHS:
			req.Interval = helpers.ConvertToSeconds(req.Interval, req.IntervalUnitType)

			if req.Counter == req.NowTriggerRequest {
				fmt.Printf("\n")
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				log.Printf("=============================== EXECUTION PROCESS SUCCESS =============================\n")
				fmt.Printf("\n")
				log.Printf("=============================== NOW TIME REQUEST: %dS =================================\n", req.NowTriggerRequest)
				log.Printf("=============================== NEXT TIME REQUEST: %dS ================================\n", req.NextTriggerRequest)
				fmt.Printf("\n")
				log.Printf("=======================================================================================\n")
				fmt.Printf("\n")
				fmt.Printf("\n")

				go HandleFunc()
			}

		default:
			panic("Unit converter not supported")
		}

		if req.Counter == req.Interval {
			req.NextTriggerRequest = 0
			req.NowTriggerRequest += req.Interval
			req.NextTriggerRequest += (req.NowTriggerRequest + req.Interval)

		} else if req.Counter == req.NowTriggerRequest {
			req.NowTriggerRequest += req.Interval
			req.NextTriggerRequest += req.Interval

		}

		if req.IntervalTarget != 0 && req.StopTriggerRequest != 0 && int64(req.Counter) >= int64(req.IntervalTarget) {
			req.Counter = 0
			req.NowTriggerRequest = req.Interval
			req.NextTriggerRequest = (req.NowTriggerRequest + req.Interval)

			time.AfterFunc(time.Duration(time.Second*time.Duration(req.StopTriggerRequest)), func() {
				req.Release <- true
			})

			helpers.ClearScreen()
			helpers.EndScreenTime(req.StopTriggerRequest)

			<-req.Release
		}

		continue
	}
}
