package gumbelright

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Mu   float64 = 1
	Beta float64 = 1

	Cmd = &cobra.Command{
		Use:   "gumbel-right",
		Short: "generate samples from the right-skewed Gumbel distribution",
		Long: `Generate samples from the right-skewed Gumbel distribution.

.     mu=1         mu=1.5       mu=3
.     beta=2       beta=3       beta=4
-5.0  ▏            ▊            ▉
-2.5  ████▌        █████▊       ████▋
0.0   ██████████▏  ██████████▏  █████████
2.5   █████▉       ████████▎    ██████████▏
5.0   ██           ████▉        ████████
7.5   ▊            ██▍          █████▍
10.0  ▎            █▏           ███▎
12.5  ▏            ▌            █▉
15.0  ▏            ▎            █▏
17.5  ▏            ▏            ▋

https://en.wikipedia.org/wiki/Gumbel_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.GumbelRight{
				Mu:   Mu,
				Beta: Beta,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Mu, "mu", "m", Mu, "location")
	flags.Float64VarP(&Beta, "beta", "b", Beta, "scale")
}
