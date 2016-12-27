package gnuplot

import (
	"testing"
)

func TestPlot2D(t *testing.T) {
	data1 := Dataset{
		Title: "Func 1",
		Style: StyleLines,
		Color: ColorRed,
		Data: [][2]float64{
			{0, 1},
			{1, 2},
			{2, 4},
			{3, 3},
		},
	}

	data2 := Dataset{
		Title: "Func 2",
		Style: StyleLines,
		Color: ColorGreen,
		Data: [][2]float64{
			{0, 5},
			{1, 8},
			{2, 1},
			{3, 2},
		},
	}

	plot := Plot2D{
		Title:    "Title",
		Grid:     true,
		Datasets: []*Dataset{&data1, &data2},
	}

	cmd, err := plot.Exec()
	if err != nil {
		t.Error(err)
	}

	if err := cmd.Wait(); err != nil {
		t.Error(err)
	}
}
