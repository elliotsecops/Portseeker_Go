package export

import (
	"Portseeker_Go/internal/model"
	"encoding/csv"
	"fmt"
	"os"
)

// ExportToCSV exports the data in CSV format.
func ExportToCSV(summary model.SummaryDashboard) {
	file, err := os.Create("vulnerabilities.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"CVE ID", "Severity", "Affected Service", "Description", "Recommendation"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	// Write CSV rows
	for _, vulns := range summary.ServiceSummary {
		for _, vuln := range vulns {
			row := []string{vuln.CVEID, vuln.Severity, vuln.AffectedService, vuln.Description, vuln.Recommendation}
			err := writer.Write(row)
			if err != nil {
				fmt.Println("Error writing CSV row:", err)
				return
			}
		}
	}

	fmt.Println("Exported to CSV: vulnerabilities.csv")
}
