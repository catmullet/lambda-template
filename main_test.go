package main

import (
	"github.com/kyani-inc/kms-lambda-template/local/helpers/environments"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	environments.LoadEnvironmentVariablesFromYml(environments.Staging)
	code := m.Run()

	os.Exit(code)
}
