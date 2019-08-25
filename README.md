# Advent of Code -- Go

## Architecture

This program uses Go's [plugin](https://golang.org/pkg/plugin) to dynamically load the solution code for
a given day's puzzle. The [puzzles](puzzles) directory contains subdirectories for each year of AOC
puzzles, each of which contains subdirectories for each day's puzzle. Within a given day's 
directory, there is a `main.go` and an `input.txt` file. `main.go` implements a `Solve()` function
with the signature `Solve(fname string)`, which is loaded in by `AOC.go`. `Solve()` calls functions
inside of the puzzle's `main.go` file to parse the input provided and solve the puzzle, and then
prints out the solution.

Because of Go's plugin system, each plugin is its own `main` package that is built with 

```sh
$ go build -buildmode=plugin
```

to produce a `.so` (shared object) file. Because each plugin has to be its own `main` package, I
implement an empty `main()` function in each so that Go doesn't complain that `function main is 
undeclared in the main package`.

Within `AOC.go`, I parse and validate the command line arguments to get the year and day of the
desired puzzle. If a solution has been implemented (i.e. if there is a `.so` file in the puzzle's
directory), I load it in and use `plugin.Lookup("Solve")` to be able to use that function. The 
caller function must provide a type assertion for the looked-up symbol, which in this case is
`func(string)`. Once that's all taken care of, we can finally call the function, which in turn calls
whatever functions the plugin has implemented to read the input file and solve the puzzle.

## Usage

First, build all of the `.so` files for the puzzle solutions:

```sh
$ bash build_all_plugins.sh
```

Then, build the binary for the main `AOC` program:

```sh
$ go build -o AOC
```

Finally, run `AOC` with the desired year and day:

```sh
$ ./AOC 2018 1
```
