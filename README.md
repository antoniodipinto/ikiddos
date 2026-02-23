# IkiDDoS

A Go library for HTTP load testing. **For educational purposes only.**

> Inspired by [Konstantin8105/DDoS](https://github.com/Konstantin8105/DDoS)

## Installation

### As a library

```bash
go get github.com/antoniodipinto/ikiddos
```

### As a CLI

```bash
go install github.com/antoniodipinto/ikiddos/cmd/ikiddos@latest
```

## CLI Usage

```bash
ikiddos -url http://example.com -duration 1m -clients 20
```

```bash
ikiddos -url http://example.com/api -method POST -content-type application/json -body '{"key":"value"}'
```

With custom headers:

```bash
ikiddos -url http://example.com/subscribe \
  -method POST \
  -content-type application/x-www-form-urlencoded \
  -body 'email=test@example.com&mc_language=it&subscribed=1' \
  -header 'User-Agent: Mozilla/5.0' \
  -header 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
  -header 'Accept-Language: it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7'
```

| Flag            | Default | Description                                        |
|-----------------|---------|----------------------------------------------------|
| `-url`          | —       | Target URL (required)                              |
| `-duration`     | `30s`   | Attack duration (e.g. `30s`, `1m`)                 |
| `-timeout`      | `5s`    | Per-request timeout                                |
| `-clients`      | `10`    | Number of concurrent workers                       |
| `-method`       | `GET`   | HTTP method: `GET` or `POST`                       |
| `-content-type` | —       | Content-Type header (required for POST)            |
| `-body`         | —       | Request body (POST only)                           |
| `-header`       | —       | Custom header as `Key: Value` (repeatable)         |

The CLI prints a live counter of requests, successes, and errors, and outputs a final report when finished. Press `Ctrl+C` to stop early.

## Library Usage

### GET request

```go
package main

import (
	"fmt"
	"time"

	"github.com/antoniodipinto/ikiddos"
)

func main() {
	attack := ikiddos.New()

	err := attack.SetConfig(&ikiddos.Config{
		Url:      "http://example.com",
		Duration: 30 * time.Second,
		Timeout:  5 * time.Second,
		Clients:  10,
		Method:   ikiddos.HttpGet,
	})
	if err != nil {
		panic(err)
	}

	attack.Start()

	time.Sleep(30 * time.Second)

	report := attack.Report()
	fmt.Printf("Requests: %d | Success: %d | Errors: %d\n",
		report.Requests, report.Success, report.Error)
}
```

### POST request

```go
attack := ikiddos.New()

err := attack.SetConfig(&ikiddos.Config{
	Url:         "http://example.com/api",
	Duration:    20 * time.Second,
	Timeout:     5 * time.Second,
	Clients:     5,
	Method:      ikiddos.HttpPost,
	ContentType: "application/json",
	Body:        []byte(`{"key": "value"}`),
})
```

### Custom headers

```go
err := attack.SetConfig(&ikiddos.Config{
	Url:         "http://example.com/subscribe",
	Duration:    30 * time.Second,
	Timeout:     5 * time.Second,
	Clients:     10,
	Method:      ikiddos.HttpPost,
	ContentType: "application/x-www-form-urlencoded",
	Body:        []byte("email=test@example.com&mc_language=it&subscribed=1"),
	Headers: map[string]string{
		"User-Agent":      "Mozilla/5.0",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Language": "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7",
	},
})
```

## Configuration

| Field            | Type            | Description                                      |
|------------------|-----------------|--------------------------------------------------|
| `Url`            | `string`        | Target URL (must include scheme and host)         |
| `Duration`       | `time.Duration` | How long the attack runs                          |
| `Timeout`        | `time.Duration` | Per-request timeout                               |
| `Clients`        | `int`           | Number of concurrent workers (must be > 0)        |
| `Method`         | `string`        | `ikiddos.HttpGet` or `ikiddos.HttpPost`           |
| `ContentType`    | `string`        | Required for POST requests                        |
| `Body`           | `[]byte`        | Request body (POST only)                          |
| `Headers`        | `map[string]string` | Custom HTTP headers                           |
| `Proxy`          | `*ProxyConfig`  | Optional proxy (`Address` + `Port`)               |
| `ConsoleEnabled` | `bool`          | Enable console output (default `true`)            |

## API

| Method                        | Description                          |
|-------------------------------|--------------------------------------|
| `New() *Attack`               | Create a new Attack instance         |
| `SetConfig(*Config) error`    | Validate and apply configuration     |
| `Start()`                     | Begin the attack in background       |
| `Pause()`                     | Pause the attack loop                |
| `Stop()`                      | Stop the attack permanently          |
| `IsEnabled() bool`            | Check if the attack is active        |
| `IsStopped() bool`            | Check if the attack has been stopped |
| `Report() *Report`            | Get current stats                    |

## Disclaimer

This project is strictly for **educational and authorized testing purposes**. Unauthorized use of this tool against systems you do not own or have explicit permission to test is illegal and unethical. The author assumes no liability for misuse.

