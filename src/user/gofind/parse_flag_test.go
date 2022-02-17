package main

import (
	"flag"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

type parseResult struct {
	conf *config
	err  error
}

func Test_ParseFlags(t *testing.T) {
	tests := []struct {
		args []string
		want parseResult
	}{
		{[]string{"-start", "/usr/local"},
			parseResult{nil, MissingPatternErr},
		},
		{[]string{"-pattern", "foo.baz"},
			parseResult{nil, MissingStartErr},
		},
		{[]string{"-start", "/", "-pattern", "foo.go"},
			parseResult{&config{start: "/", pattern: regexp.MustCompile("foo.go")}, nil},
		},
		{
			[]string{"-h"},
			parseResult{nil, flag.ErrHelp},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, err := parseFlags("program", tt.args)
			got := parseResult{conf, err}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conf got %+v, want %+v", got, tt.want)
			}
		})
	}
}
