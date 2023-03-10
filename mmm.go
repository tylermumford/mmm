// This is an early copy of mmm.py
// for the purposes of comparing both implementations

package main

import "fmt"
import "os"
import "strings"
import "github.com/BurntSushi/toml"

const FILE = "mmm.toml"

type Config struct {
	Version string
	Command map[string]Command
}

// Returns the commands available for use.
func (c *Config) Commands() []string {
	result := []string{}
	for i := range c.Command {
		result = append(result, i)
	}
	return result
}

type Command struct {
	Script any
}

func main() {
	// 1. Load

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

	// 2. Get input

	if len(os.Args) != 2 {
		help := strings.Join(config.Commands(), ", ")
		fmt.Printf("commands in %s: %s\n", FILE, help)
		os.Exit(0)
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

	// 3. Get steps

	scriptLines := []string{}
	if s, ok := script.(string); ok {
		scriptLines = append(scriptLines, s)
	} else if ss, ok := script.([]string); ok {
		scriptLines = ss
	} else if si, ok := script.([]interface{}); ok {
		for i := range si {
			if s, ok := si[i].(string); ok {
				scriptLines = append(scriptLines, s)
			} else {
				fmt.Printf("script array must contain strings, but found %T\n", si[i])
				os.Exit(7)
			}
		}
	} else {
		fmt.Printf("script has unhandled Go type: %T\n", script)
		os.Exit(6)
	}

	// 4. Act

	for i := range scriptLines {
		fmt.Println(scriptLines[i])
	}
}
