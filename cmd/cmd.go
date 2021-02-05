package cmd

import (
	"os"

	"github.com/inconshreveable/log15"

	"github.com/calebcase/gomnia/cmd/root"

	_ "github.com/calebcase/gomnia/cmd/filter"
	_ "github.com/calebcase/gomnia/cmd/generate/bernoulli"
	_ "github.com/calebcase/gomnia/cmd/generate/beta"
	_ "github.com/calebcase/gomnia/cmd/generate/binomial"
	_ "github.com/calebcase/gomnia/cmd/generate/chisquared"
	_ "github.com/calebcase/gomnia/cmd/generate/exponential"
	_ "github.com/calebcase/gomnia/cmd/generate/f"
	_ "github.com/calebcase/gomnia/cmd/generate/gamma"
	_ "github.com/calebcase/gomnia/cmd/generate/gumbelright"
	_ "github.com/calebcase/gomnia/cmd/generate/inversegamma"
	_ "github.com/calebcase/gomnia/cmd/generate/laplace"
	_ "github.com/calebcase/gomnia/cmd/generate/lognormal"
	_ "github.com/calebcase/gomnia/cmd/generate/normal"
	_ "github.com/calebcase/gomnia/cmd/generate/pareto"
	_ "github.com/calebcase/gomnia/cmd/generate/poisson"
	_ "github.com/calebcase/gomnia/cmd/generate/studentst"
	_ "github.com/calebcase/gomnia/cmd/generate/triangular"
	_ "github.com/calebcase/gomnia/cmd/generate/uniform"
	_ "github.com/calebcase/gomnia/cmd/generate/weibull"
	_ "github.com/calebcase/gomnia/cmd/limit"
	_ "github.com/calebcase/gomnia/cmd/plot/histogram"
	_ "github.com/calebcase/gomnia/cmd/summarize/histogram"
	_ "github.com/calebcase/gomnia/cmd/summarize/histogram/tdigest"
)

func Main() {
	if err := root.Cmd.Execute(); err != nil {
		root.Log.Error("Failed to execute command.", log15.Ctx{
			"err": err,
		})
		os.Exit(1)
	}
}
