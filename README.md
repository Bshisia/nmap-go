# pentest-kit

A lightweight, nmap-like port scanner written in Go for penetration testing and network reconnaissance.

## Features

- **Multiple Scan Types**: SYN, FIN, XMAS, NULL, UDP, Aggressive, Service Detection, OS Detection
- **Nmap-Compatible Output**: Familiar output format matching nmap
- **Fast & Efficient**: Optimized scanning with configurable timeouts
- **Service Detection**: Identifies common services running on open ports
- **OS Fingerprinting**: TCP-based operating system detection
- **Flexible Port Ranges**: Scan specific ports, ranges, or all ports

## Installation

```bash
git clone https://learn.zone01kisumu.ke/git/bshisia/pentest-kit.git
cd pentest-kit
go build
```

## Usage

### Basic Scan (Default TCP Connect)
```bash
go run . <host>
go run . 192.168.1.1
```

### Scan Types

**SYN Scan (Stealth)**
```bash
sudo go run . -sS <host> [port-range]
sudo go run . -sS 192.168.1.1
```

**FIN Scan**
```bash
sudo go run . -sF <host> [port-range]
sudo go run . -sF 192.168.1.1 1-1000
```

**XMAS Scan**
```bash
sudo go run . -sX <host> [port-range]
sudo go run . -sX 192.168.1.1
```

**NULL Scan**
```bash
sudo go run . -sN <host> [port-range]
sudo go run . -sN 192.168.1.1
```

**UDP Scan**
```bash
sudo go run . -sU <host> [port-range]
sudo go run . -sU 192.168.1.1 53
```

**Aggressive Scan (Service + OS Detection)**
```bash
go run . -A <host> [port-range]
go run . -A 192.168.1.1 1-1000
```

**Service Version Detection**
```bash
go run . -sV <host> [port-range]
go run . -sV 192.168.1.1 80-443
```

**OS Detection**
```bash
go run . -O <host> [port-range]
go run . -O 192.168.1.1
```

### Port Range Options

- **Single port**: `80`
- **Port range**: `1-1000`
- **All ports**: `-` (scans 1-65535)
- **Default**: `1-1000` (when no range specified)

### Examples

```bash
# Scan common ports on target
go run . 192.168.1.100

# FIN scan on specific port range
sudo go run . -sF 192.168.1.100 20-80

# Aggressive scan with service detection
go run . -A 192.168.1.100

# Scan all ports with SYN scan
sudo go run . -sS 192.168.1.100 -

# OS detection on default ports
go run . -O 192.168.1.100
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
- Named for "lit up" TCP flags

### NULL Scan (-sN)
- Sends packets with no flags set
- Ultra-stealthy technique
- Relies on RFC compliance

### UDP Scan (-sU)
- Scans UDP ports
- Slower than TCP scans
- Important for DNS, DHCP, SNMP services

### Aggressive Scan (-A)
- Combines service detection and OS fingerprinting
- Most comprehensive scan
- Provides detailed service information

### OS Detection (-O)
- TCP fingerprinting based on TTL and window size
- Analyzes port patterns
- Identifies target operating system

## Output Format

The tool produces nmap-compatible output:

```
Starting Nmap 7.94SVN ( https://nmap.org ) at 2026-01-14 10:13 EAT
Nmap scan report for 192.168.1.222
Host is up (0.00018s latency).
Not shown: 998 closed tcp ports (conn-refused)
PORT   STATE SERVICE
22/tcp open  ssh
80/tcp open  http

Pentest-Kit done: 1 IP address (1 host up) scanned in 0.06 seconds
```

## Requirements

- Go 1.16 or higher
- Root/sudo privileges for raw socket scans (SYN, FIN, XMAS, NULL)
- Network access to target hosts

## Limitations

- Stealth scans require root privileges
- Raw socket implementation may vary by OS
- Some firewalls may block or detect scans
- UDP scans are slower due to protocol nature

## Legal Notice

⚠️ **WARNING**: This tool is for educational and authorized testing purposes only. Unauthorized port scanning may be illegal in your jurisdiction. Always obtain proper authorization before scanning networks you don't own.

## License

MIT License - See LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.

## Author

Developed for penetration testing and network security assessment.
