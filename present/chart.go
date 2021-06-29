package present

import (
	"bufio"
	"bytes"
	_ "embed"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/johangus/speed-chart-go/openinbrowser"
)

type Series struct {
	Heading string
	Points  []Point
}

type Point struct {
	X, Y string
}

//go:embed chart.html
var templateString string

func ShowChart(dataFile string, outputPath string) {
	series := parseDataSeries(dataFile)
	htmlPage := buildHTMLPage(series)
	savePage(htmlPage, outputPath)
	openinbrowser.Open(outputPath)
}

func parseDataSeries(path string) []Series {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	headerSplit := strings.Split(scanner.Text(), "\t")
	series := make([]Series, len(headerSplit)-1)

	for index := range series {
		series[index].Heading = headerSplit[index+1]
	}

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "\t")
		time := split[0]
		for index := range series {
			serie := series[index]
			series[index].Points = append(serie.Points, Point{time, split[index+1]})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return series
}

func buildHTMLPage(series []Series) string {
	type plotlyData struct {
		X    []string `json:"x"`
		Y    []string `json:"y"`
		Type string   `json:"type"`
		Name string   `json:"name"`
	}

	templ := template.Must(template.New("chart").Parse(templateString))
	buf := new(bytes.Buffer)

	data := make([]plotlyData, len(series))

	for index, serie := range series[1:] {
		var x []string
		var y []string

		for _, point := range serie.Points {
			x = append(x, point.X)
			y = append(y, point.Y)
		}

		data[index] = plotlyData{
			X:    x,
			Y:    y,
			Type: "scatter",
			Name: serie.Heading,
		}
	}

	err := templ.Execute(buf, data)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func savePage(htmlPage string, path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(htmlPage)
}
