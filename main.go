package main

import (
	"fmt"
	"os"

	"github.com/oucema001/task/cmd"
	"github.com/oucema001/task/data"
)

func main() {
	//data.Init()
	//cmd.Execute()

	must(data.Init())
	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
