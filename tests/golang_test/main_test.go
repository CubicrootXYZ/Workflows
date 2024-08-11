package golangtest_test

import (
	"golangtest"
	"os"
	"testing"
)

func TestSomething(t *testing.T) {
	golangtest.DoSomething("a test")
}

func TestEnvVarsAreSet(t *testing.T) {
	v := os.Getenv("TESTENV")
	t.Logf("env var is: %s", v)
	if v != "value" {
		t.Fail()
	}
}
