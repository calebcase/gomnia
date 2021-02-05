package poisson

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Lambda float64 = 1

	Cmd = &cobra.Command{
		Use:   "poisson",
		Short: "generate samples from the Poisson distribution",
		Long: `Generate samples from the Poisson distribution.

.     lambda=1     lambda=4     lambda=10
0.0   ██████████▏  ██████▏      ▏
2.0   ▉            ██████████▏  █▏
4.0   ▏            ██████▊      ████
6.0   ▏            ██▍          ████████▏
8.0   ▏            ▌            ██████████▏
10.0  ▏            ▏            ████████▍
12.0  ▏            ▏            █████
14.0  ▏            ▏            ██▍
16.0  ▏            ▏            ▉
18.0  ▏            ▏            ▎

https://en.wikipedia.org/wiki/Poisson_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Poisson{
				Lambda: Lambda,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Lambda, "lambda", "l", Lambda, "rate")
}
