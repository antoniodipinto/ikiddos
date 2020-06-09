package ikiddos_test

import (
	"github.com/antoniodipinto/ikiddos"
	"testing"
	"time"
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
}
