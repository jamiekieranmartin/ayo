package main

import (
	"flag"
	"fmt"

	"github.com/jamiekieranmartin/ayo"
)

const cliVersion = "0.0.1"

const helpMessage = `
ayo is a minimal input/output proxy for common interfaces.
	ayo v%s

Configurable via TOML.

# Inputs

[[input.tcp]]
port="5001"

[[input.udp]]
port="5002"

[[input.http]]
port="80"

# Outputs

[[output.tcp]]
hostname="1.1.1.1"
port="6001"

[[output.udp]]
hostname="1.1.1.1"
port="6002"

[[output.http]]
uri="https://example.com"

`

func main() {
	flag.Usage = func() {
		fmt.Printf(helpMessage, cliVersion)
		flag.PrintDefaults()
	}

	// cli arguments
	path := flag.String("config", "./config.toml", "TOML configuration file")

	version := flag.Bool("version", false, "Print version string and exit")
	help := flag.Bool("help", false, "Print help message and exit")

	flag.Parse()

	// if asked for version, disregard everything else
	if *version {
		fmt.Printf("ayo v%s\n", cliVersion)
		return
	} else if *help {
		flag.Usage()
		return
	}

	// generate config
	config, err := ayo.ReadConfigFile(*path)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	// parse input given expression
	io, err := ayo.New(*config)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	err = io.ListenAndServe()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
}
