package pareto

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Xm    float64 = 1
	Alpha float64 = 1

	Cmd = &cobra.Command{
		Use:   "pareto",
		Short: "generate samples from the Pareto distribution",
		Long: `Generate samples from the Pareto distribution.

.    xm=1         xm=1         xm=1
.    alpha=5      alpha=3      alpha=1
0.0  ▏            ▏            ▏
0.5  ▏            ▏            ▏
1.0  ██████████▏  ██████████▏  ██████████▏
1.5  █▎           ██▌          █████
2.0  ▎            ▉            ███▏
2.5  ▏            ▌            ██
3.0  ▏            ▎            █▌
3.5  ▏            ▏            █▏
4.0  ▏            ▏            ▉
4.5  ▏            ▏            ▊

https://en.wikipedia.org/wiki/Pareto_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Pareto{
				Xm:    Xm,
				Alpha: Alpha,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Xm, "xm", "x", Xm, "scale")
	flags.Float64VarP(&Alpha, "alpha", "a", Alpha, "shape")
}
