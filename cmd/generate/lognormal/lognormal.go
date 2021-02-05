package lognormal

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Mu    float64
	Sigma float64 = 1

	Cmd = &cobra.Command{
		Use:   "log-normal",
		Short: "generate samples from the log-normal distribution",
		Long: `Generate samples from the log-normal distribution.

.    mu=0         mu=0         mu=0
.    sigma=1      sigma=0.5    sigma=0.25
0.0  ██████▏      ▍            ▏
0.3  ██████████▏  █████▋       ▌
0.6  ████████▏    ██████████▏  ███████▌
0.9  ██████       ████████▋    ██████████▏
1.2  ████▌        █████▋       ████▎
1.5  ███▍         ███▍         █▏
1.8  ██▋          ██           ▎
2.1  ██▏          █▏           ▏
2.4  █▋           ▋            ▏
2.7  █▍           ▍            ▏

https://en.wikipedia.org/wiki/Log-normal_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.LogNormal{
				Mu:    Mu,
				Sigma: Sigma,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Mu, "mu", "m", Mu, "expected value")
	flags.Float64VarP(&Sigma, "sigma", "s", Sigma, "standard deviation")
}
