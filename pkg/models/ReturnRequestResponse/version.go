package ReturnRequestResponse

type CAMT_029_001_VESION string

const (
	CAMT_029_001_03 CAMT_029_001_VESION = "camt.029.001.03"
	CAMT_029_001_04 CAMT_029_001_VESION = "camt.029.001.04"
	CAMT_029_001_05 CAMT_029_001_VESION = "camt.029.001.05"
	CAMT_029_001_06 CAMT_029_001_VESION = "camt.029.001.06"
	CAMT_029_001_07 CAMT_029_001_VESION = "camt.029.001.07"
	CAMT_029_001_08 CAMT_029_001_VESION = "camt.029.001.08"
	CAMT_029_001_09 CAMT_029_001_VESION = "camt.029.001.09"
	CAMT_029_001_10 CAMT_029_001_VESION = "camt.029.001.10"
	CAMT_029_001_11 CAMT_029_001_VESION = "camt.029.001.11"
	CAMT_029_001_12 CAMT_029_001_VESION = "camt.029.001.12"
)

var VersionNameSpaceMap = map[CAMT_029_001_VESION]string{
	CAMT_029_001_03: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.03",
	CAMT_029_001_04: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.04",
	CAMT_029_001_05: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.05",
	CAMT_029_001_06: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.06",
	CAMT_029_001_07: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.07",
	CAMT_029_001_08: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.08",
	CAMT_029_001_09: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09",
	CAMT_029_001_10: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.10",
	CAMT_029_001_11: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.11",
	CAMT_029_001_12: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.12",
}

var NameSpaceVersonMap = map[string]CAMT_029_001_VESION{
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.03": CAMT_029_001_03,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.04": CAMT_029_001_04,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.05": CAMT_029_001_05,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.06": CAMT_029_001_06,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.07": CAMT_029_001_07,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.08": CAMT_029_001_08,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09": CAMT_029_001_09,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.10": CAMT_029_001_10,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.11": CAMT_029_001_11,
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.12": CAMT_029_001_12,
}

var VersionPathMap = map[CAMT_029_001_VESION]map[string]string{
	CAMT_029_001_03: PathMapV3(),
	CAMT_029_001_04: PathMapV4(),
	CAMT_029_001_05: PathMapV5(),
	CAMT_029_001_06: PathMapV6(),
	CAMT_029_001_07: PathMapV7(),
	CAMT_029_001_08: PathMapV8(),
	CAMT_029_001_09: PathMapV9(),
	CAMT_029_001_10: PathMapV10(),
	CAMT_029_001_11: PathMapV11(),
	CAMT_029_001_12: PathMapV12(),
}
