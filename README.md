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

max.yml
```
FirstName:  Lee
LastName:   Taylor
MiddleName: No Idea
Company:    Burning Buttons
Title: CEO
Note: For more than 10 years I've been helping our clients to move forward from ideas to solutions by providing highly profitable applications.
Email:
  Personal: email_Personal@example.com
  Business: email_Business@example.com
HomeAddr:
  Street: "qqq"
  City: "qqq"
  State: "qqq"
  Country: "qqq"
  ZipCode: "qqq"
  CountryCode: "qqq"
  Type: "HOME"
  Preferenced: n
WorkAddr:
  Street: "qqq"
  City: "qqq"
  State: "qqq"
  Country: "qqq"
  ZipCode: "qqq"
  CountryCode: "qqq"
  Type: "WORK"
  Preferenced: y
Website: "https://example.com"
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

