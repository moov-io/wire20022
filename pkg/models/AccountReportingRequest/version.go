package ArchiveAccountReportingRequest

type CAMT_060_001_VESION string

const (
	CAMT_060_001_02 CAMT_060_001_VESION = "camt.060.001.02"
	CAMT_060_001_03 CAMT_060_001_VESION = "camt.060.001.03"
	CAMT_060_001_04 CAMT_060_001_VESION = "camt.060.001.04"
	CAMT_060_001_05 CAMT_060_001_VESION = "camt.060.001.05"
	CAMT_060_001_06 CAMT_060_001_VESION = "camt.060.001.06"
	CAMT_060_001_07 CAMT_060_001_VESION = "camt.060.001.07"
)

var VersionNameSpaceMap = map[CAMT_060_001_VESION]string{
	CAMT_060_001_02: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.02",
	CAMT_060_001_03: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.03",
	CAMT_060_001_04: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.04",
	CAMT_060_001_05: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05",
	CAMT_060_001_06: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.06",
	CAMT_060_001_07: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.07",
}
var NameSpaceVersonMap = map[string]CAMT_060_001_VESION{
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.02": CAMT_060_001_02,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03": CAMT_060_001_03,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04": CAMT_060_001_04,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05": CAMT_060_001_05,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.06": CAMT_060_001_06,
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.07": CAMT_060_001_07,
}

var VersionPathMap = map[CAMT_060_001_VESION]map[string]string{
	CAMT_060_001_02: PathMapV2(),
	CAMT_060_001_03: PathMapV3(),
	CAMT_060_001_04: PathMapV4(),
	CAMT_060_001_05: PathMapV5(),
	CAMT_060_001_06: PathMapV6(),
	CAMT_060_001_07: PathMapV7(),
}
