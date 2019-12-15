package Beta

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Alpha float64 = 1
	Beta  float64 = 1

	Cmd = &cobra.Command{
		Use:   "beta",
		Short: "generate samples from the beta distribution",
		Long: `Generate samples from the beta distribution.

.    alpha=0.5    alpha=5      alpha=2      alpha=2
.    beta=0.5     beta=1       beta=2       beta=4
0.0  ██████████   ▏            █▉           ███▉
0.1  ████▍        ▏            █████▎       ████████▋
0.2  ███▋         ▏            ███████▌     ██████████▏
0.3  ███▎         ▍            █████████    █████████▍
0.4  ███▏         ▉            ██████████   ███████▌
0.5  ███▎         █▌           ██████████▏  █████▎
0.6  ███▎         ██▋          █████████▏   ███▏
0.7  ███▋         ████▍        ███████▌     █▌
0.8  ████▍        ██████▊      █████▎       ▌
0.9  ██████████▏  ██████████▏  ██           ▏

https://en.wikipedia.org/wiki/Beta_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Beta{
				Alpha: Alpha,
				Beta:  Beta,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Alpha, "alpha", "a", Alpha, "alpha")
	flags.Float64VarP(&Beta, "beta", "b", Beta, "beta")
}
