package model

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/aknopov/xmlcomparator"
)

func WriteXMLTo(filePath string, xml []byte) error {
	os.Mkdir("generated", 0755)
	xmlFileName := filepath.Join("generated", filePath)

	return os.WriteFile(xmlFileName, xml, 0644)
}
func ReadXMLFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
func removeAttributes(input []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	var output bytes.Buffer
	encoder := xml.NewEncoder(&output)
	defer encoder.Close() // Ensure the encoder is always closed

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch tok := t.(type) {
		case xml.StartElement:
			tok.Attr = nil
			encoder.EncodeToken(tok)
		case xml.EndElement, xml.CharData, xml.Comment, xml.ProcInst, xml.Directive:
			encoder.EncodeToken(tok)
		}
	}
	encoder.Flush()
	return output.Bytes(), nil
}

// Matches ISO8601 date and datetime (very basic)
var dateRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var datetimeRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}[T ]\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[+-]\d{2}:\d{2})?$`)

func isDateOrDatetime(s string) bool {
	return dateRegex.MatchString(s) || datetimeRegex.MatchString(s)
}
func removeDateValues(input []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	var output bytes.Buffer
	encoder := xml.NewEncoder(&output)

	defer encoder.Close() // Ensure the encoder is closed even if an error occurs

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch tok := t.(type) {
		case xml.StartElement:
			// Remove date/datetime attribute values
			newAttrs := make([]xml.Attr, 0, len(tok.Attr))
			for _, attr := range tok.Attr {
				if isDateOrDatetime(attr.Value) {
					attr.Value = "" // blank out
				}
				newAttrs = append(newAttrs, attr)
			}
			tok.Attr = newAttrs
			encoder.EncodeToken(tok)
		case xml.CharData:
			// If the text is a date/datetime, blank it out
			if isDateOrDatetime(strings.TrimSpace(string(tok))) {
				encoder.EncodeToken(xml.CharData([]byte("")))
			} else {
				encoder.EncodeToken(tok)
			}
		case xml.EndElement:
			encoder.EncodeToken(tok)
		default:
			encoder.EncodeToken(tok)
		}
	}
	encoder.Flush()
	return output.Bytes(), nil
}
func CompareXMLs(filePath1 string, filePath2 string) bool {
	xml1, err := ReadXMLFile(filePath1)
	if err != nil {
		return false
	}
	xml2, err := ReadXMLFile(filePath2)
	if err != nil {
		return false
	}
	clean1, err := removeAttributes(xml1)
	if err != nil {
		log.Fatal(err)
	}
	clean2, err := removeAttributes(xml2)
	if err != nil {
		log.Fatal(err)
	}
	clean1, err = removeDateValues(clean1)
	if err != nil {
		log.Fatal(err)
	}
	clean2, err = removeDateValues(clean2)
	if err != nil {
		log.Fatal(err)
	}
	diffs := xmlcomparator.CompareXmlStrings(string(clean1), string(clean2), true)
	if len(diffs) == 0 {
		log.Println("XML files are equal (ignoring date/datetime fields).")
		return true
	} else {
		log.Println("Differences found (ignoring date/datetime fields):")
		for _, diff := range diffs {
			log.Println(diff)
		}
		return false
	}
}
