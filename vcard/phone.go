package vcard

// Phone type
type Phone struct {
	Main     string `yaml:"Main"`
	Home     string `yaml:"Home"`
	Mobile   string `yaml:"Mobile"`
	Business string `yaml:"Business"`
}

func (p Phone) String() string {
	output := ""

	if len(p.Main) > 0 {
		output += "TEL;type=WORK;type=pref:" + p.Main + "\n"
	}

	if len(p.Mobile) > 0 {
		output += "TEL;type=CELL:" + p.Mobile + "\n"
	}

	if len(p.Home) > 0 {
		output += "TEL;type=HOME:" + p.Home + "\n"
	}

	if len(p.Business) > 0 {
		output += "TEL;type=WORK:" + p.Business + "\n"
	}

	return output
}
