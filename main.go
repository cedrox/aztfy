package main

import (
	"fmt"
	"log"
	"os"

	"flag"

	"github.com/magodo/aztfy/internal/config"
	"github.com/magodo/aztfy/internal/ui"
)

var (
	flagVersion *bool
)

func init() {
	flagVersion = flag.Bool("v", false, "Print version")
}

const usage = `aztfy [option] <resource group name>
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.NewConfig(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	prog, err := ui.NewProgram(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := prog.Start(); err != nil {
		log.Fatal(err)
	}
}
