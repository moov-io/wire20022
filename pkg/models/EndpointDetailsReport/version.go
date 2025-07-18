package EndpointDetailsReport

type CAMT_052_001_VERSION string

const (
	CAMT_052_001_02 CAMT_052_001_VERSION = "camt.052.001.02"
	CAMT_052_001_03 CAMT_052_001_VERSION = "camt.052.001.03"
	CAMT_052_001_04 CAMT_052_001_VERSION = "camt.052.001.04"
	CAMT_052_001_05 CAMT_052_001_VERSION = "camt.052.001.05"
	CAMT_052_001_06 CAMT_052_001_VERSION = "camt.052.001.06"
	CAMT_052_001_07 CAMT_052_001_VERSION = "camt.052.001.07"
	CAMT_052_001_08 CAMT_052_001_VERSION = "camt.052.001.08"
	CAMT_052_001_09 CAMT_052_001_VERSION = "camt.052.001.09"
	CAMT_052_001_10 CAMT_052_001_VERSION = "camt.052.001.10"
	CAMT_052_001_11 CAMT_052_001_VERSION = "camt.052.001.11"
	CAMT_052_001_12 CAMT_052_001_VERSION = "camt.052.001.12"
)

var VersionNameSpaceMap = map[CAMT_052_001_VERSION]string{
	CAMT_052_001_02: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.02",
	CAMT_052_001_03: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.03",
	CAMT_052_001_04: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.04",
	CAMT_052_001_05: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.05",
	CAMT_052_001_06: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.06",
	CAMT_052_001_07: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.07",
	CAMT_052_001_08: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
	CAMT_052_001_09: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.09",
	CAMT_052_001_10: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.10",
	CAMT_052_001_11: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.11",
	CAMT_052_001_12: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.12",
}
var NameSpaceVersionMap = map[string]CAMT_052_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.02": CAMT_052_001_02,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.03": CAMT_052_001_03,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.04": CAMT_052_001_04,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.05": CAMT_052_001_05,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.06": CAMT_052_001_06,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.07": CAMT_052_001_07,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08": CAMT_052_001_08,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.09": CAMT_052_001_09,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.10": CAMT_052_001_10,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.11": CAMT_052_001_11,
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.12": CAMT_052_001_12,
}
var VersionPathMap = map[CAMT_052_001_VERSION]map[string]any{
	CAMT_052_001_02: pathMapV2(),
	CAMT_052_001_03: pathMapV3(),
	CAMT_052_001_04: pathMapV4(),
	CAMT_052_001_05: pathMapV5(),
	CAMT_052_001_06: pathMapV6(),
	CAMT_052_001_07: pathMapV7(),
	CAMT_052_001_08: pathMapV8(),
	CAMT_052_001_09: pathMapV9(),
	CAMT_052_001_10: pathMapV10(),
	CAMT_052_001_11: pathMapV11(),
	CAMT_052_001_12: pathMapV12(),
}
