// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"fmt"
	"os/exec"
	"strings"
)

func ReadMany(names []string) (*Prog, error) {
	lx := &lexer{}
	var prog *Prog
	for _, name := range names {
		data, err := preprocess(name)
		if err != nil {
			return nil, err
		}
		lx.start = startProg
		lx.lexInput = lexInput{
			input:  string(data),
			file:   name,
			lineno: 1,
		}
		lx.scope = &Scope{}
		lx.parse()
		if lx.errors != nil {
			return nil, fmt.Errorf("%v", lx.errors[0])
		}
		if prog == nil {
			prog = lx.prog
		} else {
			prog.Span.End = lx.prog.Span.End
			prog.Decls = mergeDecls(prog.Decls, lx.prog.Decls)
		}
		lx.prog = nil
		for sc := lx.scope; sc != nil; sc = sc.Next {
			for name, decl := range sc.Decl {
				if decl.Storage&Static != 0 || (decl.Storage&Typedef != 0 && strings.HasSuffix(decl.Span.Start.File, ".c")) {
					delete(sc.Decl, name)
				}
			}
			for name, typ := range sc.Tag {
				if strings.HasSuffix(typ.Span.Start.File, ".c") {
					delete(sc.Tag, name)
				}
			}
		}
	}
	lx.prog = prog
	lx.assignComments()
	lx.typecheck(lx.prog)
	if lx.errors != nil {
		return nil, fmt.Errorf("%v", strings.Join(lx.errors, "\n"))
	}

	removeDuplicates(lx.prog)

	return lx.prog, nil
}

func mergeDecls(dest, toAdd []*Decl) []*Decl {
	have := make(map[string]bool)
	for _, d := range dest {
		have[fmt.Sprintf("%v:%x", d.Span, d.Name)] = true
	}
	for _, d := range toAdd {
		if !have[fmt.Sprintf("%v:%x", d.Span, d.Name)] {
			dest = append(dest, d)
		}
	}
	return dest
}

// preprocess runs the GCC preprocessor on the specified file, and returns the
// output.
func preprocess(filename string) ([]byte, error) {
	gcc := exec.Command("gcc", "-C", "-E", "-DC2GO", filename)
	data, err := gcc.Output()
	if exitErr, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("%s", exitErr.Stderr)
	}
	data = append(data, '\n')
	return data, err
}

type Prog struct {
	SyntaxInfo
	Decls []*Decl
}

// removeDuplicates drops the duplicated declarations
// caused by forward decls from prog.
// It keeps the _last_ of each given declaration,
// assuming that's the complete one.
// This heuristic tends to preserve something like
// source order.
// It would be defeated by someone writing a "forward"
// declaration following the real definition.
func removeDuplicates(prog *Prog) {
	count := map[*Decl]int{}
	for _, d := range prog.Decls {
		count[d]++
	}
	var out []*Decl
	for _, d := range prog.Decls {
		count[d]--
		if count[d] == 0 {
			out = append(out, d)
		}
	}
	prog.Decls = out
}
