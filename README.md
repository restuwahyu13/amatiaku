# Amatiaku (On Progress)

**Amatiku** is simple watchdog implementation request with zero dependency, Watchdog is useful for use cases like if you have an api for exchange rate in your app and you want to update this exchange rate from a third party, let's say from [xe.com](https://www.xe.com/currencyconverter/), to get the latest exchange rate, you can use watchdog to handle this mechanism, call the api from xe.com every 5 seconds, and update the exchange rate in your database and for other many use cases.

## Example

```go
func main() {
	amatiaku := NewAmatiAku(&AmatiakuConfig{
		IntervalConfig: IntervalConfig{
			IntervalEveryRequest: 5,
			IntervalNextRequest:  300,
			IntervalUnitType:     SECONDS,
		},
	})

	amatiaku.ByInterval(func() {
		// Do something here ....
	})
}
```