package cli

import (
	"flag"
	"fmt"
	"os"

	internal "github.com/pl-42/hold_my_beer/internal"
)

// cliCommandHandler defines the signature for command handler functions.
type cliCommandHandler func(args []string, stdout, stderr *os.File) int

// cliCommand represents a command with its description and handler function.
type cliCommand struct {
	description string
	handler     cliCommandHandler
}

// Map that associates subcommand with corresponding description and execution function.
var cliCommands = make(map[string]cliCommand)

// Registers a new command with its description and handler function.
// name: The name of the command (e.g., "help", "version").
// description: A brief description of what the command does.
// handler: The function that will be executed when the command is invoked.
func RegisterCommand(name, description string, handler cliCommandHandler) {
	cliCommands[name] = cliCommand{
		description: description,
		handler:     handler,
	}
}

// Run executes the CLI application with the provided arguments.
// args: The command-line arguments passed to the application.
// stdout: The output stream for standard output.
// stderr: The output stream for error messages.
// Returns an exit code indicating success (0) or failure (non-zero).
func Run(args []string, stdout, stderr *os.File) int {
	fs := flag.NewFlagSet("HoldMyBeer", flag.ExitOnError)
	fs.Usage = func() {
		usage(stderr)
		stderr.WriteString("\nAvailable options:\n")
		fs.SetOutput(stderr)
		fs.PrintDefaults()
	}
	beQuiet := fs.Bool("be-quiet", false, "Suppress the startup signature")

	fs.Parse(args)
	if !*beQuiet {
		stderr.WriteString(beerArt())
	}
	a := fs.Args()
	if len(a) == 0 {
		return 0
	}

	cmdName := a[0]
	cmdArgs := []string{}
	if len(a) > 1 {
		cmdArgs = a[1:]
	}

	if cmdFunc, exists := cliCommands[cmdName]; exists {
		return cmdFunc.handler(cmdArgs, stdout, stderr)
	} else {
		stderr.WriteString("Unknown command: " + cmdName + "\n")
		return 1
	}
}

// helpCommand displays the usage information for the CLI application.
func helpCommand(args []string, stdout, stderr *os.File) int {
	usage(stdout)
	return 0
}

// versionCommand displays the version of the HoldMyBeer application.
func versionCommand(args []string, stdout, stderr *os.File) int {
	fs := flag.NewFlagSet("version", flag.ExitOnError)
	short := fs.Bool("short", false, "Display only the version number")
	fs.Parse(args)

	if *short {
		stdout.WriteString(fmt.Sprintf("%s\n", internal.VERSION))
	} else {
		stdout.WriteString(fmt.Sprintf("HoldMyBeer v%s\n", internal.VERSION))
	}
	return 0
}

// usage prints the usage information for the CLI application to the provided stream.
func usage(stream *os.File) {
	stream.WriteString("Usage: HoldMyBeer <command> [options]\n")
	stream.WriteString("Available commands:\n")
	for cmd, cmdInfo := range cliCommands {
		stream.WriteString(fmt.Sprintf("  %s\n\t%s\n", cmd, cmdInfo.description))
	}
}

// init registers the default commands for the CLI application.
func init() {
	// Register the "help" command
	RegisterCommand("help", "Display available commands", helpCommand)
	// Register the "version" command
	RegisterCommand("version", "Display the version of HoldMyBeer", versionCommand)
}
