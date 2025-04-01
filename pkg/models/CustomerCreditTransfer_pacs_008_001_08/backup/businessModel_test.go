package CustomerCreditTransfer_pacs_008_001_08

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"cloud.google.com/go/civil"
// )

// var businessModel = BusinessModel{
// 	XMLNameSpace: "http://example.com/ns",
// 	XMLNameLocal: "pacs 008 001 08",
// 	MsgId:        "123456789",
// 	CreDtTm:      time.Now(),
// 	NbOfTxs:      "3",
// 	SttlmMtd:     "123456789",
// 	ClrSys:       "123456789",
// 	InstrId:      "12345",
// 	EndToEndId:   "12345",
// 	TxId:         "12345",
// 	UETR:         "12345",
// 	InstrPrty:    "3",
// 	SvcLvl: []ServiceChoiceParam{
// 		{
// 			Cd:    "Service Level",
// 			Prtry: "3",
// 		},
// 		{
// 			Cd:    "Service Level",
// 			Prtry: "2",
// 		},
// 	},
// 	LclInstrm: "INST123",
// 	CtgyPurp: ServiceChoiceParam{
// 		Cd:    "Category Purpose",
// 		Prtry: "3",
// 	},
// 	IntrBkSttlmAmt: CurrencyAndAmountParam{
// 		Ccy: "USD",
// 		Cd:  350.00,
// 	},
// 	XchgRate: 0.75,
// 	ChrgBr:   "12345",
// 	ChrgsInf: []ChargeParam{
// 		{
// 			amount: CurrencyAndAmountParam{
// 				Ccy: "USD",
// 				Cd:  350.00,
// 			},
// 			identify: IdentificationParam{
// 				BICFI:    "12345",
// 				ClrSysId: "12345",
// 				MmbId:    "12345",
// 			},
// 			LEI: "12345",
// 			Nm:  "234235",
// 			PstlAdr: AddressParam{
// 				Dept:        "home",
// 				SubDept:     "second",
// 				StrtNm:      "1",
// 				BldgNb:      "256",
// 				BldgNm:      "linkon, st",
// 				Flr:         "5",
// 				PstBx:       "12345",
// 				Room:        "205",
// 				PstCd:       "001000",
// 				TwnNm:       "Dhaka",
// 				TwnLctnNm:   "office",
// 				DstrctNm:    "5",
// 				CtrySubDvsn: "polaris",
// 				Ctry:        "New York",
// 				AdrLine:     []string{"canlaed", "appilo"},
// 			},
// 		},
// 	},
// 	PrvsInstgAgt1: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	PrvsInstgAgt1Acct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	PrvsInstgAgt2: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	PrvsInstgAgt2Acct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	PrvsInstgAgt3: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	PrvsInstgAgt3Acct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	InstgAgt_ClrSysId: "12345",
// 	InstgAgt_MmbId:    "12345",
// 	InstdAgt_ClrSysId: "12345",
// 	InstdAgt_MmbId:    "12345",
// 	IntrmyAgt1: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	IntrmyAgt1Acct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	IntrmyAgt2: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	IntrmyAgt2Acct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	IntrmyAgt3: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	IntrmyAgt3Acct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	IntrBkSttlmDt: civil.Date{
// 		Year:  2025,
// 		Month: 3,
// 		Day:   20,
// 	},
// 	AccptncDtTm: time.Now(),
// 	InstdAmt: CurrencyAndAmountParam{
// 		Ccy: "USD",
// 		Cd:  350.00,
// 	},
// 	UltmtDbtr: PartyIdentificationParam{
// 		Nm: "John Doe",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 		Id_OrgId_AnyBIC: "12345",
// 		Id_OrgId_LEI:    "Somee",
// 		Id_OrgId_Othr: []OrganisationIdentificationParam{
// 			{
// 				Id: "abc2234",
// 				SchmeNm: ServiceChoiceParam{
// 					Cd:    "Pretty service",
// 					Prtry: "3",
// 				},
// 				Issr: "issr24",
// 			},
// 		},
// 		BirthDt: civil.Date{
// 			Year:  2025,
// 			Month: 3,
// 			Day:   20,
// 		},
// 		PrvcOfBirth: "California",
// 		CityOfBirth: "New York",
// 		CtryOfBirth: "USA",
// 		Othr:        []OrganisationIdentificationParam{},
// 		CtryOfRes:   "+1",
// 	},
// 	InitgPty: PartyIdentificationParam{
// 		Nm: "John Doe",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 		Id_OrgId_AnyBIC: "12345",
// 		Id_OrgId_LEI:    "Somee",
// 		Id_OrgId_Othr: []OrganisationIdentificationParam{
// 			{
// 				Id: "abc2234",
// 				SchmeNm: ServiceChoiceParam{
// 					Cd:    "Pretty service",
// 					Prtry: "3",
// 				},
// 				Issr: "issr24",
// 			},
// 		},
// 		BirthDt: civil.Date{
// 			Year:  2025,
// 			Month: 3,
// 			Day:   20,
// 		},
// 		PrvcOfBirth: "California",
// 		CityOfBirth: "New York",
// 		CtryOfBirth: "USA",
// 		Othr:        []OrganisationIdentificationParam{},
// 		CtryOfRes:   "+1",
// 	},
// 	Dbtr: PartyIdentificationParam{
// 		Nm: "John Doe",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 		Id_OrgId_AnyBIC: "12345",
// 		Id_OrgId_LEI:    "Somee",
// 		Id_OrgId_Othr: []OrganisationIdentificationParam{
// 			{
// 				Id: "abc2234",
// 				SchmeNm: ServiceChoiceParam{
// 					Cd:    "Pretty service",
// 					Prtry: "3",
// 				},
// 				Issr: "issr24",
// 			},
// 		},
// 		BirthDt: civil.Date{
// 			Year:  2025,
// 			Month: 3,
// 			Day:   20,
// 		},
// 		PrvcOfBirth: "California",
// 		CityOfBirth: "New York",
// 		CtryOfBirth: "USA",
// 		Othr:        []OrganisationIdentificationParam{},
// 		CtryOfRes:   "+1",
// 	},
// 	DbtrAcct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	DbtrAgt: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	DbtrAgtAcct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	CdtrAgt: ChargeParam{
// 		amount: CurrencyAndAmountParam{
// 			Ccy: "USD",
// 			Cd:  350.00,
// 		},
// 		identify: IdentificationParam{
// 			BICFI:    "12345",
// 			ClrSysId: "12345",
// 			MmbId:    "12345",
// 		},
// 		LEI: "12345",
// 		Nm:  "234235",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 	},
// 	CdtrAgtBrnchId: "branch1234",
// 	CdtrAgtAcct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	Cdtr: PartyIdentificationParam{
// 		Nm: "John Doe",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 		Id_OrgId_AnyBIC: "12345",
// 		Id_OrgId_LEI:    "Somee",
// 		Id_OrgId_Othr: []OrganisationIdentificationParam{
// 			{
// 				Id: "abc2234",
// 				SchmeNm: ServiceChoiceParam{
// 					Cd:    "Pretty service",
// 					Prtry: "3",
// 				},
// 				Issr: "issr24",
// 			},
// 		},
// 		BirthDt: civil.Date{
// 			Year:  2025,
// 			Month: 3,
// 			Day:   20,
// 		},
// 		PrvcOfBirth: "California",
// 		CityOfBirth: "New York",
// 		CtryOfBirth: "USA",
// 		Othr:        []OrganisationIdentificationParam{},
// 		CtryOfRes:   "+1",
// 	},
// 	CdtrAcct: CashAccountParam{
// 		Id_IBAN:     "12345",
// 		Id_Other_Id: "123456",
// 		Id_Other_SchmeNm: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Id_Other_Issr: "sdafe",
// 		Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Ccy: "asd",
// 		Nm:  "aads",
// 		Prxy_Tp: ServiceChoiceParam{
// 			Cd:    "Pretty service",
// 			Prtry: "3",
// 		},
// 		Prxy_Id: "12345",
// 	},
// 	UltmtCdtr: PartyIdentificationParam{
// 		Nm: "John Doe",
// 		PstlAdr: AddressParam{
// 			Dept:        "home",
// 			SubDept:     "second",
// 			StrtNm:      "1",
// 			BldgNb:      "256",
// 			BldgNm:      "linkon, st",
// 			Flr:         "5",
// 			PstBx:       "12345",
// 			Room:        "205",
// 			PstCd:       "001000",
// 			TwnNm:       "Dhaka",
// 			TwnLctnNm:   "office",
// 			DstrctNm:    "5",
// 			CtrySubDvsn: "polaris",
// 			Ctry:        "New York",
// 			AdrLine:     []string{"canlaed", "appilo"},
// 		},
// 		Id_OrgId_AnyBIC: "12345",
// 		Id_OrgId_LEI:    "Somee",
// 		Id_OrgId_Othr: []OrganisationIdentificationParam{
// 			{
// 				Id: "abc2234",
// 				SchmeNm: ServiceChoiceParam{
// 					Cd:    "Pretty service",
// 					Prtry: "3",
// 				},
// 				Issr: "issr24",
// 			},
// 		},
// 		BirthDt: civil.Date{
// 			Year:  2025,
// 			Month: 3,
// 			Day:   20,
// 		},
// 		PrvcOfBirth: "California",
// 		CityOfBirth: "New York",
// 		CtryOfBirth: "USA",
// 		Othr:        []OrganisationIdentificationParam{},
// 		CtryOfRes:   "+1",
// 	},
// 	InstrForCdtrAgt: []Instruction{
// 		{
// 			Cd:       "instreuction_001",
// 			InstrInf: "Instruction for InstrForCdtrAgt",
// 		},
// 	},
// 	Purp: ServiceChoiceParam{
// 		Cd:    "Category Purpose",
// 		Prtry: "3",
// 	},
// 	RgltryRptg: []RegulatoryReportingParam{
// 		{
// 			DbtCdtRptgInd:       "1234545",
// 			Authrty_Nm:          "Joe",
// 			Authrty_CountryCode: "+1",
// 			Dtls: []StructuredRegulatoryReporting3Param{
// 				{
// 					Tp: "report_type_0023",
// 					Dt: civil.Date{
// 						Year:  2025,
// 						Month: 3,
// 						Day:   20,
// 					},
// 					Ctry: "USA",
// 					Cd:   "Service2",
// 					Amt: CurrencyAndAmountParam{
// 						Ccy: "USD",
// 						Cd:  350.00,
// 					},
// 					Inf: []string{
// 						"Resund invoice1", "Revers Invoice 2",
// 					},
// 				},
// 			},
// 		},
// 	},
// 	RltdRmtInf_RmtId: "rmt_213124235",
// 	RmtLctnDtls: []RemittanceLocationDataParam{
// 		{
// 			Mtd:        "Calling",
// 			ElctrncAdr: "New York, Ad, USA",
// 			PstlAdrNm:  "001002",
// 			PstlAdr: AddressParam{
// 				Dept:        "home",
// 				SubDept:     "second",
// 				StrtNm:      "1",
// 				BldgNb:      "256",
// 				BldgNm:      "linkon, st",
// 				Flr:         "5",
// 				PstBx:       "12345",
// 				Room:        "205",
// 				PstCd:       "001000",
// 				TwnNm:       "Dhaka",
// 				TwnLctnNm:   "office",
// 				DstrctNm:    "5",
// 				CtrySubDvsn: "polaris",
// 				Ctry:        "New York",
// 				AdrLine:     []string{"canlaed", "appilo"},
// 			},
// 		},
// 	},
// 	RmtInfUstrd: "",
// }

