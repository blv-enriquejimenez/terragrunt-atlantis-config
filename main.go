package main

import "github.com/transcend-io/terragrunt-atlantis-config/cmd"

var (
	VERSION = "0.3.1"
)

func main() {
	cmd.Execute(VERSION)
}
