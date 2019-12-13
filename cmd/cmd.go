package cmd

import (
	"os"

	"github.com/inconshreveable/log15"

	_ "github.com/calebcase/gomnia/cmd/generate/bernoulli"
	_ "github.com/calebcase/gomnia/cmd/generate/beta"
	_ "github.com/calebcase/gomnia/cmd/generate/binomial"
	_ "github.com/calebcase/gomnia/cmd/generate/chisquared"
	_ "github.com/calebcase/gomnia/cmd/generate/exponential"
	_ "github.com/calebcase/gomnia/cmd/generate/f"
	_ "github.com/calebcase/gomnia/cmd/limit"
	_ "github.com/calebcase/gomnia/cmd/plot/histogram"
	"github.com/calebcase/gomnia/cmd/root"
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
