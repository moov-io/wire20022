package PaymentStatusRequest

type PACS_028_001_VERSION string

const (
	PACS_028_001_01 PACS_028_001_VERSION = "pacs.028.001.01"
	PACS_028_001_02 PACS_028_001_VERSION = "pacs.028.001.02"
	PACS_028_001_03 PACS_028_001_VERSION = "pacs.028.001.03"
	PACS_028_001_04 PACS_028_001_VERSION = "pacs.028.001.04"
	PACS_028_001_05 PACS_028_001_VERSION = "pacs.028.001.05"
	PACS_028_001_06 PACS_028_001_VERSION = "pacs.028.001.06"
)

var VersionNameSpaceMap = map[PACS_028_001_VERSION]string{
	PACS_028_001_01: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.01",
	PACS_028_001_02: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.02",
	PACS_028_001_03: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03",
	PACS_028_001_04: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04",
	PACS_028_001_05: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.05",
	PACS_028_001_06: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.06",
}

var NameSpaceVersionMap = map[string]PACS_028_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.01": PACS_028_001_01,
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.02": PACS_028_001_02,
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03": PACS_028_001_03,
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04": PACS_028_001_04,
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.05": PACS_028_001_05,
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.06": PACS_028_001_06,
}
var VersionPathMap = map[PACS_028_001_VERSION]map[string]any{
	PACS_028_001_01: pathMapV1(),
	PACS_028_001_02: pathMapV2(),
	PACS_028_001_03: pathMapV3(),
	PACS_028_001_04: pathMapV4(),
	PACS_028_001_05: pathMapV5(),
	PACS_028_001_06: pathMapV6(),
}
