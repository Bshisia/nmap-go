package main

import (
	"flag"
	"fmt"
	"os"
	"pentest-kit/scanner"
	"pentest-kit/tools"
	"pentest-kit/utils"
)

func main() {
	// Define flags for new tools
	target := flag.String("t", "", "Target host for port scanning")
	portsFlag := flag.String("p", "", "Ports to scan (comma-separated, e.g., 22,80,443)")
	dirURL := flag.String("d", "", "URL for directory brute-forcing")
	wordlist := flag.String("w", "", "Wordlist file for directory brute-forcing")
	hostSubnet := flag.String("h", "", "Subnet for network mapping (e.g., 192.168.1.0/24)")
	grabURL := flag.String("g", "", "URL for HTTP header analysis")
	outputFile := flag.String("o", "", "Output file for results")
	help := flag.Bool("help", false, "Show help message")
	
	flag.Parse()
	
	// Show help
	if *help || len(os.Args) == 1 {
		showHelp()
		return
	}
	
	// TinyScanner: Port scanning
	if *target != "" && *portsFlag != "" {
		tools.TinyScanner(*target, *portsFlag, *outputFile)
		return
	}
	
	// DirFinder: Directory brute-forcing
	if *dirURL != "" && *wordlist != "" {
		tools.DirFinder(*dirURL, *wordlist, *outputFile)
		return
	}
	
	// HostMapper: Network mapping
	if *hostSubnet != "" {
		tools.HostMapper(*hostSubnet, *outputFile)
		return
	}
	
	// HeaderGrabber: HTTP header analysis
	if *grabURL != "" {
		tools.HeaderGrabber(*grabURL, *outputFile)
		return
	}
	
	// Legacy nmap-style scanning
	config := ParseArgs()
	ports := utils.ParsePortRange(config.PortRange)

	if config.Verbose {
		fmt.Printf("Starting scan with verbose output...\n")
	}
	if config.Timing != "" {
		fmt.Printf("Using timing template: %s\n", config.Timing)
	}

	switch {
	case config.Aggressive:
		scanner.AggressiveScan(config.Host, ports)
	case config.SynScan:
		scanner.SynScan(config.Host, ports)
	case config.UdpScan:
		scanner.UdpScan(config.Host, ports)
	case config.FinScan:
		scanner.FinScan(config.Host, ports)
	case config.XmasScan:
		scanner.XmasScan(config.Host, ports)
	case config.NullScan:
		scanner.NullScan(config.Host, ports)
	case config.OSDetection:
		scanner.OSDetection(config.Host, ports)
	default:
		scanner.ScanPorts(config.Host, ports, config.ServiceDetection)
	}
}

func showHelp() {
	fmt.Println("PentestKit - Penetration Testing Toolkit")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  pentestkit [options]")
	fmt.Println("")
	fmt.Println("Tools:")
	fmt.Println("  TinyScanner (Port Scanning):")
	fmt.Println("    pentestkit -t <target> -p <ports> -o <output>")
	fmt.Println("    Example: pentestkit -t 192.168.1.1 -p 22,80,443 -o result1.txt")
	fmt.Println("")
	fmt.Println("  DirFinder (Directory Brute-forcing):")
	fmt.Println("    pentestkit -d <url> -w <wordlist> -o <output>")
	fmt.Println("    Example: pentestkit -d http://example.com -w /path/to/wordlist.txt -o result2.txt")
	fmt.Println("")
	fmt.Println("  HostMapper (Network Mapping):")
	fmt.Println("    pentestkit -h <subnet> -o <output>")
	fmt.Println("    Example: pentestkit -h 192.168.1.0/24 -o result3.txt")
	fmt.Println("")
	fmt.Println("  HeaderGrabber (HTTP Header Analysis):")
	fmt.Println("    pentestkit -g <url> -o <output>")
	fmt.Println("    Example: pentestkit -g http://example.com -o result4.txt")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -t string    Target host for port scanning")
	fmt.Println("  -p string    Ports to scan (comma-separated)")
	fmt.Println("  -d string    URL for directory brute-forcing")
	fmt.Println("  -w string    Wordlist file path")
	fmt.Println("  -h string    Subnet for network mapping")
	fmt.Println("  -g string    URL for HTTP header analysis")
	fmt.Println("  -o string    Output file for results")
	fmt.Println("  --help       Show this help message")
	fmt.Println("")
	fmt.Println("Legacy Nmap-style Scanning:")
	fmt.Println("  pentestkit [-sS|-sF|-sX|-sN|-sU|-A|-sV|-O] <host> [port-range]")
	fmt.Println("  Example: pentestkit -sS 192.168.1.1 1-1000")
	fmt.Println("")
	fmt.Println("Legal Notice:")
	fmt.Println("  This tool is for authorized testing only. Unauthorized use may be illegal.")
	fmt.Println("  Always obtain proper authorization before scanning networks you don't own.")
}


