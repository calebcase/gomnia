package f

import (
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/calebcase/gomnia/cmd/generate"
)

var (
	D1 float64 = 1
	D2 float64 = 1

	Cmd = &cobra.Command{
		Use:   "f",
		Short: "generate samples from the F distribution",
		Long: `Generate samples from the F distribution.

.    d1=1         d1=5         d1=100
.    d2=1         d2=2         d2=100
0.0  ██████████▏  ██████████▏  ▏
0.5  ██▉          ████████▉    ██████████▏
1.0  █▊           █████▍       █████████▋
1.5  █▎           ███▋         ▌
2.0  ▉            ██▌          ▏
2.5  ▊            █▉           ▏
3.0  ▋            █▌           ▏
3.5  ▌            █▏           ▏
4.0  ▍            ▉            ▏
4.5  ▍            ▊            ▏

https://en.wikipedia.org/wiki/F_distribution`,
		RunE: func(command *cobra.Command, args []string) (err error) {
			return generate.Sample(distuv.F{
				D1: D1,
				D2: D2,
			})
		},
	}
)

func init() {
	generate.Cmd.AddCommand(Cmd)

	flags := Cmd.Flags()
	flags.Float64Var(&D1, "d1", D1, "degree of freedom 1")
	flags.Float64Var(&D2, "d2", D2, "degree of freedom 2")
}
