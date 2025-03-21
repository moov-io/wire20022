package AccountReportingRequest_camt_060_001_05

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteJsonFile(t *testing.T) {
	doc := Document{}

	jsonData, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		fmt.Errorf("error marshal: %w", err)
	}
	jsonFileName := filepath.Join("generated", "test_write.json")
	err = os.WriteFile(jsonFileName, jsonData, 0644)
	if err != nil {
		fmt.Errorf("error writing JSON file: %w", err)
	}

	xmlData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		fmt.Errorf("error marshal: %w", err)
	}
	xmlFileName := filepath.Join("generated", "test_write.xml")
	err = os.WriteFile(xmlFileName, xmlData, 0644)
	if err != nil {
		fmt.Errorf("error writing JSON file: %w", err)
	}
}
