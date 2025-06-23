package AccountReportingRequest

type CAMT_060_001_VERSION string

const (
	CAMT_060_001_02 CAMT_060_001_VERSION = "camt.060.001.02"
	CAMT_060_001_03 CAMT_060_001_VERSION = "camt.060.001.03"
	CAMT_060_001_04 CAMT_060_001_VERSION = "camt.060.001.04"
	CAMT_060_001_05 CAMT_060_001_VERSION = "camt.060.001.05"
	CAMT_060_001_06 CAMT_060_001_VERSION = "camt.060.001.06"
	CAMT_060_001_07 CAMT_060_001_VERSION = "camt.060.001.07"
)

var VersionNameSpaceMap = map[CAMT_060_001_VERSION]string{
	CAMT_060_001_02: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.02",
	CAMT_060_001_03: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.03",
	CAMT_060_001_04: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.04",
	CAMT_060_001_05: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05",
	CAMT_060_001_06: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.06",
	CAMT_060_001_07: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.07",
}
var NameSpaceVersionMap = map[string]CAMT_060_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.02": CAMT_060_001_02,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03": CAMT_060_001_03,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04": CAMT_060_001_04,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05": CAMT_060_001_05,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.06": CAMT_060_001_06,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.07": CAMT_060_001_07,
}

var VersionPathMap = map[CAMT_060_001_VERSION]map[string]any{
	CAMT_060_001_02: pathMapV2(),
	CAMT_060_001_03: pathMapV3(),
	CAMT_060_001_04: pathMapV4(),
	CAMT_060_001_05: pathMapV5(),
	CAMT_060_001_06: pathMapV6(),
	CAMT_060_001_07: pathMapV7(),
}
