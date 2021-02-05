package inversegamma

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Alpha float64 = 1
	Beta  float64 = 1

	Cmd = &cobra.Command{
		Use:   "inverse-gamma",
		Short: "generate samples from the inverse gamma distribution",
		Long: `Generate samples from the inverse gamma distribution.

.    alpha=1      alpha=2      alpha=3      alpha=3
.    beta=1       beta=1       beta=1       beta=0.5
0.0  ██▍          ████▌        ████████▊    ██████████▏
0.3  ██████████▏  ██████████▏  ██████████▏  ██▍
0.6  █████████▎   █████▋       ███▎         ▌
0.9  ███████      ███          █▎           ▎
1.2  █████▎       █▊           ▋            ▏
1.5  ████         █▏           ▍            ▏
1.8  ███▎         ▊            ▎            ▏
2.1  ██▋          ▌            ▏            ▏
2.4  ██▏          ▍            ▏            ▏
2.7  █▊           ▍            ▏            ▏

https://en.wikipedia.org/wiki/Inverse-gamma_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.InverseGamma{
				Alpha: Alpha,
				Beta:  Beta,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Alpha, "alpha", "a", Alpha, "shape")
	flags.Float64VarP(&Beta, "beta", "b", Beta, "scale")
}
