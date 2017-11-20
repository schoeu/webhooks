package utils

import (
	"testing"
	"strings"
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

func TestAnalysis(t *testing.T) {
	command := "router"
	param := "sh ./builder"
	cmdStr := command + ":" + param
	r, c := Analysis(cmdStr)
	if r != command {
		t.Errorf("got [%s] expected [%s]", r, command)
	}

	if c != param {
		t.Errorf("got [%s] expected [%s]", c, param)
	}
}

func TestExecCmds(t *testing.T) {
	str := "123"
	c := "echo " + str
	rs := ExecCmds(c)
	if strings.TrimSpace(rs) != str {
		t.Errorf("got [%s] expected [%s]",str, rs)
	}
}
