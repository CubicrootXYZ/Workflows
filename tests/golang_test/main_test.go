package golangtest_test

import (
	golangtest "golang_test"
	"testing"
)

func TestSomething(t *testing.T) {
	golangtest.DoSomething("a test")
}
