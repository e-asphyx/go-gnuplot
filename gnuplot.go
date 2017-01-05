package gnuplot

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	Title  string
	XLabel string
	YLabel string
	XScale Scale // Linear or logarithmic only
	YScale Scale
	Grid   bool
	Data   []*Plot2Data
}

type Cmd struct {
	cmd      *exec.Cmd
	tmpFiles []*os.File
	term     chan struct{}
}

func (c *Cmd) Close() error {
	close(c.term)
	err := c.cmd.Wait()
	for _, f := range c.tmpFiles {
		os.Remove(f.Name())
	}
	return err
}

func (p *Plot2D) Exec() (*Cmd, error) {
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

	tmpFiles := make([]*os.File, 0, len(p.Data))

	// Prepare plot command
	var plotspec string
	for _, pd := range p.Data {

		// Create data file
		dataFile, err := ioutil.TempFile("", "plotdata")
		if err != nil {
			return nil, err
		}

		for i := 0; i < pd.Data.Len(); i++ {
			val := pd.Data.At(i)
			fmt.Fprintf(dataFile, "%f %f\n", val[0], val[1])
		}
		dataFile.Close()

		tmpFiles = append(tmpFiles, dataFile)

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

		plotspec += fmt.Sprintf("\"%s\" using %s %s with %s", dataFile.Name(), usingspec, titlespec, style)

		if pd.Color != nil {
			plotspec += fmt.Sprintf(" linecolor rgb \"%s\"", pd.Color.Color())
		}
	}

	fmt.Fprintf(&script, "plot %s\n", plotspec)
	fmt.Fprintf(&script, "pause -1\n")

	// Prepagre Gnuplot child
	cmd := Cmd{
		cmd:      exec.Command("gnuplot"),
		term:     make(chan struct{}),
		tmpFiles: tmpFiles,
	}
	cmd.cmd.Stderr = os.Stderr
	pipe, err := cmd.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.cmd.Start()

	if err != nil {
		return nil, err
	}

	go func() {
		io.Copy(pipe, &script)
		<-cmd.term
		pipe.Close()
	}()

	return &cmd, err
}
