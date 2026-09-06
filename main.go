// Command esp32 runs your controller on this machine, so you can watch it work
// before you flash it to a device.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/andygeiss/esp32/internal/device"
	"github.com/andygeiss/esp32/internal/simulator"
)

var errUsage = errors.New("usage error")

const usage = `esp32 runs your ESP32 controller on this machine.

Usage:
  esp32 [flags]

The controller writes to stdout, the way a serial monitor shows it. Everything
else goes to stderr. Press Ctrl-C to stop; a second Ctrl-C kills the process.

Flags:
`

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// One loop can hold the process for seconds, because timer.Delay sleeps and
	// takes no context. Hand the signal back after the first Ctrl-C so a second
	// one still kills the process.
	go func() {
		<-ctx.Done()
		stop()
	}()
	err := run(ctx, os.Args[1:], os.Stdout, os.Stderr)
	switch {
	case err == nil:
	case errors.Is(err, errUsage):
		os.Exit(2) // message already printed where the error was detected
	default:
		fmt.Fprintf(os.Stderr, "esp32: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("esp32", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		fmt.Fprint(stderr, usage)
		fs.PrintDefaults()
	}
	loops := fs.Int("n", 0, "stop after this many loops; 0 runs until you interrupt it")
	showVersion := fs.Bool("version", false, "print the version and exit")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil // -h: usage already printed, exit 0
		}
		return errUsage // bad flag: message already printed by fs, exit 2
	}
	if *showVersion {
		fmt.Fprintln(stdout, version())
		return nil
	}
	if *loops < 0 {
		fmt.Fprintf(stderr, "esp32: -n %d: want zero or more\n", *loops)
		return errUsage
	}
	return simulator.Run(ctx, device.NewController(), *loops)
}

// version reports what the toolchain stamped into the binary. No -ldflags step:
// since Go 1.24 a build from a checkout carries VCS metadata too.
func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	if v := info.Main.Version; v != "" && v != "(devel)" {
		return v
	}
	return "unknown" // no VCS metadata, so there is nothing to fall back to
}
