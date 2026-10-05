package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Income(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "34320") {
		t.Errorf("Expected output to contain '34320', but got '%s'", output)
	}
}
