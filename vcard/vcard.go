package vcard

import "fmt"

// Vcard type
type Vcard struct {
	FirstName     string
	LastName      string
	MiddleName    string
	Company       string
	EmailPersonal string
	EmailBusiness string
	Title         string
	HomeAddr      Address
	WorkAddr      Address
	Website       string
	Phone         Phone
	Birthday      Birthday
	Node          string
}

// Phone type
type Phone struct {
	Main     string
	Home     string
	Mobile   string
	Business string
}

// Address type
type Address struct {
	Street      string
	City        string
	State       string
	Country     string
	ZipCode     string
	CountryCode string
}

func (v Vcard) String() string {
	output := "BEGIN:VCARD\nVERSION:3.0\n" +
		"N:" + v.LastName + ";" + v.FirstName + ";" + v.MiddleName + ";\n"
	if len(v.FirstName) > 0 {
		output += "FN:" + v.FirstName + "\n"
	}

	if len(v.Company) > 0 {
		output += "ORG:" + v.Company + ";\n"
	}

	if len(v.Title) > 0 {
		output += "TITLE:" + v.Title + "\n"
	}

	if len(v.EmailBusiness) > 0 {
		output += "EMAIL;type=INTERNET;type=WORK;type=pref:" + v.EmailBusiness + "\n"
	}

	if len(v.EmailPersonal) > 0 {
		output += "EMAIL;type=INTERNET;type=HOME:" + v.EmailPersonal + "\n"
	}

	if len(v.Phone.Main) > 0 {
		output += "TEL;type=WORK;type=pref:" + v.Phone.Main + "\n"
	}

	if len(v.Phone.Mobile) > 0 {
		output += "TEL;type=CELL:" + v.Phone.Mobile + "\n"
	}

	if len(v.Phone.Home) > 0 {
		output += "TEL;type=HOME:" + v.Phone.Home + "\n"
	}

	if len(v.Phone.Business) > 0 {
		output += "TEL;type=WORK:" + v.Phone.Business + "\n"
	}

	if len(v.WorkAddr.Country) > 0 && len(v.WorkAddr.City) > 0 && len(v.WorkAddr.Street) > 0 {
		output += "item1.ADR;type=WORK:;;" + v.WorkAddr.Street + ";" + v.WorkAddr.City + ";" + v.WorkAddr.State + ";" + v.WorkAddr.ZipCode + ";" + v.WorkAddr.Country + "\n"
		if len(v.WorkAddr.CountryCode) > 0 {
			output += "item1.X-ABADR:" + v.WorkAddr.CountryCode + "\n"
		}
	}

	if len(v.HomeAddr.Country) > 0 && len(v.HomeAddr.City) > 0 && len(v.HomeAddr.Street) > 0 {
		output += "item2.ADR;type=HOME;type=pref:;;" + v.HomeAddr.Street + ";" + v.HomeAddr.City + ";" + v.HomeAddr.State + ";" + v.HomeAddr.ZipCode + ";" + v.HomeAddr.Country + "\n"
		if len(v.HomeAddr.CountryCode) > 0 {
			output += "item2.X-ABADR:" + v.HomeAddr.CountryCode + "\n"
		}
	}

	if v.Birthday.valid() {
		output += fmt.Sprint(v.Birthday)
	}

	if len(v.Node) > 0 {
		output += "NOTE:" + v.Node + "\n"
	}

	if len(v.Website) > 0 {
		output += "item3.URL;type=pref:" + v.Website + "\n"
		output += "item3.X-ABLabel:_$!<HomePage>!$_\n"
	}

	output += "END:VCARD"

	return output
}