// func TestWriteJsonXMLFromPackage(t *testing.T) {
// 	doc, err := Write(businessModel)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	jsonData, err := json.MarshalIndent(doc, "", "  ")
// 	if err != nil {
// 		fmt.Errorf("error marshal: %w", err)
// 	}
// 	jsonFileName := filepath.Join("generated", "test_package_write.json")
// 	err = os.WriteFile(jsonFileName, jsonData, 0644)
// 	if err != nil {
// 		fmt.Errorf("error writing JSON file: %w", err)
// 	}

// 	xmlData, err := xml.MarshalIndent(doc, "", "  ")
// 	if err != nil {
// 		fmt.Errorf("error marshal: %w", err)
// 	}
// 	xmlFileName := filepath.Join("generated", "test_package_write.xml")
// 	err = os.WriteFile(xmlFileName, xmlData, 0644)
// 	if err != nil {
// 		fmt.Errorf("error writing JSON file: %w", err)
// 	}
// }
// func TestWriteJsonXMLFile(t *testing.T) {

// 	doc, err := Write(businessModel)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	jsonData, err := json.MarshalIndent(doc, "", "  ")
// 	if err != nil {
// 		fmt.Errorf("error marshal: %w", err)
// 	}
// 	jsonFileName := filepath.Join("generated", "test_write.json")
// 	err = os.WriteFile(jsonFileName, jsonData, 0644)
// 	if err != nil {
// 		fmt.Errorf("error writing JSON file: %w", err)
// 	}

// 	xmlData, err := xml.MarshalIndent(doc, "", "  ")
// 	if err != nil {
// 		fmt.Errorf("error marshal: %w", err)
// 	}
// 	xmlFileName := filepath.Join("generated", "test_write.xml")
// 	err = os.WriteFile(xmlFileName, xmlData, 0644)
// 	if err != nil {
// 		fmt.Errorf("error writing JSON file: %w", err)
// 	}
// }
