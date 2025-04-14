package model

import (
	"os"
	"path/filepath"
)

func WriteXMLTo(filePath string, xml []byte) error {
	os.Mkdir("generated", 0755)
	xmlFileName := filepath.Join("generated", filePath)

	return os.WriteFile(xmlFileName, xml, 0644)
}
