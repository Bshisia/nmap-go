# pentest-kit

A comprehensive penetration testing toolkit written in Go for security assessments and network reconnaissance.

## Features

- **TinyScanner**: Fast port scanning tool
- **DirFinder**: Directory brute-forcing for web applications
- **HostMapper**: Network mapping and host discovery
- **HeaderGrabber**: HTTP header analysis and security assessment
- **Multiple Scan Types**: SYN, FIN, XMAS, NULL, UDP, Aggressive, Service Detection, OS Detection
- **Nmap-Compatible Output**: Familiar output format
- **File Output**: Save results to files for reporting

## Installation

```bash
git clone https://learn.zone01kisumu.ke/git/bshisia/pentest-kit.git
cd pentest-kit
go build -o pentestkit
```

## Usage

### Help Command

```bash
go run . --help
# or
pentestkit --help
```

### TinyScanner (Port Scanning)

Scan specific ports on a target host:

```bash
# Basic port scan
go run . -t 192.168.1.1 -p 22,80,443

# Port scan with output file
go run . -t 192.168.1.1 -p 22,80,443 -o result1.txt

# Scan multiple ports
go run . -t 127.0.0.1 -p 21,22,23,25,53,80,110,143,443,993,995
```

**Options:**
- `-t`: Target host IP address (required)
- `-p`: Comma-separated list of ports to scan (required)
- `-o`: Output file (optional)

**Example Output:**
```
Port 22: OPEN
Port 80: OPEN
Port 443: CLOSED
```

### DirFinder (Directory Brute-forcing)

Discover hidden directories on web servers:

```bash
# Basic directory brute-force
go run . -d https://example.com -w wordlist.txt

# Directory brute-force with output file
go run . -d https://example.com -w wordlist.txt -o result2.txt

# Scan specific website
go run . -d https://learn.zone01kisumu.ke -w wordlist.txt -o result2.txt
```

