package chisquared

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	K float64 = 1
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&K, "k", "k", K, "degrees of freedom")
}

var Cmd = &cobra.Command{
	Use:   "chi-squared",
	Short: "generate samples from the chi-squared distribution",
	Long: `Generate samples from the chi-squared distribution.

https://en.wikipedia.org/wiki/Chi-squared_distribution`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		return generate.Sample(distuv.ChiSquared{
			K: K,
		})
	},
}
