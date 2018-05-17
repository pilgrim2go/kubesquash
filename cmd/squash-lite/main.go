package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/solo-io/kubesquash/pkg/cmd"
)

const descriptionUsage = `Normally squash lite requires no arguments. just run it!
it works by creating additional privileged debug pod and then attaching to it. 
Kubernetes with CRI is needed. Due to a technical limitation, squash-lite doesn't support 
scratch images at the moment (squash lite relys on the 'ls' command present in the image). 
`

func main() {
	var cfg cmd.SquashConfig
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", descriptionUsage)
		flag.PrintDefaults()
	}

	flag.BoolVar(&cfg.NoClean, "no-clean", false, "don't clean temporar pod when existing")
	flag.BoolVar(&cfg.ChooseDebugger, "no-guess-debugger", false, "don't auto detect debugger to use")
	flag.BoolVar(&cfg.ChoosePod, "no-guess-pod", false, "don't auto detect pod to use")
	flag.BoolVar(&cfg.NoDetectSkaffold, "no-detect-pod", false, "don't auto settigns based on skaffold configuration present in current folder")
	flag.BoolVar(&cfg.DebugServer, "debug-server", false, "start a debug server instead of an interactive session")
	flag.IntVar(&cfg.TimeoutSeconds, "timeout", 300, "timeout in seconds to wait for debug pod to be ready")
	flag.StringVar(&cfg.DebugContainerVersion, "container-version", cmd.ImageVersion, "debug container version to use")
	flag.StringVar(&cfg.DebugContainerRepo, "container-repo", cmd.ImageRepo, "debug container repo to use")
	flag.Parse()

	err := cmd.StartDebugContainer(cfg)
	if err != nil {
		fmt.Println(err)
	}
}
