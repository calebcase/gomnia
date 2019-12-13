package binomial

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	N float64 = 1
	P float64 = 0.5
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&N, "n", "n", N, "number of trials")
	flags.Float64VarP(&P, "p", "p", P, "success probability for each trial")
}

var Cmd = &cobra.Command{
	Use:   "binomial",
	Short: "generate samples from the binomial distribution",
	Long: `Generate samples from the binomial distribution.

https://en.wikipedia.org/wiki/Binomial_distribution`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		return generate.Sample(distuv.Binomial{
			N: N,
			P: P,
		})
	},
}
