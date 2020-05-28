package main

import "os"
import "fmt"
import "flag"

func main() {
	// Command-line arguments are a common way to parameterize execution of programs. For example, go run hello.go uses run and hello.go arguments to the go program.
	argsWithProg := os.Args        // os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program.
	argsWithoutProg := os.Args[1:] // os.Args[1:] holds the arguments to the program.
	arg := os.Args[3]              // You can get individual args with normal indexing.

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// Command-Line Flags
	// Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.
	wordPtr := flag.String("word", "foo", "a string") // Basic flag declarations are available for string, integer, and boolean options. This flag.String function returns a string pointer (not a string value). Here we declare a string flag word with a default value "foo" and a short description.
	numbPtr := flag.Int("numb", 42, "an int")         // This declares numb and fork flags, using a similar approach to the word flag.
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") // It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
	flag.Parse()                                         // Once all flags are declared, call flag.Parse() to execute the command-line parsing.

	fmt.Println("word:", *wordPtr) // Here we’ll just dump out the parsed options and any trailing positional arguments. Note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// Commnad
// $ go build command-line-arguments.go
// $ ./command-line-arguments a b c d

// Command
// $ go build command.line.args.go
// $ ./command.line.args -word=opt -numb=7 -fork -svar=flag // Try out the built program by first giving it values for all flags.
// word: opt
// numb: 7
// fork: true
// svar: flag
// tail: []

// $ ./command.line.args -word=opt // Note that if you omit flags they automatically take their default values.
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: []

// $ ./command.line.args -word=opt a1 a2 a3 // Trailing positional arguments can be provided after any flags.
// word: opt
// ...
// tail: [a1 a2 a3]

// $ ./command.line.args -word=opt a1 a2 a3 -numb=7 // Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]

// $ ./command.line.args -h // Use -h or --help flags to get automatically generated help text for the command-line program.
// Usage of ./command.line.args:
//   -fork=false: a bool
//   -numb=42: an int
//   -svar="bar": a string var
//   -word="foo": a string

// $ ./command.line.args -wat // If you provide a flag that wasn’t specified to the flag package, the program will print an error message an show the help text again.
// flag provided but not defined: -wat
// Usage of ./command.line.args:
// ...

// $ go build command.line.args.go
// $ ./command.line.args -word=opt -numb=7 -fork -svar=flag // Try out the built program by first giving it values for all flags.
// word: opt
// numb: 7
// fork: true
// svar: flag
// tail: []

// $ ./command.line.args -word=opt // Note that if you omit flags they automatically take their default values.
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: []

// $ ./command.line.args -word=opt a1 a2 a3 // Trailing positional arguments can be provided after any flags.
// word: opt
// ...
// tail: [a1 a2 a3]

// $ ./command.line.args -word=opt a1 a2 a3 -numb=7 // Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments).
// word: opt
// numb: 42
// fork: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]

// $ ./command.line.args -h // Use -h or --help flags to get automatically generated help text for the command-line program.
// Usage of ./command.line.args:
//   -fork=false: a bool
//   -numb=42: an int
//   -svar="bar": a string var
//   -word="foo": a string

// $ ./command.line.args -wat // If you provide a flag that wasn’t specified to the flag package, the program will print an error message an show the help text again.
// flag provided but not defined: -wat
// Usage of ./command.line.args:
// ...
