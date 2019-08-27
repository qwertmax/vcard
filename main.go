package main

import (
	"fmt"
	"io/ioutil"

	"github.com/qwertmax/vcard-generator/vcard"
	"gopkg.in/yaml.v2"
)

func main() {
	// vcard := vcard.Vcard{
	// 	FirstName:  "Jef",
	// 	LastName:   "Lee",
	// 	MiddleName: "Taylor",
	// 	Company:    "Burning Buttons",
	// 	Email: vcard.Email{
	// 		Business: "max@example.com",
	// 		Personal: "test@gmail.com",
	// 	},
	// 	Title: "CEO",
	// 	Node:  "For more than 10 years I've been helping our clients to move forward from ideas to solutions by providing highly profitable applications.",
	// 	WorkAddr: vcard.Address{
	// 		Street:      "TRA TA TA",
	// 		City:        "Saint-Petersburg",
	// 		Country:     "Russia",
	// 		CountryCode: "RU",
	// 		State:       "Leningradskaya Oblast",
	// 		ZipCode:     "190005",
	// 		Type:        "WORK",
	// 		Preferenced: true,
	// 	},
	// 	Website: "https://burningbuttons.com",
	// 	HomeAddr: vcard.Address{
	// 		Street:      "example",
	// 		City:        "Saint-Petersburg",
	// 		Country:     "Russia",
	// 		CountryCode: "RU",
	// 		State:       "Leningradskaya Oblast",
	// 		ZipCode:     "196084",
	// 		Type:        "HOME",
	// 	},
	// 	Birthday: vcard.Birthday{
	// 		Year:  1985,
	// 		Month: 12,
	// 		Day:   25,
	// 	},
	// 	Phone: vcard.Phone{
	// 		Main:     "+70000000000",
	// 		Business: "+70000000001",
	// 		Home:     "+70000000002",
	// 	},
	// }

	yamlFile, err := ioutil.ReadFile("max.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	v := vcard.Vcard{}

	err = yaml.Unmarshal(yamlFile, &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)

	// fmt.Println(vcard)
	// qrcode.WriteFile(fmt.Sprint(vcard), qrcode.High, 256, "qr.png")
}
