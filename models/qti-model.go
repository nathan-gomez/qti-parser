package models

import "encoding/xml"

type QtiItem struct {
	XMLName             xml.Name              `xml:"assessmentItem"`
	ResponseDeclaration []ResponseDeclaration `xml:"responseDeclaration"`
	OutcomeDeclaration  []OutcomeDeclaration  `xml:"outcomeDeclaration"`
}

type OutcomeDeclaration struct {
	Identifier   string   `xml:"identifier,attr"`
	Cardinality  string   `xml:"cardinality,attr"`
	BaseType     string   `xml:"baseType,attr"`
	DefaultValue []string `xml:"defaultValue>value"`
}

type ResponseDeclaration struct {
	Identifier      string   `xml:"identifier,attr"`
	Cardinality     string   `xml:"cardinality,attr"`
	BaseType        string   `xml:"baseType,attr"`
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
