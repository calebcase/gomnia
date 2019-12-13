package root

import (
	"os"
	"strconv"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
)

var (
	Log = log15.New()

	Cmd = &cobra.Command{
		Use:   "gomnia",
		Short: "numerical and scientific algorithms",
		Long: `Gomnia is a CLI tool designed to make using numerical and scientific algorithms
from the terminal productive and performant.
`,
	}
)

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))

	lvl := log15.LvlWarn
	lvlStr, lvlProvided := os.LookupEnv("GOMNIA_LOG_LEVEL")
	if lvlProvided {
		lvlParsed, err := log15.LvlFromString(lvlStr)
		if err == nil {
			lvl = lvlParsed
		}
	}

	var verbosity uint = 0
	verbosityStr, verbosityProvided := os.LookupEnv("GOMNIA_LOG_VERBOSITY")
	if verbosityProvided {
		verbosityParsed, err := strconv.ParseUint(verbosityStr, 10, 64)
		if err == nil {
			verbosity = uint(verbosityParsed)
		}
	}

	SetLogger(lvl, verbosity, log15.TerminalFormat())
}

func SetLogger(lvl log15.Lvl, verbosity uint, format log15.Format) {
	sh := log15.StreamHandler(os.Stderr, format)
	fh := log15.LvlFilterHandler(lvl, sh)

	if verbosity >= 1 {
		fh = log15.CallerFileHandler(fh)
	}

	if verbosity >= 2 {
		fh = log15.CallerFuncHandler(fh)
	}

	if verbosity >= 3 {
		fh = log15.CallerStackHandler("%+v", fh)
	}

	Log.SetHandler(fh)
}
