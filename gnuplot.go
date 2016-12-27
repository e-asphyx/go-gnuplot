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

type Dataset struct {
	Title string
	Style Style
	Color Color
	Data  [][2]float64
}

type Plot2D struct {
	Title    string
	XLabel   string
	YLabel   string
	XLog     bool
	YLog     bool
	Grid     bool
	Datasets []*Dataset
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

	if p.XLog || p.YLog {
		var axes string
		if p.XLog {
			axes = "x"
		}
		if p.YLog {
			axes += "y"
		}
		fmt.Fprintf(&script, "set logscale %s\n", axes)
	}

	if p.Grid {
		fmt.Fprintf(&script, "set grid\n")
	}

	// Prepare plot command
	var plotspec string
	for _, ds := range p.Datasets {
		if plotspec != "" {
			plotspec += ", "
		}

		var titlespec string
		if ds.Title != "" {
			titlespec = fmt.Sprintf("title \"%s\"", ds.Title)
		} else {
			titlespec = "notitle"
		}

		style := ds.Style
		if style == "" {
			style = StyleLines
		}
		plotspec += fmt.Sprintf("\"-\" using 1:2 %s with %s", titlespec, style)

		if ds.Color != nil {
			plotspec += fmt.Sprintf(" linecolor \"%s\"", ds.Color.Color())
		}
	}

	fmt.Fprintf(&script, "plot %s\n", plotspec)

	for _, ds := range p.Datasets {
		for _, val := range ds.Data {
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
