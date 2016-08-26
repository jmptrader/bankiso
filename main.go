package main

import (
	"fmt"
	"reflect"

	"github.com/figassis/bankiso/iso"
	"github.com/figassis/bankiso/iso20022"
)

func main() {
	example := `<?xml version="1.0" encoding="UTF-8"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.009.001.04"><MndtInitnReq><GrpHdr><MsgId>BBBB654322</MsgId><CreDtTm>2013-06-10T11:00:00</CreDtTm><InitgPty><Nm>Jersey Mobile Phone</Nm><PstlAdr><StrtNm>Virginia Lane</StrtNm><BldgNb>36</BldgNb><PstCd>NJ 07311</PstCd><TwnNm>Jersey City</TwnNm><Ctry>US</Ctry></PstlAdr></InitgPty></GrpHdr><Mndt><MndtReqId>Johns/005</MndtReqId><Ocrncs><SeqTp>RCUR</SeqTp><Frqcy><Tp>MNTH</Tp></Frqcy><FrstColltnDt>2013-06-25</FrstColltnDt></Ocrncs><Cdtr><Nm>Jersey Mobile Phone</Nm></Cdtr><CdtrAcct><Id><Othr><Id>76543</Id></Othr></Id></CdtrAcct><CdtrAgt><FinInstnId><BICFI>DDDDUS31</BICFI></FinInstnId></CdtrAgt><Dbtr><Nm>Johnson</Nm></Dbtr><DbtrAcct><Id><Othr><Id>5544732</Id></Othr></Id></DbtrAcct><DbtrAgt><FinInstnId><BICFI>FFFFUS91</BICFI></FinInstnId></DbtrAgt><RfrdDoc><Tp><CdOrPrtry><Cd>DISP</Cd></CdOrPrtry></Tp><Nb>JMP/24653</Nb><RltdDt>2013-06-11</RltdDt></RfrdDoc></Mndt></MndtInitnReq></Document>`

	message, ok := iso.DecodeISO20022(example)

	if false {
		fmt.Println("ISO Main: ", reflect.TypeOf(message), ok)
	}

	//Type assertion to use in specific function
	message_struct := message.(*iso20022.PAIN00900104)
	message_struct.Name = "Assis"

	fmt.Println("ISO Main: ", reflect.TypeOf(message), message_struct, ok)

}
