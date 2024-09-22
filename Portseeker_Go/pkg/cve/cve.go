// pkg/cve/cve.go
package cve

import (
	"encoding/json"
	"log"
	"os"
)

type CVEItem struct {
	CVE struct {
		CVEDataMeta struct {
			ID string `json:"ID"`
		} `json:"CVE_data_meta"`
		Description struct {
			DescriptionData []struct {
				Value string `json:"value"`
			} `json:"description_data"`
		} `json:"description"`
	} `json:"CVE"`
}

type CVEList struct {
	CVEItems []CVEItem `json:"CVE_Items"`
}

func LoadCVE(filename string) (*CVEList, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open CVE file: %v", err)
		return nil, err
	}
	defer file.Close()

	var cves CVEList
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cves)
	if err != nil {
		log.Fatalf("Failed to parse CVE JSON: %v", err)
		return nil, err
	}
	return &cves, nil
}
