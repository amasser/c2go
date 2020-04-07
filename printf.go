// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/andybalholm/c2go/cc"
)

func tryPrintf(curfn *cc.Decl, x *cc.Expr, name string, fmtpos int, newName string) bool {
	if (x.Left.Text == name || (strings.Contains(name, ".") || strings.Contains(name, "->")) && x.Left.String() == name) && len(x.List) >= fmtpos+1 && x.List[fmtpos].Op == cc.String {
		x.List = append(x.List[:fmtpos+1], fixPrintFormat(curfn, x.List[fmtpos], x.List[fmtpos+1:])...)
		if newName != "" {
			x.Left.Text = newName
			x.Left.XDecl = nil
		}
		return true
	}
	return false
}

func fixPrintf(curfn *cc.Decl, x *cc.Expr) bool {
	if x.Op != cc.Call {
		return false
	}
	if tryPrintf(curfn, x, "sprintf", 1, "fmt.Sprintf") {
		targ := x.List[0]
		x.List = x.List[1:]
		x.Right = copyExpr(x)
		x.List = nil
		x.Left = targ
		x.Op = cc.Eq

		// sprint(strchr(s, 0), "hello") => s += "hello"
		if targ.Op == cc.Call && len(targ.List) == 2 && targ.Left.Text == "strchr" && targ.List[1].Text == "0" {
			x.Left = targ.List[0]
			x.Op = cc.AddEq
		} else if targ.Op == cc.Add && targ.Right.Op == cc.Call && targ.Right.Left.Text == "strlen" && len(targ.Right.List) == 1 && GoString(targ.Left) == GoString(targ.Right.List[0]) {
			x.Left = targ.Left
			x.Op = cc.AddEq
		}
		return true
	}
	if tryPrintf(curfn, x, "snprintf", 2, "fmt.Sprintf") {
		targ := x.List[0]
		x.List = x.List[2:]
		x.Right = copyExpr(x)
		x.List = nil
		x.Left = targ
		x.Op = cc.Eq
		return true
	}
	if tryPrintf(curfn, x, "printf", 0, "fmt.Printf") {
		return true
	}
	if tryPrintf(curfn, x, "fprintf", 1, "fmt.Fprintf") {
		if x.List[0].String() == "2" {
			x.List[0] = &cc.Expr{Op: cc.Name, Text: "os.Stderr"}
		}
		return true
	}
	if tryPrintf(curfn, x, "warn", 0, "") {
		return true
	}

	return false
}

func fixPrintFormat(curfn *cc.Decl, fx *cc.Expr, args []*cc.Expr) []*cc.Expr {
	for _, arg := range args {
		fixGoTypesExpr(curfn, arg, nil)
		cc.Preorder(arg, func(x cc.Syntax) {
			if x, ok := x.(*cc.Expr); ok && x.Op == cc.Name && strings.HasPrefix(x.Text, "bigP") {
				x.Text = "p"
				x.XDecl = nil
			}
		})
	}

	narg := 0
	for j, text := range fx.Texts {
		format, err := strconv.Unquote(text)
		if err != nil {
			fprintf(fx.Span, "cannot parse quoted string: %v", err)
			return args
		}

		var buf bytes.Buffer
		start := 0
		for i := 0; i < len(format); i++ {
			if format[i] != '%' {
				continue
			}
			buf.WriteString(format[start:i])
			start = i
			i++
			if i < len(format) && format[i] == '%' {
				buf.WriteByte('%')
				buf.WriteByte('%')
				start = i + 1
				continue
			}
			for i < len(format) {
				c := format[i]
				switch c {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '#', '-', '.', ',', ' ', 'h', 'j', 'l', '*':
					i++
					continue
				}
				break
			}
			if i >= len(format) {
				fprintf(fx.Span, "print format ends mid-verb")
				return args
			}
			flags, verb := format[start:i], format[i]
			start = i + 1

			allFlags := flags
			_ = allFlags

			flags = strings.Replace(flags, "h", "", -1)
			flags = strings.Replace(flags, "j", "", -1)
			flags = strings.Replace(flags, "l", "", -1)

			if j := strings.Index(flags, "#0"); j >= 0 && verb == 'x' {
				k := j + 2
				for k < len(flags) && '0' <= flags[k] && flags[k] <= '9' {
					k++
				}
				n, _ := strconv.Atoi(flags[j+2 : k])
				flags = flags[:j+2] + fmt.Sprint(n-2) + flags[k:]
			}

			narg += strings.Count(allFlags, "*")

			convert := ""
			switch verb {
			default:
				fprintf(fx.Span, "unrecognized format %s%c", flags, verb)
				buf.WriteString("%")
				buf.WriteString(flags)
				buf.WriteString(string(verb))

			case 'f', 'e', 'g', 's', 'c', 'p':
				// usual meanings
				buf.WriteString(flags)
				buf.WriteString(string(verb))

			case 'x', 'X', 'o', 'd', 'b', 'i':
				if verb == 'i' {
					verb = 'd'
				}
				buf.WriteString(flags)
				buf.WriteString(string(verb))

			case 'u':
				buf.WriteString(flags)
				buf.WriteString("d")
				if narg >= len(args) {
					break
				}
				arg := args[narg]
				if t := arg.XType.Def(); t != nil {
					switch t.Kind {
					case Int8:
						convert = "uint8"
					case Int16:
						convert = "uint16"
					case Int32:
						convert = "uint32"
					case Int64:
						convert = "uint64"
					case Int:
						convert = "uint"
					}
				}
			}

			if convert != "" && narg < len(args) {
				arg := args[narg]
				args[narg] = &cc.Expr{Op: cc.Call, Left: &cc.Expr{Op: cc.Name, Text: convert}, List: []*cc.Expr{arg}, XType: stringType}
			}

			narg++
		}
		buf.WriteString(format[start:])
		fx.Texts[j] = strconv.Quote(buf.String())
	}

	return args
}

