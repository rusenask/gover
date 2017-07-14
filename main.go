package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Masterminds/semver"
)

var usageStr = `
Usage: gover [options]
Gover options:
    --current                         Current version (ie: 0.1.1)

    --increase-major                  Increase major version
    --increase-minor                  Increase minor version
    --increase-patch                  Increase patch version

Usage example:
	$ gover --current 1.2.3 --increase-minor
	1.3.0

Common Options:
    -h, --help                       Show this message    
`

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func main() {

	major := flag.Bool("increase-major", false, "increase major version")
	minor := flag.Bool("increase-minor", false, "increase minor version")
	patch := flag.Bool("increase-patch", false, "increase patch version")

	current := flag.String("current", "", "current version")

	flag.Usage = usage

	flag.Parse()

	if *current == "" {
		fmt.Println("current version not specified")
		os.Exit(1)
		return
	}

	version, err := semver.NewVersion(*current)
	if err != nil {
		fmt.Println("failed to parse version: %s", err)
		os.Exit(1)
		return
	}

	updated := *version

	if *major {
		updated = version.IncMajor()
	}

	if *minor {
		updated = updated.IncMinor()
	}

	if *patch {
		updated = updated.IncPatch()
	}

	fmt.Printf("%s", updated.String())
}
