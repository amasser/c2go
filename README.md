C2go converts (some) C programs to Go.
It was developed by Russ Cox to convert the Go runtime and compiler to Go.
It is progressing toward becoming a general-purpose translation tool,
but it does not (and never will) translate all possible C programs.
You will need to refactor your C program into the subset of C that c2go
can handle, and give it translation hints in the configuration file.

## Goals

 - Generate Go that is readable, maintainable,
   and similar to the original C code.

 - Generate type-safe and memory-safe code;
   it should not use the "unsafe" package.

 - Handle as much as possible of the subset of C that is compatible
   with the other goals.

## Non-Goals

 - Handle *all* C code.

 - Be fully automatic. (It’s OK to need some hints in the config file.)

## Limitations

Some limitations are virtually inherent in the goals of producing safe, readable code;
others are simply things that haven’t been implemented yet.

 - Preprocessor support: The only preprocessor directives that c2go understands
   are `#include` and simple constant `#define`.

 - Pointer arithmetic: Pointer arithmetic is translated as slice operations when possible.
   But slices can only be moved forward, not backward.
   (`p + n` becomes `p[n:]`; `p - n` doesn’t work.)

 - Some constructs that are equivalent in C are only recognized one way by c2go.
   For example, when a count of items is multiplied by a `sizeof` expression, 
   the `sizeof` needs to come last.
   In C, `&p[n]` and `p + n` are equivalent, but c2go translates them differently, 
   as `&p[n]` and as `p[n:]`.

# How to Use c2go

Before you even try running c2go on your program, 
you should refactor it into the kind of C that c2go is likely to be able to understand.
Throughout the refactoring process, 
run tests regularly to make sure you haven’t introduced any bugs.

## Remove Portability Macros

Most C projects rely heavily on macros and `#ifdef`s for portability.
These are hard for c2go to translate, and what’s more, they point out
code that is likely to be problematic: unsafe, low-level, and machine-dependent.
Replace that code with something that will translate cleanly to Go.

If the program has a header like "platform.h" or "port.h", start there.
Keep going till there are no `#ifdef`s left (except header include guards).

## Simplify Memory Management

If you allocate memory with `malloc`, c2go will generally be able to translate 
it to `make` or `new`. But that won’t happen with a custom memory allocator. 
So if the program uses a custom allocator, replace it with `malloc` and `free`.

## Reduce Preprocessor Use

Function-like macros are hard to translate to Go;
c2go doesn’t even try at this point.
You will either need to expand them in place,
or convert them to regular functions.

Other preprocessor tricks (like including a file multiple times as a template,
with different macros defined each time) 
will need to be eliminated as well.

## Create a Config File

The config file is an important part of the process of using c2go.
It can contain various types of translation hints and fixups,
but we’ll start by just specifying what package your Go code will be in.

Create a file called `c2go.cfg`.
If your project is a command, put the line `package * main` in the config file;
otherwise you should replace `main` with your package’s import path.
(You can also specify multiple packages, each with its own pattern that specifies which files go in that package.
But most likely all your files will go in the same package, and one line with a `*` is all you need.)

When you run the c2go tool, always point it to your config file with the `-c` option.
You will probably be adding a lot more to the file before you are done.

## Start Translating

I like to start by translating just one file; 
then once it translates and compiles successfully,
I add another file to the set of files that I’m translating,
until eventually I am translating the whole project.

When you first run c2go, you will probably wonder where the output went.
By default, the Go files are created in subdirectories of `/tmp/c2go`.
You can send them somewhere else (such as your default GOPATH) with the `-dst` option.

## Fix Errors

You will probably have two types of translation errors to deal with:
errors from c2go itself, and errors from the Go compiler when you try to compile the output.
Errors from c2go generally should be fixed by editing the C code to make it easier for
c2go to parse or translate it.

In my experience, the majority (though not all, by any means) of errors from the Go compiler
can be fixed by adding type hints to `c2go.cfg`.
Type hints look like this:

    slice ReplicateValue.table

This indicates that the parameter or local variable `table`
in the function `ReplicateValue` should be translated as a slice instead of as a pointer.
(The default is to translate `char *` as `[]byte`, and other pointer types as pointers.)
You will need a `slice` type hint for each pointer variable
that is used to point to an array of objects rather than a single object.
The other type hints available for pointer variables are `ptr` and `string`.

Besides the type hints for pointers, there is also `bool`
to indicate that an `int` variable should be translated as a `bool`.

The variable reference in a type hint is usually of the form
functionName.variableName or typeName.memberName.
But for global variables, or for variables in blocks nested inside a function,
you just use the variable name by itself.

If your error can’t be fixed with a type hint, 
you can usually fix it by editing the C code.
Make sure it still compiles and passes the tests.

If the error can’t be fixed by editing the C code,
you will need to fix it in the Go output by putting a diff block
in the config file. A diff block replaces one or more
lines of Go output, like this:

```
diff {
-	htrees        []*HuffmanCode
+	htrees        [][]HuffmanCode
}
```
