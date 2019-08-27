package vcard

import "fmt"

// Vcard type
type Vcard struct {
	FirstName  string `yaml:"FirstName"`
	LastName   string `yaml:"LastName"`
	MiddleName string `yaml:"MiddleName"`
	Company    string `yaml:"Company"`
	Email      Email
	Title      string `yaml:"Title"`
	HomeAddr   Address
	WorkAddr   Address
	Website    string `yaml:"Website"`
	Phone      Phone
	Birthday   Birthday
	Node       string `yaml:"Note"`
}

func (v Vcard) String() string {
	output := "BEGIN:VCARD\nVERSION:4.0\n" +
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

	output += fmt.Sprint(v.Email)
	output += fmt.Sprint(v.Phone)
	output += fmt.Sprint(v.WorkAddr)
	output += fmt.Sprint(v.HomeAddr)

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
