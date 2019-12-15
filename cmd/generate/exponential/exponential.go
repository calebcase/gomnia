package exponential

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	Rate float64 = 1

	Cmd = &cobra.Command{
		Use:   "exponential",
		Short: "generate samples from the exponential distribution",
		Long: `Generate samples from the exponential distribution.

.    rate=0.5     rate=1       rate=2
0.0  ██████████▏  ██████████▏  ██████████▏
0.5  ███████▊     ██████▏      ███▊
1.0  ██████       ███▊         █▍
1.5  ████▊        ██▎          ▌
2.0  ███▊         █▍           ▎
2.5  ██▉          ▉            ▏
3.0  ██▎          ▋            ▏
3.5  █▊           ▍            ▏
4.0  █▌           ▎            ▏
4.5  █▏           ▏            ▏

https://en.wikipedia.org/wiki/Exponential_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.Exponential{
				Rate: Rate,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64VarP(&Rate, "rate", "r", Rate, "rate or inverse scale")
}
