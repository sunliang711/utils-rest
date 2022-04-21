package utils

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	command := "solc --version"
	output, err := Run(command, nil)
	if err != nil {
		t.Fatalf("run error: %v", err)
	}
	t.Logf("run result: %s", string(output))
	splits := strings.Split(string(output), "Version: ")
	t.Logf("'%v'", strings.TrimSpace(splits[1]))
}
