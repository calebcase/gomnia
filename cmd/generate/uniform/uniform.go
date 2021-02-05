package uniform

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Cmd = &cobra.Command{
		Use:   "uniform",
		Short: "generate samples from the uniform distribution",
		Long: `Generate samples from the uniform distribution.

https://en.wikipedia.org/wiki/Uniform_distribution_%28continuous%29`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Uniform{
				Min: generate.ScaledMin,
				Max: generate.ScaledMax,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)
}
