// Package cli provides CLI application utils
package cli

import (
	"fmt"
	"os"
)

// Command represents CLI command
type Command struct {
	Name  string
	Help  string
	Usage string
	Args  Args
	Run   func(args []string)
}

// App represents CLI application
type App struct {
	Name     string
	Help     string
	commands []Command
}

// RunWith runs CLI application with provided commands
func (a *App) RunWith(commands ...Command) {
	a.commands = make([]Command, 0, len(commands)+1)
	a.commands = append(a.commands, commands...)
	a.commands = append(a.commands, Command{
		Name:  "help",
		Help:  "Show this help message",
		Usage: "",
		Run: func(_ []string) {
			a.printHelp()
		},
	})
	args := os.Args[1:]
	if len(args) == 0 {
		a.die("No command provided")
	}
	for _, cmd := range a.commands {
		if cmd.Name == args[0] {
			if len(args)-1 < len(cmd.Args) {
				a.die("Not enough arguments")
			} else if len(args)-1 > len(cmd.Args) {
				a.die("Too many arguments")
			}
			args = args[1:]
			UnpackArgs(args, cmd.Args)
			cmd.Run(args)
			os.Exit(0)
		}
	}
	a.die("Unknown command")
}

func (a *App) die(message string) {
	print(Red(message) + "\n")
	a.printUsage()
	os.Exit(1)
}

func (a *App) printHelp() {
	fmt.Println(a.Help + "\n")
	a.printUsage()
}

func (a *App) printUsage() {
	usage := "Usage:\n"
	usage += fmt.Sprintf("  %s <command> [arguments]\n", a.Name)
	usage += "The commands are:\n"
	for _, cmd := range a.commands {
		usage += fmt.Sprintf("  %s %s", cmd.Name, cmd.Usage)
		if len(cmd.Help) > 0 {
			usage += fmt.Sprintf(" — %s", cmd.Help)
		}
		usage += "\n"
	}
	fmt.Print(usage)
}
