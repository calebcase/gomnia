package normal

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Mu    float64
	Sigma float64 = 1

	Cmd = &cobra.Command{
		Use:   "normal",
		Short: "generate samples from the normal distribution",
		Long: `Generate samples from the normal distribution.

.     mu=0          mu=0         mu=0          mu=-2
.     sigma=0.4472  sigma=1      sigma=2.2361  sigma=0.7071
-4.0  ▏             ▏            ██▉           █
-3.2  ▏             ▍            ████▊         █████▊
-2.4  ▏             █▋           ██████▊       ██████████▏
-1.6  ▉             █████▌       ████████▉     █████▋
-0.8  ██████████▏   ██████████   ██████████▏   █
0.0   ██████████    ██████████▏  ██████████    ▏
0.8   ▉             █████▍       ████████▉     ▏
1.6   ▏             █▋           ██████▊       ▏
2.4   ▏             ▎            ████▋         ▏
3.2   ▏             ▏            ██▉           ▏

https://en.wikipedia.org/wiki/Normal_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Normal{
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
