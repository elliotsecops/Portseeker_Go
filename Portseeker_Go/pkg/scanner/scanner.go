// pkg/scanner/scanner.go

package scanner

import (
	"Portseeker_Go/pkg/cve" // Correct import path
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// RunNmap runs an Nmap scan on the target host and ports, returning the results as a string
func RunNmap(hostname string, startPort, endPort int) (string, error) {
	// Construct Nmap command
	ports := fmt.Sprintf("%d-%d", startPort, endPort)
	cmd := exec.Command("nmap", "-p", ports, "-sV", hostname)

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run Nmap: %v", err)
		return "", err
	}
	return string(output), nil
}

// ParseNmapOutput parses the Nmap scan results to extract open ports and services
func ParseNmapOutput(nmapOutput string) []PortService {
	var results []PortService
	lines := strings.Split(nmapOutput, "\n")

	// Regular expression to match the Nmap output format
	re := regexp.MustCompile(`(\d+)/tcp\s+open\s+(\S+)\s+(.*)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 4 {
			port := match[1]
			service := match[2]
			version := match[3]
			results = append(results, PortService{Port: port, Service: service, Version: version})
		}
	}
	return results
}

// PortService represents a mapping between a port and its service
type PortService struct {
	Port    string
	Service string
	Version string
}

// ScanWithNmap scans the hostname using Nmap and matches open ports with CVEs
func ScanWithNmap(hostname string, startPort, endPort int, cveList []cve.CVEItem) (Summary, error) {
	// Run Nmap scan
	nmapOutput, err := RunNmap(hostname, startPort, endPort)
	if err != nil {
		fmt.Println("Error running Nmap:", err)
		return Summary{}, err
	}

	// Parse Nmap output to extract open ports and services
	openPorts := ParseNmapOutput(nmapOutput)

	// Initialize summary
	summary := Summary{
		TotalOpenPorts:  len(openPorts),
		Vulnerabilities: make(map[string][]Vulnerability),
	}

	// Match each open port/service with CVEs
	for _, ps := range openPorts {
		fmt.Printf("Port %s is open, service: %s, version: %s\n", ps.Port, ps.Service, ps.Version)
		vulns := MatchCVEs(ps.Service, ps.Version, cveList)
		summary.Vulnerabilities[ps.Service] = append(summary.Vulnerabilities[ps.Service], vulns...)
	}

	// Calculate total vulnerabilities and severity breakdown
	for _, vulns := range summary.Vulnerabilities {
		summary.TotalVulnerabilities += len(vulns)
		for _, vuln := range vulns {
			switch vuln.Severity {
			case "High":
				summary.SeverityBreakdown.High++
			case "Medium":
				summary.SeverityBreakdown.Medium++
			case "Low":
				summary.SeverityBreakdown.Low++
			}
		}
	}

	return summary, nil
}

// MatchCVEs matches the service and version with relevant CVEs
func MatchCVEs(service, version string, cveList []cve.CVEItem) []Vulnerability {
	var vulnerabilities []Vulnerability
	for _, cveItem := range cveList {
		// Check if the CVE description contains the service name
		if strings.Contains(cveItem.CVE.Description.DescriptionData[0].Value, service) {
			// Check if the CVE description contains the version
			if version != "" && strings.Contains(cveItem.CVE.Description.DescriptionData[0].Value, version) {
				vulnerabilities = append(vulnerabilities, Vulnerability{
					CVEID:    cveItem.CVE.CVEDataMeta.ID,
					Service:  service,
					Version:  version,
					Severity: "High", // Placeholder for severity
				})
			} else {
				// If version is not specified, just match the service name
				vulnerabilities = append(vulnerabilities, Vulnerability{
					CVEID:    cveItem.CVE.CVEDataMeta.ID,
					Service:  service,
					Version:  version,
					Severity: "Medium", // Placeholder for severity
				})
			}
		}
	}
	return vulnerabilities
}

// Summary represents the summary of the scan results
type Summary struct {
	TotalOpenPorts       int
	TotalVulnerabilities int
	SeverityBreakdown    struct {
		High   int
		Medium int
		Low    int
	}
	Vulnerabilities map[string][]Vulnerability
}

// Vulnerability represents a vulnerability found during the scan
type Vulnerability struct {
	CVEID    string
	Service  string
	Version  string
	Severity string
}
