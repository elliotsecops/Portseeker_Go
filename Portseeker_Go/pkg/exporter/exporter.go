// pkg/exporter/exporter.go

package exporter

import (
	"Portseeker_Go/pkg/scanner"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf" // For PDF export
)

func ExportToJSON(summary scanner.Summary, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(summary)
}

func ExportToCSV(summary scanner.Summary, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Service", "CVE ID", "Severity"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write CSV rows
	for service, vulns := range summary.Vulnerabilities {
		for _, vuln := range vulns {
			row := []string{service, vuln.CVEID, vuln.Severity}
			if err := writer.Write(row); err != nil {
				return err
			}
		}
	}

	return nil
}

func ExportToPDF(summary scanner.Summary, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Portseeker Scan Results")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Total Open Ports: "+fmt.Sprintf("%d", summary.TotalOpenPorts))
	pdf.Ln(8)
	pdf.Cell(40, 10, "Total Vulnerabilities: "+fmt.Sprintf("%d", summary.TotalVulnerabilities))
	pdf.Ln(8)
	pdf.Cell(40, 10, "Severity Breakdown: High: "+fmt.Sprintf("%d", summary.SeverityBreakdown.High)+", Medium: "+fmt.Sprintf("%d", summary.SeverityBreakdown.Medium)+", Low: "+fmt.Sprintf("%d", summary.SeverityBreakdown.Low))
	pdf.Ln(12)

	for service, vulns := range summary.Vulnerabilities {
		pdf.Cell(40, 10, "Vulnerabilities for "+service+":")
		pdf.Ln(8)
		for _, vuln := range vulns {
			pdf.Cell(40, 10, "CVE ID: "+vuln.CVEID+", Severity: "+vuln.Severity)
			pdf.Ln(6)
		}
		pdf.Ln(8)
	}

	return pdf.OutputFileAndClose(filename)
}
