package histogram

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"

	"github.com/calebcase/gomnia/cmd/summarize"
)

var (
	N int = 10
)

func init() {
	summarize.Cmd.AddCommand(Cmd)

	flags := Cmd.PersistentFlags()
	flags.IntVarP(&N, "n", "n", N, "number of bins")
}

var Cmd = &cobra.Command{
	Use:   "histogram",
	Short: "summarize data with a histogram",
	Long: `Summarize data with a histogram.

NOTE: This will read all data into memory. Histograms for large sets of data
can be approximated with the t-digest sub-command.

https://en.wikipedia.org/wiki/Histogram`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		vs := []float64{}
		scanner := bufio.NewScanner(os.Stdin)

		var v float64
		for scanner.Scan() {
			v, err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				return
			}

			vs = append(vs, v)
		}
		err = scanner.Err()
		if err != nil {
			return
		}

		sort.Float64s(vs)
		dividers := Fence(vs)
		hist := stat.Histogram(nil, dividers, vs, nil)

		for i := 0; i < len(hist); i++ {
			fmt.Printf("%g %g\n", dividers[i], hist[i])
		}

		return
	},
}

func Fence(vs []float64) []float64 {
	dividers := make([]float64, N+1)
	min := floats.Min(vs)

	// Get the next, next float; This is necessary since the max value in
	// the set must be strictly less than the last fence value.
	max := math.Nextafter(math.Nextafter(floats.Max(vs), math.MaxFloat64), math.MaxFloat64)

	_ = floats.Span(dividers, min, max)

	return dividers
}
