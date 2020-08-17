package main

import (
	"fmt"
	"os"

	"jirago/cmd"
	"jirago/lib/client"
	"jirago/lib/config"
)

func main() {
	if err := config.Setup(); err != nil {
		fmt.Println("Command Failed", err)
		os.Exit(1)
	}
	client.Setup()
	cmd.Execute()
}
