package main

import (
	"fmt"
	"os"

	"github.com/gueckmooh/argparse"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("positionals", "Prints provided positional strings to stdout")
	// Create string flag
	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print"})
	p := parser.PosString("somename", &argparse.Options{Help: "Positional strings to print"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	// Print the -s string
	fmt.Println(*s)
	// Print any remaining strings as the positional arg
	fmt.Println(*p)
}
