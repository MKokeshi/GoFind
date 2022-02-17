// Package gofind implements una ricerca del file concurrente
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type config struct {
	start   string
	pattern *regexp.Regexp
}

var (
	MissingPatternErr = errors.New("è richiesto il pattern")
	MissingStartErr   = errors.New("è richiesto lo start")
)

func parseFlags(programName string, args []string) (*config, error) {
	flags := flag.NewFlagSet(programName, flag.ContinueOnError)

	var rawStart string
	var rawPattern string
	flags.StringVar(&rawStart, "start", "", "Lo starting path per la ricerca")
	flags.StringVar(&rawPattern, "pattern", "", "Un pattern con cui far matchare i file")
	err := flags.Parse(args)
	if err != nil {
		return nil, err
	}

	if rawPattern == "" {
		return nil, MissingPatternErr
	}

	if rawStart == "" {
		return nil, MissingStartErr
	}

	absStart, err := filepath.Abs(rawStart)
	if err != nil {
		return nil, err
	}
	start := filepath.Clean(absStart)

	pattern, err := regexp.Compile(rawPattern)
	if err != nil {
		return nil, err
	}

	return &config{start, pattern}, nil
}

func main() {
	conf, err := parseFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		flag.PrintDefaults()
		os.Exit(2)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Errore: %s\n", err)
		os.Exit(1)
	}

	finder := NewFinder(conf.pattern)
	fmt.Printf("start: %s, pattern: %s\n", conf.start, conf.pattern)
	matches, err := finder.Find(conf.start)
	for _, match := range matches {
		fmt.Println(match)
	}
}
