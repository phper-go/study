package main

import (
	"os"
	_ "study/expamle/controller"
	"study/expamle/run"

	"github.com/phper-go/frame"
)

func main() {

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "start")
	}

	frame.Run(run.App, os.Args)
}
