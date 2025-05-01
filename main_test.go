package main

import (
	"bytes"
	"testing"
)

func TestPrintArgsHappyPath(t *testing.T) {
	outputBuffer := bytes.Buffer{}
	errorBuffer := bytes.Buffer{}

	cfg, err := NewCliConfig(CliConfigWithOutputWriter(&outputBuffer), CliConfigWithErrorWriter(&errorBuffer))
	if err != nil {
		t.Fatalf("NewCliConfig error: %v", err)
	}

	inputArgs := []string{"one", "two", "three", "four", "five"}

	err = cfg.PrintArgs(inputArgs)
	if err != nil {
		t.Fatalf("PrintArgs error: %v", err)
	}

	expectedErrorOutput := "one\ntwo\nthree\n"
	expectedStandardOutput := "four\nfive\n"

	if outputBuffer.String() != expectedStandardOutput {
		t.Fatalf("PrintArgs output:\n%s\nexpected:\n%s", outputBuffer.String(), expectedStandardOutput)
	}

	if errorBuffer.String() != expectedErrorOutput {
		t.Fatalf("PrintArgs error:\n%s\nexpected:\n%s", errorBuffer.String(), expectedErrorOutput)
	}
}
