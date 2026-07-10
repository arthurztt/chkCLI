package main

import "github.com/arthurztt/chkCLI/cmd"

var version = "dev"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
