# Amatiaku (On Progress)

**Amatiku** is a simple watchdog implementation request with zero dependency, Watchdog is useful for use cases like if you have an api for exchange rate in your app and you want to update this exchange rate from a third party, let's say from [xe.com](https://www.xe.com/currencyconverter/), to get the latest exchange rate, you can use watchdog to handle this mechanism, call the api from xe.com every 5 seconds, and update the exchange rate in your database. It is also useful for many other use cases.

## Example

```go
func main() {
	amatiaku := NewAmatiAku(&AmatiakuConfig{
		IntervalConfig: IntervalConfig{
			IntervalEveryRequest:  5, // Execution process every 5 seconds.
			IntervalTargetRequest: 60, // Stop the process if the time has reached 1 minute, counting from the time the function is executed.
			IntervalNextRequest: NextRequest{ // The next execution process starts again after 1 minute.
				Delay: 1,
				Type:  MINUTES,
			},
			IntervalUnitType: SECONDS,
		},
	})

	amatiaku.ByInterval(func() {
			// Do something here ....
	})
}
```


```go
func main() {
	amatiaku := NewAmatiAku(&AmatiakuConfig{
		IntervalCounterConfig: IntervalCounterConfig{
			IntervalEveryRequest: 5, // Execution process every 5 seconds.
			IntervalUnitType:     SECONDS,
			IntervalNextRequest: NextRequest{ // The next execution process starts again after 1 minute.
				Delay: 1,
				Type:  MINUTES,
			},
			MaxCounterRequest: 30, // Stop the process if the counter equal 30x request.
		},
	})

	amatiaku.ByIntervalCounter(func() {
			// Do something here ....
	})
}
```