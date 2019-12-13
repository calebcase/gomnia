package histogram

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/guptarohit/asciigraph"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	cmdPlot "github.com/calebcase/gomnia/cmd/plot"
)

var (
	Format  string = "ascii"
	Variant string = "default"

	Width  float64 = 20
	Height float64 = 20
)

func init() {
	cmdPlot.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()

	flags.StringVarP(&Format, "format", "f", Format, "output format [ascii, eps, jpg, jpeg, pdf, png, svg, tif, tiff]")
	flags.StringVar(&Variant, "variant", Variant, "variant")

	flags.Float64Var(&Width, "width", Width, "width in lines/centimeters")
	flags.Float64Var(&Height, "height", Height, "height in characters/centimeters")
}

var Cmd = &cobra.Command{
	Use:   "histogram",
	Short: "plot a histogram",
	Long: `Plot a histogram.

Data should already be in histogram shape (x by count; see generate histogram
sub-command).

https://en.wikipedia.org/wiki/Histogram`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		centroids := plotter.XYs{}
		scanner := bufio.NewScanner(os.Stdin)

		var x, y float64
		var xs, ys []float64
		ys = append(ys, 0)
		for scanner.Scan() {
			_, err = fmt.Sscanf(scanner.Text(), "%g %g", &x, &y)
			if err != nil {
				return
			}

			centroids = append(centroids, plotter.XY{
				X: x,
				Y: y,
			})

			xs = append(xs, x)
			ys = append(ys, y)
		}
		err = scanner.Err()
		if err != nil {
			return
		}

		var of *os.File

		if cmdPlot.OutputPath == "-" {
			of = os.Stdout
		} else {
			of, err = os.Create(cmdPlot.OutputPath)
			if err != nil {
				return
			}
		}

		flags := command.Flags()
		if flags.Changed("output") && !flags.Changed("format") {
			ext := filepath.Ext(cmdPlot.OutputPath)
			if ext != "" {
				Format = ext[1:]
			}
		}

		switch Format {
		case "ascii":
			switch Variant {
			case "default":
				fallthrough
			case "horizontal":
				w := tabwriter.NewWriter(of, 0, 0, 1, ' ', 0)

				for _, point := range centroids {
					fmt.Fprintf(w, "%g\t:\t%s %.0f\n", point.X, renderHBar(rescale(point.Y, floats.Min(ys), floats.Max(ys), 0, 40), true), point.Y)
				}

				w.Flush()
			case "vertical":
				vbars := []float64{}

				for _, y := range ys {
					vbars = append(vbars, 0, y, y, y)
				}
				vbars = append(vbars, 0, 0, 0, 0, 0)

				fmt.Fprintln(of, asciigraph.Plot(vbars, asciigraph.Height(int(Height))))
			default:
				err = errors.New("invalid variant specified")
				return
			}
		default:
			var p *plot.Plot
			p, err = plot.New()
			if err != nil {
				return
			}

			p.Title.Text = "histogram"
			p.X.Label.Text = "x"
			p.Y.Label.Text = "count"

			var h *plotter.Histogram
			h, err = plotter.NewHistogram(centroids, len(centroids))
			if err != nil {
				return
			}
			p.Add(h)

			var w io.WriterTo
			w, err = p.WriterTo(vg.Length(Width)*vg.Centimeter, vg.Length(Height)*vg.Centimeter, Format)
			if err != nil {
				return
			}

			_, err = w.WriteTo(of)
			if err != nil {
				return
			}
		}

		return
	},
}

func rescale(x, xMin, xMax, tMin, tMax float64) float64 {
	return tMin + ((x-xMin)*(tMax-tMin))/(xMax-xMin)
}

func renderHBar(x float64, eigths bool) string {
	i, f := math.Modf(x)

	is := "▇"
	fs := ""

	if eigths {
		is = "█"

		switch {
		case f <= 1./8:
			fs = "▏"
		case f <= 2./8:
			fs = "▎"
		case f <= 3./8:
			fs = "▍"
		case f <= 4./8:
			fs = "▌"
		case f <= 5./8:
			fs = "▋"
		case f <= 6./8:
			fs = "▊"
		case f <= 7./8:
			fs = "▉"
		default:
			fs = "█"
		}
	}

	return strings.Repeat(is, int(i)) + fs
}
