package vcard

// Address type
type Address struct {
	Street      string
	City        string
	State       string
	Country     string
	ZipCode     string
	CountryCode string
	Type        string
	Preferenced bool
}

func (v Address) String() string {
	output := ""

	if len(v.Country) > 0 && len(v.City) > 0 && len(v.Street) > 0 {
		preferenced := ""
		if v.Preferenced {
			preferenced = "type=pref"
		}

		output += "item2.ADR;type=" + v.Type + ";" + preferenced + ":;;" + v.Street + ";" + v.City + ";" + v.State + ";" + v.ZipCode + ";" + v.Country + "\n"
		if len(v.CountryCode) > 0 {
			output += "item2.X-ABADR:" + v.CountryCode + "\n"
		}
	}

	return output
}
