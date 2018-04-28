package config_test

import (
	"github.com/m0cchi/postal_worker/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	libPath := "/etc/sysconfig/postal_worker.d/lib"
	config, err := config.NewConfig("testdata/postal_worker.toml")
	if err != nil {
		t.Fatal(err)
	}

	err = config.Validate()
	if err == nil {
		t.Fatalf("expect: stat %s ~~~~~", libPath)
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

func TestEmpty(t *testing.T) {
	_, err := config.NewConfig("")
	if err == nil {
		t.Fatal("expect: no such file or directory")
	}
}
