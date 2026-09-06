package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"testing/synctest"
)

// -n 1 drives the real controller, which sleeps for six seconds of wall time and
// writes to package-level state in esp32-controller. So these run inside a
// synctest bubble, and not in parallel.
//
// The controller's own output never reaches the stdout buffer below. It goes
// straight to the process, because serial.Print stands in for an Arduino call
// that has no writer to hand it. These assert on what run itself prints.

func TestRun(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantErr     error
		wantVersion bool
		wantUsage   bool
	}{
		{name: "one loop runs cleanly", args: []string{"-n", "1"}},
		{name: "-version prints one line of data", args: []string{"-version"}, wantVersion: true},
		{name: "-h prints the usage and succeeds", args: []string{"-h"}, wantUsage: true},
		{name: "a negative count is a usage error", args: []string{"-n", "-1"}, wantErr: errUsage},
		{name: "an unknown flag is a usage error", args: []string{"-bogus"}, wantErr: errUsage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				var stdout, stderr bytes.Buffer

				err := run(t.Context(), tt.args, &stdout, &stderr)

				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("run: got %v, want %v (stderr: %s)", err, tt.wantErr, stderr.String())
				}
				if tt.wantVersion {
					if got := stdout.String(); strings.Count(got, "\n") != 1 || strings.TrimSpace(got) == "" {
						t.Errorf("stdout: got %q, want one non-empty line", got)
					}
					if stderr.Len() != 0 {
						t.Errorf("stderr: got %q, want nothing", stderr.String())
					}
				}
				if tt.wantUsage && !strings.Contains(stderr.String(), "Usage:") {
					t.Errorf("stderr: got %q, want the usage text", stderr.String())
				}
				if tt.wantErr == nil && !tt.wantVersion && !tt.wantUsage && stderr.Len() != 0 {
					t.Errorf("stderr: got %q, want nothing on a clean run", stderr.String())
				}
			})
		})
	}
}

func TestVersionIsNeverEmpty(t *testing.T) {
	if version() == "" {
		t.Error("version: got an empty string, want a version or \"unknown\"")
	}
}
