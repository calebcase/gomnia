package Beta

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Alpha float64 = 1
	Beta  float64 = 1
)

func init() {
	generate.Cmd.AddCommand(BetaCmd)

	flags := BetaCmd.Flags()
	flags.Float64VarP(&Alpha, "Alpha", "a", Alpha, "alpha")
	flags.Float64VarP(&Beta, "Beta", "b", Beta, "beta")
}

var BetaCmd = &cobra.Command{
	Use:   "beta",
	Short: "generate samples from the beta distribution",
	Long: `Generate samples from the beta distribution.

https://en.wikipedia.org/wiki/Beta_distribution`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		return generate.Sample(distuv.Beta{
			Alpha: Alpha,
			Beta:  Beta,
		})
	},
}
