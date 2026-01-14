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
pentestkit --help
```

### TinyScanner (Port Scanning)

Scan specific ports on a target host:

```bash
pentestkit -t 192.168.1.1 -p 22,80,443 -o result1.txt
```

**Options:**
- `-t`: Target host IP address
- `-p`: Comma-separated list of ports to scan
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
pentestkit -d http://example.com -w /path/to/wordlist.txt -o result2.txt
```

**Options:**
- `-d`: Target URL
- `-w`: Path to wordlist file
- `-o`: Output file (optional)

**Example Output:**
```
http://example.com/admin - Status: 200
http://example.com/login - Status: 200
http://example.com/api - Status: 404
```

### HostMapper (Network Mapping)

Discover live hosts on a subnet:

```bash
pentestkit -h 192.168.1.0/24 -o result3.txt
```

**Options:**
- `-h`: Subnet in CIDR notation
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
pentestkit -g http://example.com -o result4.txt
```

**Options:**
- `-g`: Target URL
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
# Basic scan
pentestkit 192.168.1.1

# SYN scan
sudo pentestkit -sS 192.168.1.1 1-1000

# FIN scan
sudo pentestkit -sF 192.168.1.1

# Aggressive scan
pentestkit -A 192.168.1.1

# OS detection
pentestkit -O 192.168.1.1
```

## Scan Types Explained

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
