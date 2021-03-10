package easy_git

import (
	"fmt"
	"os/exec"
)

// Command holds the path of the git executable, remote repo, and a pointer to a cmd
type Command struct {
	GitPath    string
	Repository string
	Cmd        *exec.Cmd
}

// Input determines which commands should be executed
type Input struct {
	Initalize   bool
	AddFiles    bool
	CommitFiles bool
	AddOrigin   bool
	PushFiles   bool
}

// Init starts the program
func (c *Command) Init() {
	var input Input

	if !StrToBool(ReadInput("Yes to all", true)) {
		input.Initalize = StrToBool(ReadInput("Initalize New Repository", true))
		input.AddFiles = StrToBool(ReadInput("Add All Files", true))
		input.CommitFiles = StrToBool(ReadInput("Commit Files", true))
		input.AddOrigin = StrToBool(ReadInput("Add Origin", true))
		input.PushFiles = StrToBool(ReadInput("Push Files", true))
	} else {
		input.Initalize = true
		input.AddFiles = true
		input.CommitFiles = true
		input.AddOrigin = true
		input.PushFiles = true
	}

	c.Repository = ReadInput("Enter repository", false)

	if input.Initalize {
		c.RunCommand(exec.Command(c.GitPath, []string{"init"}...))
	}

	//c.RunCommand(exec.Command("git", []string{"branch", "-M", "main"}...))

	if input.AddOrigin {
		c.RunCommand(exec.Command(c.GitPath, []string{
			"remote",
			"set-url",
			c.Repository,
		}...))
	}

	if input.AddFiles {
		c.RunCommand(exec.Command(c.GitPath, []string{
			"add",
			".",
		}...))
	}

	if input.CommitFiles {
		c.RunCommand(exec.Command(c.GitPath, []string{
			"commit",
			"-m",
			"testabcdfg",
		}...))
	}

	if input.PushFiles {
		c.RunCommand(exec.Command(c.GitPath, []string{
			"push",
			"-u",
			"origin",
			"main",
		}...))
	}
}

// StrToBool parses the user input and converts it to a boolean
func StrToBool(str string) bool {
	if str == "Y" || str == "y" {
		return true
	}

	return false
}

// RunCommand starts a command and prints the result
func (c *Command) RunCommand(cmd *exec.Cmd) error {
	bytes, err := cmd.Output()

	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}

	fmt.Println(string(bytes))
	return nil
}

// ReadInput reads user input through stdin
func ReadInput(message string, boolAnswer bool) string {
	var input string

	if boolAnswer {
		fmt.Printf("%s y/n: ", message)
	} else {
		fmt.Printf("%s: ", message)
	}

	fmt.Scanln(&input)

	return input
}
