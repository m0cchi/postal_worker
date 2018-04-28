package config_test

import (
	"fmt"
	"github.com/m0cchi/postal_worker/config"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	libPath := "/etc/sysconfig/postal_worker.d/lib"
	config, err := config.NewConfig("testdata/postal_worker.toml")
	if err != nil {
		message := fmt.Sprintf("%v", err)
		if !strings.Contains(message, libPath) {
			fmt.Println(message)
			t.Fatal(err)
		}
	}

	if config.Module.Lib != libPath {
		t.Fatalf("config.Module.Lib is %v", config.Module.Lib)
	}

	if config.Server.Host != "0.0.0.0" {
		t.Fatalf("config.Server.Address is %v", config.Server.Host)
	}

	if config.Server.Port != 3000 {
		t.Fatalf("config.Server.Port is %v", config.Server.Port)
	}
}
