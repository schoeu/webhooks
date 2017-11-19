package utils

import (
	"testing"
)

func TestGetCwd(t *testing.T) {
	cwd := GetCwd()
	if cwd == "" {
		t.Errorf("got [%s] is not expected", cwd)
	}
}

func TestValidCmd(t *testing.T) {
	command := "echo"
	param := "1"
	testStr := command + " " + param
	r, c, ok := ValidCmd(testStr)

	if ok != true {
		t.Errorf("got [%s] expected [%s]", ok, true)
	}

	if r != command {
		t.Errorf("got [%s] expected [%s]", r, command)
	}

	if c != param {
		t.Errorf("got [%s] expected [%s]", c, param)
	}
}
