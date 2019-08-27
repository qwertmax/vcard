package vcard

// Email type
type Email struct {
	Personal string `yaml:"Personal"`
	Business string `yaml:"Business"`
}

func (e Email) String() string {
	output := ""

	if len(e.Business) > 0 {
		output += "EMAIL;type=INTERNET;type=WORK;type=pref:" + e.Business + "\n"
	}

	if len(e.Personal) > 0 {
		output += "EMAIL;type=INTERNET;type=HOME:" + e.Personal + "\n"
	}

	return output
}
