package ikiddos_test

import (
	"testing"
	"time"

	"github.com/antoniodipinto/ikiddos"
)

func TestAttack_SetConfig(t *testing.T) {
	attack := ikiddos.New()

	err := attack.SetConfig(&ikiddos.Config{
		Url:            "http://localhost",
		Duration:       time.Second * 20,
		Proxy:          nil,
		ConsoleEnabled: false,
		Clients:        1,
		Timeout:        0,
		Method:         ikiddos.HttpGet,
		ContentType:    "",
	})

	if err != nil {
		t.Error("Wrong configuration", err)
	}

}

func TestAttack_Pause(t *testing.T) {
	attack := ikiddos.New()

	err := attack.SetConfig(&ikiddos.Config{
		Url:            "http://localhost",
		Duration:       time.Second * 20,
		Proxy:          nil,
		ConsoleEnabled: false,
		Clients:        1,
		Timeout:        0,
		Method:         ikiddos.HttpGet,
		ContentType:    "",
	})

	if err != nil {
		t.Error("Wrong configuration", err)
	}

	attack.Start()

	attack.Pause()

	if attack.IsEnabled() {
		t.Error("Attack should be paused")
	}
}

func TestAttack_Stop(t *testing.T) {
	attack := ikiddos.New()

	err := attack.SetConfig(&ikiddos.Config{
		Url:            "http://localhost",
		Duration:       time.Second * 20,
		Proxy:          nil,
		ConsoleEnabled: false,
		Clients:        1,
		Timeout:        0,
		Method:         ikiddos.HttpGet,
		ContentType:    "",
	})

	if err != nil {
		t.Error("Wrong configuration", err)
	}

	attack.Start()

	attack.Stop()

	if !attack.IsStopped() {
		t.Error("Attack should be stopped")
	}
}

func TestAttack_PostRequest(t *testing.T) {
	attack := ikiddos.New()

	err := attack.SetConfig(&ikiddos.Config{
		Url:            "http://localhost",
		Duration:       time.Second * 20,
		Proxy:          nil,
		ConsoleEnabled: false,
		Clients:        1,
		Timeout:        0,
		Method:         ikiddos.HttpPost,
		ContentType:    "application/json",
		Body:           []byte(`{"key": "value"}`),
	})

	if err != nil {
		t.Error("Wrong configuration", err)
	}

	attack.Start()

	if !attack.IsEnabled() {
		t.Error("Attack should be started")
	}
}

func TestAttack_Start(t *testing.T) {
	attack := ikiddos.New()

	err := attack.SetConfig(&ikiddos.Config{
		Url:            "http://localhost",
		Duration:       time.Second * 20,
		Proxy:          nil,
		ConsoleEnabled: false,
		Clients:        1,
		Timeout:        0,
		Method:         ikiddos.HttpGet,
		ContentType:    "",
	})

	if err != nil {
		t.Error("Wrong configuration", err)
	}

	attack.Start()

	if !attack.IsEnabled() {
		t.Error("Attack should be started")
	}
}
