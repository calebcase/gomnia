package exponential

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Rate float64 = 1
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Rate, "rate", "r", Rate, "rate or inverse scale")
}

var Cmd = &cobra.Command{
	Use:   "exponential",
	Short: "generate samples from the exponential distribution",
	Long: `Generate samples from the exponential distribution.

https://en.wikipedia.org/wiki/Exponential_distribution`,
	RunE: func(command *cobra.Command, args []string) (err error) {
		return generate.Sample(distuv.Exponential{
			Rate: Rate,
		})
	},
}
