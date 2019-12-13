package tdigest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/caio/go-tdigest"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat"

	"github.com/calebcase/gomnia/cmd/summarize/histogram"
)

var (
	Compression float64 = 100
)

func init() {
	histogram.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Compression, "compression", "c", Compression, "compression")
}

var Cmd = &cobra.Command{
	Use:   "t-digest",
	Short: "approximate histogram with t-digest method",
	Long: `Approximate histogram with t-digest method.

https://github.com/tdunning/t-digest/blob/master/docs/t-digest-paper/histo.pdf`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		digest, err := tdigest.New(tdigest.Compression(Compression))
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(os.Stdin)

		var v float64
		for scanner.Scan() {
			v, err = strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				return
			}

			digest.Add(v)
		}
		err = scanner.Err()
		if err != nil {
			return
		}

		means := []float64{}
		counts := []float64{}
		digest.ForEachCentroid(func(mean float64, count uint64) bool {
			means = append(means, mean)
			counts = append(counts, float64(count))

			return true
		})

		dividers := histogram.Fence(means)
		hist := stat.Histogram(nil, dividers, means, counts)

		for i, v := range dividers[:len(dividers)-1] {
			fmt.Printf("%g %g\n", v, hist[i])
		}

		return
	},
}
