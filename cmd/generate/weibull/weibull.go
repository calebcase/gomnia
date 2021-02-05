package weibull

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	K      float64 = 1
	Lambda float64 = 1

	Cmd = &cobra.Command{
		Use:   "weibull",
		Short: "generate samples from the Weibull distribution",
		Long: `Generate samples from the Weibull distribution.

https://en.wikipedia.org/wiki/Weibull_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Weibull{
				K:      K,
				Lambda: Lambda,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&K, "k", "k", K, "shape")
	flags.Float64VarP(&Lambda, "lambda", "l", Lambda, "scale")
}
