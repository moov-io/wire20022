package document

import (
	"os"
	"path/filepath"
	"testing"
)

func TestJsonXmlWithDocumentAccountReportingRequest_camt_060_001_05(t *testing.T) {
	inputXml, err := os.ReadFile(filepath.Join("..", "..", "test", "xsd", "Fedwire_Funds_Service_Release_2025_AccountReportingRequest_camt_060_001_05_20241122_1718_iso15.xsd"))
}
