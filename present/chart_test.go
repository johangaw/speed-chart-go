package present

import (
	"strings"
	"testing"
)

func Test_buildHTMLPage(t *testing.T) {
	type args struct {
		series []Series
	}
	tests := []struct {
		name          string
		args          args
		wantToContain []string
	}{
		{
			name: "It should include all points as X and Y data",
			args: args{series: []Series{
				{Heading: "Latency", Points: []Point{{"2021-06-24", "1ms"}, {"2021-06-25", "2ms"}, {"2021-06-26", "3ms"}}},
				{Heading: "Down", Points: []Point{{"2021-06-24", "24"}, {"2021-06-25", "25"}, {"2021-06-26", "26"}}},
			}},
			wantToContain: []string{
				`"y":["24","25","26"]`,
				`"x":["2021-06-24","2021-06-25","2021-06-26"]`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildHTMLPage(tt.args.series)
			for _, subString := range tt.wantToContain {
				if !strings.Contains(got, subString) {
					t.Errorf("buildHTMLPage() = %v, want %v", got, subString)
				}
			}
		})
	}
}
