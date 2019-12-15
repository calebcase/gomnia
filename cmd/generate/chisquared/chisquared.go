package chisquared

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	K float64 = 1

	Cmd = &cobra.Command{
		Use:   "chi-squared",
		Short: "generate samples from the chi-squared distribution",
		Long: `Generate samples from the chi-squared distribution.

.    k=1          k=4          k=9
0.0  ██████████▏  █████▏       ▏
1.0  ██▍          █████████▊   ▉
2.0  █▏           ██████████▏  ██▋
3.0  ▋            ████████▌    █████▏
4.0  ▍            ██████▊      ███████▌
5.0  ▎            █████        █████████▏
6.0  ▏            ███▌         ██████████▏
7.0  ▏            ██▌          ██████████
8.0  ▏            █▊           █████████▍
9.0  ▏            █▎           ████████▌

https://en.wikipedia.org/wiki/Chi-squared_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.ChiSquared{
				K: K,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&K, "k", "k", K, "degrees of freedom")
}
