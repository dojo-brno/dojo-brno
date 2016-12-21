package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		comment string
		// input
		stdin  string
		stdout io.Writer
		// want
		exitCode         int
		wantStdout       string
		wantStderrPrefix string
	}{
		{
			comment:    "hello in hexadecimal plus line feed",
			stdin:      "hello",
			exitCode:   0,
			wantStdout: "68656c6c6f\n",
		},
		{
			comment:    "empty input should produce empty output",
			stdin:      "",
			exitCode:   0,
			wantStdout: "",
		},
		{
			comment:          "failure to write to stdout should print error to stderr",
			stdin:            "hello",
			stdout:           (*os.File)(nil),
			exitCode:         1,
			wantStderrPrefix: "error: ",
		},
	}
	for _, tt := range tests {
		var stdout, stderr bytes.Buffer
		cmd := helperCommand()
		cmd.Stdin = strings.NewReader(tt.stdin)
		cmd.Stdout = &stdout
		if tt.stdout != nil {
			cmd.Stdout = tt.stdout
		}
		cmd.Stderr = &stderr
		err := cmd.Run()
		if (err == nil) != (tt.exitCode == 0) {
			t.Log(tt.comment)
			t.Errorf("got %v, want exit status %d", err, tt.exitCode)
		}
		if got, want := stdout.String(), tt.wantStdout; got != want {
			t.Log(tt.comment)
			t.Errorf("stdout: got %q, want %q", got, want)
		}
		if got, want := stderr.String(), tt.wantStderrPrefix; !strings.HasPrefix(got, want) {
			t.Log(tt.comment)
			t.Errorf("stderr: got %q, want %q", got, want)
		}
	}
}

func helperCommand() *exec.Cmd {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcess")
	cmd.Env = []string{"WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("WANT_HELPER_PROCESS") == "1" {
		main()
		os.Exit(0)
	}
}
