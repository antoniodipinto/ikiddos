# IkiDDoS

A Go library for HTTP load testing. **For educational purposes only.**

> Inspired by [Konstantin8105/DDoS](https://github.com/Konstantin8105/DDoS)

## Installation

```bash
go get github.com/antoniodipinto/ikiddos
```

## Usage

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

