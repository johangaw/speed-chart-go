package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/johangus/speed-chart-go/measure"
	"github.com/johangus/speed-chart-go/present"
)

func main() {

	measureCmd := flag.NewFlagSet("measure", flag.ExitOnError)
	intervalPrt := measureCmd.Int64("interval", 60*5, "Time between each measurement in seconds")
	measureOutput := measureCmd.String("output", "./internet_speed.csv", "Path to the file where the result will be saved")

	presentCmd := flag.NewFlagSet("present", flag.ExitOnError)
	dataFilePathPtr := presentCmd.String("path", "./internet_speed.csv", "Path to the data file created by the 'measure' command")
	presentOutputPtr := presentCmd.String("output", "./outfile-chart.html", "Path to the html chart file")

	if len(os.Args) < 2 {
		fmt.Println("expected 'measure' or 'present' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "measure":
		measureCmd.Parse(os.Args[2:])
		measure.Measure(*measureOutput, *intervalPrt)
	case "present":
		presentCmd.Parse(os.Args[2:])
		present.ShowChart(*dataFilePathPtr, *presentOutputPtr)
	default:
		fmt.Println("expected 'measure' or 'present' subcommands")
		os.Exit(1)
	}

}
