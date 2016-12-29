package gnuplot

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

type Style string

var (
	StyleLines       Style = "lines"
	StyleLinespoints Style = "lp"
	StyleDots        Style = "dots"
	StylePoints      Style = "points"
)

type Dataset2D interface {
	Len() int
	At(int) [2]float64
}

type DataY []float64

func (d DataY) Len() int {
	return len(d)
}

func (d DataY) At(i int) [2]float64 {
	return [2]float64{float64(i), d[i]}
}

type DataXY [][2]float64

func (d DataXY) Len() int {
	return len(d)
}

func (d DataXY) At(i int) [2]float64 {
	return d[i]
}

type Plot2Data struct {
	Title string
	Style Style
	Color Color
	Data  Dataset2D
}

type Scale int

const (
	ScaleLin Scale = iota
	ScaleLog
	ScaleDb
)

type Plot2D struct {
	Title    string
	XLabel   string
	YLabel   string
	XScale   Scale // Linear or logarithmic only
	YScale   Scale
	Grid     bool
	PlotData []*Plot2Data
}

func (p *Plot2D) Exec() (*exec.Cmd, error) {
	var script bytes.Buffer

	if p.Title != "" {
		fmt.Fprintf(&script, "set title \"%s\"\n", p.Title)
	}

	if p.XLabel != "" {
		fmt.Fprintf(&script, "set xlabel \"%s\"\n", p.XLabel)
	}

	if p.YLabel != "" {
		fmt.Fprintf(&script, "set ylabel \"%s\"\n", p.YLabel)
	}

	if p.XScale == ScaleLog || p.YScale == ScaleLog {
		var axes string
		if p.XScale == ScaleLog {
			axes = "x"
		}

		if p.YScale == ScaleLog {
			axes += "y"
		}
		fmt.Fprintf(&script, "set logscale %s\n", axes)
	}

	if p.Grid {
		fmt.Fprintf(&script, "set grid\n")
	}

	// Prepare plot command
	var plotspec string
	for _, pd := range p.PlotData {
		if plotspec != "" {
			plotspec += ", "
		}

		var titlespec string
		if pd.Title != "" {
			titlespec = fmt.Sprintf("title \"%s\"", pd.Title)
		} else {
			titlespec = "notitle"
		}

		style := pd.Style
		if style == "" {
			style = StyleLines
		}

		var usingspec string
		if p.YScale == ScaleDb {
			usingspec = "1:(20*log10(column(2)))"
		} else {
			usingspec = "1:2"
		}

		plotspec += fmt.Sprintf("\"-\" using %s %s with %s", usingspec, titlespec, style)

		if pd.Color != nil {
			plotspec += fmt.Sprintf(" linecolor rgb \"%s\"", pd.Color.Color())
		}
	}

	fmt.Fprintf(&script, "plot %s\n", plotspec)

	for _, pd := range p.PlotData {
		for i := 0; i < pd.Data.Len(); i++ {
			val := pd.Data.At(i)
			fmt.Fprintf(&script, "%f %f\n", val[0], val[1])
		}
		fmt.Fprintf(&script, "e\n")
	}
	fmt.Fprintf(&script, "pause mouse\n")

	// Prepagre Gnuplot child
	cmd := exec.Command("gnuplot", "-")
	cmd.Stdin = &script
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	return cmd, err
}
