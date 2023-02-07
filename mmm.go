// This is an early copy of mmm.py
// for the purposes of comparing both implementations

package main

import "fmt"
import "os"
import "github.com/BurntSushi/toml"

const FILE = "mmm.toml"

type Config struct {
	Version string
	Command map[string]Command
}

type Command struct {
	Script any
}

func main() {
	data, err := os.ReadFile(FILE)
	if err != nil {
		fmt.Printf("could not read %s:\n%s\n", FILE, err)
		os.Exit(1)
	}

	config := Config{}
	_, err = toml.Decode(string(data), &config)
	if err != nil {
		fmt.Printf("could not decode %s:\n%s\n", FILE, err)
		os.Exit(2)
	}

	if len(os.Args) != 2 {
		fmt.Println("no argument given; please specify a command")
		os.Exit(3)
	}
	arg := os.Args[1]

	command := config.Command[arg]
	if (command == Command{}) {
		fmt.Printf("command '%s' not found in %s\n", arg, FILE)
		os.Exit(4)
	}

	script := command.Script
	if script == "" {
		fmt.Printf("script not defined for '%s' in %s\n", arg, FILE)
		os.Exit(5)
	}

	fmt.Printf("I would run %v\n", script)
}
