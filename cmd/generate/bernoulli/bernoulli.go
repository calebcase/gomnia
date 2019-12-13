package bernoulli

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	P float64 = 0.5
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&P, "p", "p", P, "probability of taking on a value of 1")
}

var Cmd = &cobra.Command{
	Use:   "bernoulli",
	Short: "generate samples from the bernoulli distribution",
	Long: `Generate samples from the bernoulli distribution.

https://en.wikipedia.org/wiki/Bernoulli_distribution`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		return generate.Sample(distuv.Bernoulli{
			P: P,
		})
	},
}
