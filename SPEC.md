# SPEC

The project brief. Every task brief is a delta against this one.

**Job.** Write the code that runs on your ESP32 in Go, watch it run on your own
machine, then transpile it and flash it to the device.

**Why.** The Arduino IDE cannot tell you whether a sketch works without flashing
it first, so every mistake costs a compile, a flash and a stare at the serial
monitor. Here the same controller is an ordinary Go program: `go test` catches
the mistake in a second, and the hardware only sees code that already works.

**Guardrails.**

- `internal/device/controller.go` is the only file the transpiler reads, and it
  understands a small subset of Go. Keep that file to plain functions,
  constants, package-level variables and calls into `esp32-controller`.
- Waivers and the decisions this project has named live in `README.md` under
  *Baseline deviations*.
- ESP-IDF is a README install, not a Make target. `make` never clones an SDK or
  downloads a compiler.
- `make check` must not need ESP-IDF. It stays green on a Mac carrying nothing
  but Go, the way esp32-watch's does.

**Done means.** The CLI checklist in the
[engineering baseline](https://github.com/andygeiss/baseline) is walked, and
`make ci` is green.
