package f

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	D1 float64 = 1
	D2 float64 = 1
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64Var(&D1, "d1", D1, "degree of freedom 1")
	flags.Float64Var(&D2, "d2", D2, "degree of freedom 2")
}

var Cmd = &cobra.Command{
	Use:   "f",
	Short: "generate samples from the F distribution",
	Long: `Generate samples from the F distribution.

https://en.wikipedia.org/wiki/F_distribution`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		return generate.Sample(distuv.F{
			D1: D1,
			D2: D2,
		})
	},
}
