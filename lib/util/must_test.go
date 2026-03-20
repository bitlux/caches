package util

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestMustSuccess(t *testing.T) {
	Must(nil)
}

func TestMustFail(t *testing.T) {
	if os.Getenv("MUST_FAIL") == "1" {
		Must(errors.New("die"))
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMustFail")
	cmd.Env = append(os.Environ(), "MUST_FAIL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