**Options:**
- `-d`: Target URL (must include http:// or https://) (required)
- `-w`: Path to wordlist file (required)
- `-o`: Output file (optional)

**Example Output:**
```
https://example.com/admin - Status: 200
https://example.com/login - Status: 404
https://example.com/robots.txt - Status: 200
```

### HostMapper (Network Mapping)

Discover live hosts on a subnet:

```bash
# Basic network mapping
go run . -h 192.168.1.0/24

# Network mapping with output file
go run . -h 192.168.1.0/24 -o result3.txt

# Scan different subnet
go run . -h 10.0.0.0/24 -o network_scan.txt
```

**Options:**
- `-h`: Subnet in CIDR notation (e.g., 192.168.1.0/24) (required)
- `-o`: Output file (optional)

**Example Output:**
```
192.168.1.1 - ALIVE
192.168.1.100 - ALIVE
192.168.1.222 - ALIVE
```

### HeaderGrabber (HTTP Header Analysis)

Analyze HTTP headers and security configurations:

```bash
# Basic header analysis
go run . -g https://example.com

# Header analysis with output file
go run . -g https://example.com -o result4.txt

# Analyze specific website
go run . -g https://www.passafrika.xyz -o headers.txt
```

**Options:**
- `-g`: Target URL (must include http:// or https://) (required)
- `-o`: Output file (optional)

**Example Output:**
```
Status Code: 200 OK

Headers:
Server: nginx
Content-Type: text/html

Security Analysis:
[+] X-Frame-Options: DENY (Protects against clickjacking)
[-] X-Content-Type-Options: MISSING (Prevents MIME type sniffing)
```

### Legacy Nmap-style Scanning

The tool also supports nmap-style command syntax:

```bash
# Basic TCP connect scan (default)
go run . 192.168.1.1
go run . localhost

# SYN scan (requires root)
sudo go run . -sS 192.168.1.1
sudo go run . -sS 192.168.1.1 1-1000

# FIN scan (requires root)
sudo go run . -sF 192.168.1.1
sudo go run . -sF localhost

# XMAS scan (requires root)
sudo go run . -sX 192.168.1.1

# NULL scan (requires root)
sudo go run . -sN 192.168.1.1

# UDP scan
go run . -sU 192.168.1.1
go run . -sU 192.168.1.1 53,67,68,123

# Aggressive scan (service + OS detection)
go run . -A 192.168.1.1
go run . -A localhost

# Service version detection
go run . -sV 192.168.1.1
go run . -sV 192.168.1.1 80-443

# OS detection
go run . -O 192.168.1.1
go run . -O localhost

# Verbose output
go run . -v -sS 192.168.1.1

# Custom port ranges
go run . -sS 192.168.1.1 1-65535    # All ports
go run . -sS 192.168.1.1 80-443     # Port range
go run . -sS 192.168.1.1 -          # All ports (alternative)
```

## Quick Reference

### All Available Commands

```bash
# Help
go run . --help

# TinyScanner (Port Scanning)
go run . -t <target> -p <ports> [-o output.txt]
go run . -t 192.168.1.1 -p 22,80,443 -o result1.txt

# DirFinder (Directory Brute-forcing)
go run . -d <url> -w <wordlist> [-o output.txt]
go run . -d https://example.com -w wordlist.txt -o result2.txt

# HostMapper (Network Mapping)
go run . -h <subnet> [-o output.txt]
go run . -h 192.168.1.0/24 -o result3.txt

# HeaderGrabber (HTTP Header Analysis)
go run . -g <url> [-o output.txt]
go run . -g https://example.com -o result4.txt

# Legacy Nmap-style Scanning
go run . [-sS|-sF|-sX|-sN|-sU|-A|-sV|-O] <host> [port-range]
go run . -sS 192.168.1.1 1-1000
go run . -A localhost
go run . 192.168.1.1  # Basic scan
```

### Common Usage Patterns

```bash
# Quick port scan
go run . -t 192.168.1.1 -p 22,80,443

# Full network discovery
go run . -h 192.168.1.0/24 -o network.txt

# Website security assessment
go run . -g https://example.com -o headers.txt
go run . -d https://example.com -w wordlist.txt -o dirs.txt

# Comprehensive host scan
go run . -A 192.168.1.100 -o full_scan.txt

# Stealth scanning (requires root)
sudo go run . -sF 192.168.1.1
sudo go run . -sX 192.168.1.1
```

### TCP Connect Scan (Default)
- Completes full TCP handshake
- Most reliable but easily detected
- No root privileges required

### SYN Scan (-sS)
- Half-open scan (doesn't complete handshake)
- Stealthy and fast
- Requires root privileges

### FIN Scan (-sF)
- Sends FIN packets to ports
- Bypasses some firewalls
- Closed ports respond with RST

### XMAS Scan (-sX)
- Sends packets with FIN+PSH+URG flags
- Useful for firewall evasion

### NULL Scan (-sN)
- Sends packets with no flags set
- Ultra-stealthy technique

### UDP Scan (-sU)
- Scans UDP ports
- Important for DNS, DHCP, SNMP services

### Aggressive Scan (-A)
- Combines service detection and OS fingerprinting
- Most comprehensive scan

### OS Detection (-O)
- TCP fingerprinting based on TTL and window size
- Identifies target operating system

## Requirements

- Go 1.16 or higher
- Root/sudo privileges for raw socket scans (SYN, FIN, XMAS, NULL)
- Network access to target hosts
- Wordlist file for directory brute-forcing (sample included)

## Configuration Files

- `wordlist.txt`: Sample wordlist for directory brute-forcing
- Customize with your own wordlists for better results

## Output Files

All tools support the `-o` flag to save results:
- Results are saved in plain text format
- Easy to parse and include in reports
- Timestamped for audit trails

## Ethical and Legal Use

⚠️ **IMPORTANT - READ CAREFULLY**

This tool is designed for **authorized security testing only**. You must:

1. **Obtain Written Permission**: Always get explicit authorization before scanning any network or system
2. **Respect Scope**: Only test systems within the agreed scope
3. **Follow Laws**: Unauthorized port scanning and penetration testing may be illegal in your jurisdiction
4. **Professional Use**: Use this tool responsibly as part of legitimate security assessments
5. **No Malicious Intent**: Never use these tools for unauthorized access or malicious purposes

**Legal Consequences**: Unauthorized use of penetration testing tools can result in:
- Criminal prosecution
- Civil lawsuits
- Fines and penalties
- Imprisonment

**Best Practices**:
- Always work within a defined scope of work
- Document all testing activities
- Report findings responsibly
- Respect privacy and data protection laws

## Limitations

- Stealth scans require root privileges
- Raw socket implementation may vary by OS
- Some firewalls may block or detect scans
- UDP scans are slower due to protocol nature
- Directory brute-forcing depends on wordlist quality
- Network mapping may be slow on large subnets

## Troubleshooting

**Permission Denied Errors**:
```bash
sudo pentestkit -sS 192.168.1.1
```

**Timeout Issues**:
- Increase timeout values in source code
- Check network connectivity
- Verify firewall rules

**Wordlist Not Found**:
```bash
pentestkit -d http://example.com -w ./wordlist.txt -o results.txt
```

## Contributing

Contributions are welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Submit a pull request
4. Follow Go coding standards

## License

MIT License - See LICENSE file for details

## Disclaimer

THIS SOFTWARE IS PROVIDED "AS IS" WITHOUT WARRANTY OF ANY KIND. THE AUTHORS ARE NOT RESPONSIBLE FOR ANY MISUSE OR DAMAGE CAUSED BY THIS TOOL. USE AT YOUR OWN RISK.

## Author

Developed for educational purposes and authorized penetration testing engagements.

## Acknowledgments

- Inspired by nmap and other industry-standard security tools
- Built with Go for performance and portability
- Community contributions welcome
