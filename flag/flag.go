package flag

import (
	"os"

	"github.com/theoriz0/go-public/version"

	"github.com/spf13/pflag"
)

var (
	configFile  = pflag.StringP("config", "c", "./configs/dal.yaml", "Set config file (default ./configs/dal.yaml)")
	help        = pflag.BoolP("help", "h", false, "Show this help message")
	versionFlag = pflag.BoolP("version", "v", false, "Show version info")
)

func HandleFlags() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		os.Exit(0)
	}

	if *versionFlag {
		version.Get().PrintInfo()
		os.Exit(0)
	}
}

func GetConfigFile() string {
	return *configFile
}
