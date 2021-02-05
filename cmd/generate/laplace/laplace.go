package laplace

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Mu    float64
	Scale float64 = 1

	Cmd = &cobra.Command{
		Use:   "laplace",
		Short: "generate samples from the Laplace distribution",
		Long: `Generate samples from the Laplace distribution.

.      mu=0         mu=0         mu=0         mu=-5
.      scale=1      scale=2      scale=4      scale=4
-10.0  ▏            ▎            █▍           ████▎
-8.0   ▏            ▌            ██▎          ██████▉
-6.0   ▎            █▍           ███▊         ██████████▏
-4.0   █▍           ███▊         ██████▏      ██████▉
-2.0   ██████████▏  ██████████   ██████████   ████▍
0.0    ██████████   ██████████▏  ██████████▏  ██▋
2.0    █▍           ███▊         ██████▏      █▋
4.0    ▎            █▍           ███▊         █
6.0    ▏            ▌            ██▎          ▋
8.0    ▏            ▎            █▍           ▍

https://en.wikipedia.org/wiki/Laplace_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Laplace{
				Mu:    Mu,
				Scale: Scale,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Mu, "mu", "m", Mu, "location")
	flags.Float64VarP(&Scale, "scale", "s", Scale, "scale")
}
