package easy_git

import (
	"fmt"
	"log"
	"os/exec"
)

type Command struct {
	Repository string
	Cmd        *exec.Cmd
}

type Input struct {
	Initalize   bool
	AddFiles    bool
	CommitFiles bool
	AddOrigin   bool
	PushFiles   bool
}

func (c *Command) Init() {
	var input Input

	input.Initalize = StrToBool(ReadInput("Initalize New Repository", true))
	input.AddFiles = StrToBool(ReadInput("Add All Files", true))
	input.CommitFiles = StrToBool(ReadInput("Commit Files", true))
	input.AddOrigin = StrToBool(ReadInput("Add Origin", true))
	input.PushFiles = StrToBool(ReadInput("Push Files", true))
	c.Repository = ReadInput("Enter repository", false)

	if input.Initalize {
		c.RunCommand(exec.Command("/usr/bin/git", []string{"init"}...))
	}

	//c.RunCommand(exec.Command("git", []string{"branch", "-M", "main"}...))

	if input.AddFiles {
		c.RunCommand(exec.Command("/usr/bin/git", []string{"add", "."}...))
	}

	if input.CommitFiles {
		c.RunCommand(exec.Command("/usr/bin/git", []string{"commit", "-m", "testabcdfg"}...))
	}

	if input.AddOrigin {
		c.RunCommand(exec.Command("/usr/bin/git", []string{"remote", "add", "origin", c.Repository}...))
	}

	if input.PushFiles {
		c.RunCommand(exec.Command("/usr/bin/git", []string{"push", "-u", "origin", "main"}...))
	}

	fmt.Println(input)
	fmt.Println(c)
}

func StrToBool(str string) bool {
	if str == "Y" || str == "y" {
		return true
	}

	return false
}

func (c *Command) RunCommand(cmd *exec.Cmd) error {
	fmt.Println(cmd)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}

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
