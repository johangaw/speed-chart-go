package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedTest struct {
	Time        time.Time
	Latency     time.Duration
	Download    float64
	UploadSpeed float64
}

func main() {

	intervalPrt := flag.Int64("interval", 60*5, "Time between each measurement in seconds")
	output := flag.String("output", "./internet_speed.csv", "Path to the file where the result will be saved")
	flag.Parse()

	writeHeaderLine(*output)

	for {
		fmt.Print("Running...")
		speedTest, err := runTest()
		if err != nil {
			panic(err)
		}
		writeToFile(speedTest, *output)

		fmt.Println("done")
		time.Sleep(time.Duration((*intervalPrt) * 1000_000_000))
	}

}

func runTest() (SpeedTest, error) {
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		return SpeedTest{
			Time:        time.Now(),
			Latency:     s.Latency,
			Download:    s.DLSpeed,
			UploadSpeed: s.ULSpeed,
		}, nil
	}
	return SpeedTest{}, errors.New("No target servers found")
}

func writeHeaderLine(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		if _, err = file.WriteString("Time\tLatency\tDownload speed\tUpload speed"); err != nil {
			panic(err)
		}
	}

}

func writeToFile(speedTest SpeedTest, path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	csvLine := fmt.Sprintf("\n%s\t%s\t%f\t%f", speedTest.Time.Format(time.RFC3339), speedTest.Latency, speedTest.Download, speedTest.UploadSpeed)
	if _, err = file.WriteString(csvLine); err != nil {
		panic(err)
	}
}
