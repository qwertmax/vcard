package main

import (
	"fmt"

	"github.com/qwertmax/vcard-generator/vcard"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	vcard := vcard.Vcard{
		FirstName:     "Max",
		LastName:      "Tishchenko",
		MiddleName:    "Nikolaevish",
		Company:       "Burning Buttons",
		EmailBusiness: "max@burningbuttons.com",
		EmailPersonal: "qwertmax@gmail.com",
		Title:         "CEO",
		Node:          "For more than 10 years I've been helping our clients to move forward from ideas to solutions by providing highly profitable applications.",
		WorkAddr: vcard.Address{
			Street:      "Mitrofanevskoe shosse 2k2 office 209",
			City:        "Saint-Petersburg",
			Country:     "Russia",
			CountryCode: "RU",
			State:       "Leningradskaya Oblast",
			ZipCode:     "190005",
		},
		Website: "https://burningbuttons.com",
		HomeAddr: vcard.Address{
			Street:      "nabereznaya obvodnogo kanala 108 apt 469",
			City:        "Saint-Petersburg",
			Country:     "Russia",
			CountryCode: "RU",
			State:       "Leningradskaya Oblast",
			ZipCode:     "196084",
		},
		Birthday: vcard.Birthday{
			Year:  1985,
			Month: 12,
			Day:   25,
		},
	}

	// str := "BEGIN:VCARD\nVERSION:3.0\nN:Doe;John;;;\nFN:John Doe\nORG:Example.com Inc.;\nTITLE:Imaginary test person\nEMAIL;type=INTERNET;type=WORK;type=pref:johnDoe@example.org\nTEL;type=WORK;type=pref:+1 617 555 1212\nTEL;type=CELL:+1 781 555 1212\nTEL;type=HOME:+1 202 555 1212\nTEL;type=WORK:+1 (617) 555-1234\nitem1.ADR;type=WORK:;;2 Example Avenue;Anytown;NY;01111;USA\nitem1.X-ABADR:us\nitem2.ADR;type=HOME;type=pref:;;3 Acacia Avenue;Newtown;MA;02222;USA\nitem2.X-ABADR:us\nNOTE:John Doe has a long and varied history\\, being documented on more police files that anyone else. Reports of his death are alas numerous.\nitem3.URL;type=pref:http\\://www.example/com/doe\nitem3.X-ABLabel:_$!<HomePage>!$_\nitem4.URL:http\\://www.example.com/Joe/foaf.df\nitem4.X-ABLabel:FOAF\nitem5.X-ABRELATEDNAMES;type=pref:Jane Doe\nitem5.X-ABLabel:_$!<Friend>!$_\nCATEGORIES:Work,Test group\nX-ABUID:5AD380FD-B2DE-4261-BA99-DE1D1DB52FBE\\:ABPerson\nEND:VCARD"
	// qrcode.WriteFile(str, qrcode.High, 256, "qr.png")
	qrcode.WriteFile(fmt.Sprint(vcard), qrcode.High, 256, "qr.png")
}
