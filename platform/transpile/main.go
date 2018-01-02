package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andygeiss/esp32/application/transpile"
	"github.com/andygeiss/esp32/infrastructure/ino"
)

var (
	// APPNAME ...
	APPNAME string
	// BUILD ...
	BUILD string
	// VERSION ...
	VERSION string
)

func main() {
	mapping := flag.String("mapping", os.Getenv("HOME")+"/.goat.json", "API Mapping file")
	source := flag.String("source", "", "Golang source file")
	target := flag.String("target", "", "Arduino sketch file")
	flag.Parse()
	if *mapping == "" || *source == "" || *target == "" {
		printUsage()
		return
	}
	// Read the Golang source file.
	in, err := os.Open(*source)
	if err != nil {
		log.Fatalf("Go source file [%s] could not be opened! %v", *source, err)
	}
	defer in.Close()
	// Create the Arduino sketch file.
	os.Remove(*target)
	out, err := os.OpenFile(*target, os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		log.Fatalf("Arduino sketch file [%s] could not be opened! %v", *target, err)
	}
	// Transpiles the Golang source into Arduino sketch.
	m := ino.NewMapping(*mapping)
	if err := m.Read(); err != nil {
		log.Fatal(err)
	}
	worker := ino.NewWorker(in, out, m)
	trans := transpile.NewTranspiler(worker)
	if err := trans.Transpile(); err != nil {
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Printf("%s %s (build %s)\n\n", APPNAME, VERSION, BUILD)
	fmt.Print("GOAT transpiles Golang source into corresponding Arduino sketches.\n\n")
	fmt.Print("Options:\n")
	flag.PrintDefaults()
	fmt.Print("\n")
	fmt.Print("Example:\n")
	fmt.Printf("\t%s -source application/blink/controller.go -target application/blink/controller.ino\n\n", APPNAME)
}
