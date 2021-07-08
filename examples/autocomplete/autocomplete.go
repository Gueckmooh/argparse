package main

import (
	"github.com/gueckmooh/argparse"
)

// Run this as `go run commands.go [start|stop]`
func main() {
	// Create new parser object
	parser := argparse.NewParser("commands", "Simple example of argparse commands")

	_ = parser.Flag("v", "verbose", &argparse.Options{Help: "Enable verbose mode"})
	_ = parser.String("s", "string", &argparse.Options{Help: "Enable verbose mode"})

	// Add top level command `start`
	startCmd := parser.NewCommand("start", "Will start a process")

	_ = startCmd.Flag("", "wiggle", &argparse.Options{Help: "Makes the dog to wiggle its tail"})
	_ = startCmd.String("w", "woof", &argparse.Options{Help: "Makes the dog to wiggle its tail"})

	// Add top level commands `stop`
	_ = parser.NewCommand("stop", "Will stop a process")

	parser.GenBashAutoComplete()

	// Parse command line arguments and in case of any error print error and help information
	// err := parser.Parse(os.Args)
	// if err != nil {
	// 	fmt.Print(parser.Usage(err))
	// 	return
	// }

	// // Check if `start` command was given
	// if startCmd.Happened() {
	// 	// Starting a process
	// 	fmt.Println("Started process")
	// } else if stopCmd.Happened() { // Check if `stop` command was given
	// 	// Stopping a process
	// 	fmt.Println("Stopped process")
	// } else {
	// 	// In fact we will never hit this one
	// 	// because commands and sub-commands are considered as required
	// 	err := fmt.Errorf("bad arguments, please check usage")
	// 	fmt.Print(parser.Usage(err))
	// }
}
