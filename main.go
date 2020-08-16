package main

import (
	"jirago/cmd"
	"jirago/lib/client"
	"jirago/lib/config"
)

func main() {
	config.Setup()
	client.Setup()
	cmd.Execute()
}