func fixFormatter(fn *cc.Decl) {
	// Find va_arg assignment.
	var arg *cc.Expr
	var ps []*cc.Expr
	cc.Preorder(fn.Body, func(x cc.Syntax) {
		switch x := x.(type) {
		case *cc.Expr:
			if x.Op == cc.Name && strings.HasPrefix(x.Text, "bigP") {
				ps = append(ps, x)
			}
		case *cc.Stmt:
			stmt := x
			if stmt.Op != cc.StmtExpr {
				return
			}
			expr := stmt.Expr
			if expr.Op != cc.Eq {
				return
			}
			if expr.Left.Op == cc.Name && strings.HasPrefix(expr.Left.Text, "bigP") {
				stmt.Op = cc.Empty
				stmt.Expr = nil
				return
			}
			if expr.Op != cc.Eq || expr.Right.Op != cc.VaArg {
				return
			}
			if arg != nil {
				fprintf(fn.Span, "multiple va_arg in formatter")
			}
			arg = expr.Left
			stmt.Op = cc.Empty
		}
	})

	fp := fn.Type.Decls[0]
	fp.Type = stringType
	fn.Type.Base = stringType
	if arg != nil {
		fn.Type.Decls[0] = arg.XDecl
	} else {
		if len(fn.Type.Decls) == 1 {
			fprintf(fn.Span, "missing va_arg in formatter")
			return
		}
		fn.Type.Decls = fn.Type.Decls[1:]
		decl := &cc.Stmt{
			Op:   cc.StmtDecl,
			Decl: fp,
		}
		fn.Body.Block = append([]*cc.Stmt{decl}, fn.Body.Block...)
	}

	if strings.HasPrefix(fn.Name, "Dconv") && len(ps) > 0 {
		pd := &cc.Decl{Name: "p", Type: ps[0].XDecl.Type}
		fd := &cc.Decl{Name: "flag", Type: intType}
		for _, p := range ps {
			p.XDecl = pd
		}
		fn.Type.Decls = []*cc.Decl{
			pd,
			fd,
			arg.XDecl,
		}
	}
	if len(fn.Name) == 5 && strings.HasSuffix(fn.Name, "conv") {
		switch fn.Name[0] {
		case 'B', 'E', 'F', 'H', 'J', 'N', 'O', 'Q', 'S', 'T', 'V', 'Z':
			fn.Type.Decls = append(fn.Type.Decls, &cc.Decl{
				Name: "flag",
				Type: intType,
			})
		}
	}

	cc.Preorder(fn.Body, func(x cc.Syntax) {
		switch x := x.(type) {
		case *cc.Stmt:
			if arg != nil && x.Op == cc.StmtDecl && x.Decl == arg.XDecl {
				x.Decl = fp
			}

			if x.Op == cc.Return && x.Expr != nil && x.Expr.Text == "0" {
				x.Expr = &cc.Expr{Op: cc.Name, Text: fp.Name, XDecl: fp}
			}

		case *cc.Expr:
			if x.Op == cc.Arrow && x.Text == "flags" {
				x.Op = cc.Name
				x.Text = "flag"
				x.XDecl = nil
				x.Left = nil
			}
		}
	})

}

func fixFormatStmt(fn *cc.Decl, x *cc.Stmt) {
	if x.Op == cc.StmtDecl && GoString(x.Decl.Type) == "Fmt" {
		x.Decl.Type = stringType
		return
	}
}
