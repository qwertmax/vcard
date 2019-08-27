package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/qwertmax/vcard/vcard"
	"github.com/skip2/go-qrcode"
	"gopkg.in/yaml.v2"
)

func main() {
	ymlFile := flag.String("f", "", "input file name [-f example.yml]")
	vcfFile := flag.String("o", "out.vcf", "output file name")
	qrFile := flag.String("qr", "v_card.png", "QR Code image file")
	flag.Parse()

	if len(*ymlFile) == 0 {
		fmt.Println("specify yml file '-f example.yml'")
		return
	}

	yamlFile, err := ioutil.ReadFile(*ymlFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	vcard := vcard.Vcard{}

	err = yaml.Unmarshal(yamlFile, &vcard)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(*vcfFile, []byte(fmt.Sprint(vcard)), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(*qrFile) > 0 {
		qrcode.WriteFile(fmt.Sprint(vcard), qrcode.High, 256, *qrFile)
	}
}
