package main

import easy_git "github.com/DustiasTheGuy/easy-git/easy-git"

func main() {

	cmd := easy_git.Command{
		Repository: "https://github.com/DustiasTheGuy/easy-git.git",
	}

	cmd.Init()
}
