package main

import (
	"fmt"
	"os"
)

type Config struct {
	ServiceDetection bool
	SynScan         bool
	Aggressive      bool
	UdpScan         bool
	FinScan         bool
	XmasScan        bool
	NullScan        bool
	OSDetection     bool
	Verbose         bool
	Host            string
	PortRange       string
	Timing          string
}

func ParseArgs() *Config {
	config := &Config{}

	switch {
	case len(os.Args) == 6 && os.Args[1] == "-v" && (os.Args[4][:2] == "-T"):
		config.Verbose = true
		config.Timing = os.Args[4]
		setScanType(config, os.Args[2])
		config.Host = os.Args[3]
		config.PortRange = os.Args[5]
	case len(os.Args) == 5 && (os.Args[3][:2] == "-T"):
		config.Timing = os.Args[3]
		setScanType(config, os.Args[1])
		config.Host = os.Args[2]
		config.PortRange = os.Args[4]
	case len(os.Args) == 5 && os.Args[1] == "-v":
		config.Verbose = true
		setScanType(config, os.Args[2])
		config.Host = os.Args[3]
		config.PortRange = os.Args[4]
	case len(os.Args) == 4 && isScanFlag(os.Args[1]):
		setScanType(config, os.Args[1])
		config.Host = os.Args[2]
		config.PortRange = os.Args[3]
	case len(os.Args) == 3:
		config.Host = os.Args[1]
		config.PortRange = os.Args[2]
	default:
		showUsage()
	}

	return config
}

func setScanType(config *Config, flag string) {
	switch flag {
	case "-sV":
		config.ServiceDetection = true
	case "-sS":
		config.SynScan = true
	case "-A":
		config.Aggressive = true
	case "-sU":
		config.UdpScan = true
	case "-sF":
		config.FinScan = true
	case "-sX":
		config.XmasScan = true
	case "-sN":
		config.NullScan = true
	case "-O":
		config.OSDetection = true
	}
}

func isScanFlag(flag string) bool {
	return flag == "-sV" || flag == "-sS" || flag == "-A" || flag == "-sU" || 
		   flag == "-sF" || flag == "-sX" || flag == "-sN" || flag == "-O"
}

func showUsage() {
	fmt.Println("Usage: go run main.go [-v] [-sV|-sS|-A|-sU|-sF|-sX|-sN|-O] <host> [-T0-T5] <port-range|->")
	fmt.Println("Example: go run main.go -A 192.168.1.1 80-443")
	fmt.Println("Example: go run main.go -sS 192.168.1.1 - (scan all ports)")
	os.Exit(1)
}