package export

import (
	"Portseeker_Go/internal/model"
	"encoding/json"
	"fmt"
	"os"
)

// ExportToJSON exports the data in JSON format.
func ExportToJSON(summary model.SummaryDashboard) {
	file, err := os.Create("vulnerabilities.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(summary)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("Exported to JSON: vulnerabilities.json")
}
