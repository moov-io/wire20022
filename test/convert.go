package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/moov-io/wire20022/pkg/document"
)

func main() {

	xsdFileName := "Fedwire_Funds_Service_Release_2025_AccountReportingRequest_camt_060_001_05_20241122_1718_iso15"
	xsdPath := filepath.Join("xsd", xsdFileName+".xsd")
	var err error
	xsdContent, err := ioutil.ReadFile(xsdPath)
	if err != nil {
		fmt.Errorf("error reading XSD file: %w", err)
	}
	doc, err := document.ParseIso20022Document(xsdContent)
	if err != nil {
		fmt.Errorf("error parsing XSD file: %w", err)
	}
	jsonData, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		fmt.Errorf("error marshal: %w", err)
	}
	jsonFileName := filepath.Join("json", xsdFileName+".json")
	err = os.WriteFile(jsonFileName, jsonData, 0644)
	if err != nil {
		fmt.Errorf("error writing JSON file: %w", err)
	}
}
