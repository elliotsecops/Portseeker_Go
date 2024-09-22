// cmd/main.go

package main

import (
	"Portseeker_Go/pkg/cve"
	"Portseeker_Go/pkg/exporter"
	"Portseeker_Go/pkg/scanner"
	"flag"
	"fmt"
	"log"
)

func main() {
	// Define command-line flags
	hostname := flag.String("host", "127.0.0.1", "The hostname to scan")
	startPort := flag.Int("start", 1, "Start of port range")
	endPort := flag.Int("end", 1024, "End of port range")
	cveFile := flag.String("cve", "cve_json_files/nvdcve-1.1-2024.json", "Path to the CVE JSON file")
	flag.Parse()

	// Load CVE data
	cveList, err := cve.LoadCVE(*cveFile)
	if err != nil {
		log.Fatalf("Error loading CVE: %v", err)
	}

	// Scan and match CVEs using Nmap
	fmt.Printf("Scanning %s from port %d to %d using Nmap...\n", *hostname, *startPort, *endPort)
	summary, err := scanner.ScanWithNmap(*hostname, *startPort, *endPort, cveList.CVEItems)
	if err != nil {
		log.Fatalf("Error scanning with Nmap: %v", err)
	}

	// Display summary dashboard
	fmt.Println("\nSummary Dashboard")
	fmt.Printf("Total Open Ports: %d\n", summary.TotalOpenPorts)
	fmt.Printf("Total Vulnerabilities: %d\n", summary.TotalVulnerabilities)
	fmt.Printf("Severity Breakdown: High: %d, Medium: %d, Low: %d\n", summary.SeverityBreakdown.High, summary.SeverityBreakdown.Medium, summary.SeverityBreakdown.Low)

	// Display vulnerabilities by service
	for service, vulns := range summary.Vulnerabilities {
		fmt.Printf("\nVulnerabilities for %s:\n", service)
		for _, vuln := range vulns {
			fmt.Printf("CVE ID: %s, Severity: %s\n", vuln.CVEID, vuln.Severity)
		}
	}

	// Ask user if they want to export the results
	var exportChoice string
	fmt.Print("\nDo you want to export the results? (yes/no): ")
	fmt.Scanln(&exportChoice)

	if exportChoice == "yes" {
		var exportFormat string
		fmt.Print("In which format do you want to export the results? (json/csv/pdf): ")
		fmt.Scanln(&exportFormat)

		var exportFile string
		fmt.Print("Enter the filename for export: ")
		fmt.Scanln(&exportFile)

		switch exportFormat {
		case "json":
			err := exporter.ExportToJSON(summary, exportFile+".json")
			if err != nil {
				log.Fatalf("Error exporting to JSON: %v", err)
			}
			fmt.Println("Results exported to JSON successfully.")
		case "csv":
			err := exporter.ExportToCSV(summary, exportFile+".csv")
			if err != nil {
				log.Fatalf("Error exporting to CSV: %v", err)
			}
			fmt.Println("Results exported to CSV successfully.")
		case "pdf":
			err := exporter.ExportToPDF(summary, exportFile+".pdf")
			if err != nil {
				log.Fatalf("Error exporting to PDF: %v", err)
			}
			fmt.Println("Results exported to PDF successfully.")
		default:
			log.Fatalf("Unsupported export format: %s", exportFormat)
		}
	}
}
