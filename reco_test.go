package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	var filename string = "test-text.txt"

	testCases := []struct {
		fileName    string
		args        string
		expectedOut string
	}{
		{filename, "-w", "word count: 400\n"},
		{filename, "-v", "vowel count: 800\n"},
		{filename, "-c", "consonant count: 1000\n"},
		{filename, "-l", "letter count: 1800\n"},
		{filename, "-s", "space count: 300\n"},
		{filename, "-d", "digit count: 300\n"},
		{filename, "-p", "pmark count: 200\n"},
	}

	// args := []string{"-w", "-v", "-c", "-l", "-s", "-d", "-p"}

	for _, test := range testCases {
		testname := fmt.Sprintf("ReCo count %s %s", test.args, test.fileName)

		t.Run(testname, func(t *testing.T) {
			cmdTest := exec.Command("ReCo", "count", test.args, test.fileName)
			output, _ := cmdTest.Output()
			if string(output) != test.expectedOut {
				t.Errorf("Expected %s, got %s", test.expectedOut, string(output))
			}
		})
	}

}
