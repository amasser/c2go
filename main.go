// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// C2go converts (some) C programs to Go.
// It was developed by Russ Cox to convert the Go runtime and compiler to Go.
// It is progressing toward becoming a general-purpose translation tool,
// but it does not (and never will) translate all possible C programs.
// You will need to refactor your C program into the subset of C that c2go
// can handle, and give it translation hints in the configuration file.
//
// Usage of c2go:
//   c2go [options] *.c
//
// Options:
//   -c string
//     	config file
//   -dst string
//     	GOPATH root of destination (default "/tmp/c2go")
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andybalholm/c2go/cc"
)

var (
	cfgFile = flag.String("c", "", "config file")
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: c2go [options] *.c\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	cfg := new(Config)
	if *cfgFile != "" {
		cfg.read(*cfgFile)
	}
	for _, t := range cfg.typedefs {
		cc.AddTypeName(t)
	}

	files := args
	prog, err := cc.ReadMany(files)
	if err != nil {
		log.Fatal(err)
	}

	rewriteTypes(cfg, prog)
	rewriteSyntax(cfg, prog)
	rewriteLen(cfg, prog)
	fixGoTypes(cfg, prog)
	simplifyBool(cfg, prog)
	renameDecls(cfg, prog)
	exportDecls(cfg, prog)
	writeGoFiles(cfg, prog)

	for _, d := range cfg.diffs {
		if d.used == 0 {
			fmt.Fprintf(os.Stderr, "%s: unused diff\n", d.line)
		}
	}
}
