package models

import "encoding/xml"

type Item struct {
	XMLName             xml.Name            `xml:"assessmentItem"`
	ResponseDeclaration ResponseDeclaration `xml:"responseDeclaration"`
}

type ResponseDeclaration struct {
	CorrectResponse []string `xml:"correctResponse>value"`
	Mapping         Mapping  `xml:"mapping"`
}

type Mapping struct {
	MapEntries []MapEntry `xml:"mapEntry"`
}

type MapEntry struct {
	MapKey      string `xml:"mapKey,attr"`
	MappedValue string `xml:"mappedValue,attr"`
}
