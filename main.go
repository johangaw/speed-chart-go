package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/johangus/speed-chart-go/measure"
)

func main() {

	measureCmd := flag.NewFlagSet("measure", flag.ExitOnError)
	intervalPrt := flag.Int64("interval", 60*5, "Time between each measurement in seconds")
	output := flag.String("output", "./internet_speed.csv", "Path to the file where the result will be saved")

	if len(os.Args) < 2 {
		fmt.Println("expected 'measure' or 'present' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "measure":
		measureCmd.Parse(os.Args[2:])
		measure.Measure(*output, *intervalPrt)
	case "present":
		panic("Not yet implemented")
	default:
		fmt.Println("expected 'measure' or 'present' subcommands")
		os.Exit(1)
	}

}
