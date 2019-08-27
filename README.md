# vcard generator

## Introduction
[vCard](https://en.wikipedia.org/wiki/VCard), also known as VCF (Virtual Contact File), is a file format standard for electronic business cards. vCards are often attached to e-mail messages, but can be exchanged in other ways, such as Multimedia Messaging Service (MMS), on the World Wide Web, instant messaging or through QR code. They can contain name and address information, telephone numbers, e-mail addresses, URLs, logos, photographs, and audio clips.

## vCard 4.0
The latest standard, which is built upon the RFC 6350 standard.

```
BEGIN:VCARD
VERSION:4.0
N:Gump;Forrest;;Mr.;
FN:Forrest Gump
ORG:Bubba Gump Shrimp Co.
TITLE:Shrimp Man
PHOTO;MEDIATYPE=image/gif:http://www.example.com/dir_photos/my_photo.gif
TEL;TYPE=work,voice;VALUE=uri:tel:+1-111-555-1212
TEL;TYPE=home,voice;VALUE=uri:tel:+1-404-555-1212
ADR;TYPE=WORK;PREF=1;LABEL="100 Waters Edge\nBaytown\, LA 30314\nUnited States of America":;;100 Waters Edge;Baytown;LA;30314;United States of America
ADR;TYPE=HOME;LABEL="42 Plantation St.\nBaytown\, LA 30314\nUnited States of America":;;42 Plantation St.;Baytown;LA;30314;United States of America
EMAIL:forrestgump@example.com
REV:20080424T195243Z
x-qq:21588891
END:VCARD
```

## License
The yaml package is licensed under the Apache License 2.0. Please see the LICENSE file for details.


To install it, run:
```go
go get github.com/qwertmax/vcard
```

## Example

```bash
vcard -f max.yml -o max.vcf -qr max.png
```

*max.yml*
```
FirstName:  Lee
LastName:   Taylor
MiddleName: J
Company:    Burning Buttons
Title: Golang Developer
Note: Golang the best.
Email:
  Personal: email_Personal@example.com
  Business: email_Business@example.com
HomeAddr:
  Street: "Marshal Zhukov Street 21 5th floor"
  City: "Omsk"
  State: "Omskaya oblast"
  Country: "Russia"
  ZipCode: "644024"
  CountryCode: "RU"
  Type: "HOME"
  Preferenced: n
WorkAddr:
  Street: "Marshal Zhukov Street 21 5th floor"
  City: "Omsk"
  State: "Omskaya oblast"
  Country: "Russia"
  ZipCode: "644024"
  CountryCode: "RU"
  Type: "WORK"
  Preferenced: y
Website: "https://qwertmax.ru"
Birthday:
  Year: 1900
  Month: 11
  Day: 05
Phone:
  Main: +7000000000
  Home: +7000000000
  Mobile: +7000000000
  Business: +7000000000
```

here is example of Contact QR Code

![Maxim Tishchenko](https://raw.githubusercontent.com/wiki/qwertmax/vcard/imgs/max.png)