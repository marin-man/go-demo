package main

import (
	"helloworld/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail:", err.Error())
		os.Exit(-1)
	}
}