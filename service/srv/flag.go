package srv

import (
	"flag"
	"fmt"
	"os"
)

var (
	Version     string
	VersionFlag = flag.Bool("v", false, "show version")

	HelpFlag = flag.Bool("h", false, "show help")
)

func FlagInit() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *VersionFlag {
		fmt.Println(Version)
		os.Exit(0)
	}

	if *HelpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}
}
