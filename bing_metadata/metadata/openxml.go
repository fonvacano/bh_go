package metadata

import (
	"archive/zip"
	"encoding/xml"
	"strings"
)

type OfficeCoreProperty struct {
	XMLName        xml.Name `xml:"coreProperties"`
	Creator        string   `xml:"creator"`
	LastModifiedBy string   `xml:"lastModifiedBy"`
}

type OfficeAppProperty struct {
	XMLName     xml.Name `xml:"Properties"`
	Application string   `xml:"Application"`
	Company     string   `xml:"Company"`
	Version     string   `xml:"Version"`
}

var OfficeVersions = map[string]string{
	"16": "2016",
	"15": "2013",
	"14": "2010",
	"12": "2007",
	"11": "2003",
}

func (a *OfficeAppProperty) GetMajorVersion() string {
	tokens := strings.Split(a.Version, ".")
	if len(tokens) != 2 {
		return "Unknown"
	}

	v, ok := OfficeVersions[tokens[0]]

	if !ok {
		return "Unknown"
	}
	return v
}

func NewProperties(r *zip.Reader) (*OfficeCoreProperty, *OfficeAppProperty, error) {
	var coreProps OfficeCoreProperty
	var appProps OfficeAppProperty

	for _, file := range r.File {
		switch file.Name {
		case "docProps/core.xml":
			if err := process(file, &coreProps); err != nil {
				return nil, nil, err
			}
		case "docProps/app.xml":
			if err := process(file, &appProps); err != nil {
				return nil, nil, err
			}
		default:
			continue
		}
	}
	return &coreProps, &appProps, nil
}

func process(f *zip.File, s interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}

	defer rc.Close()

	if err := xml.NewDecoder(rc).Decode(s); err != nil {
		return err
	}
	return nil
}
