package spold1

import (
	"encoding/xml"
)

// EcoSpold is the root element of the XML document it can have multiple data
// sets.
type EcoSpold struct {
	XMLName  xml.Name  `xml:"http://www.EcoInvent.org/EcoSpold01 ecoSpold"`
	DataSets []DataSet `xml:"dataset"`
}

// DataSet contains the content of a process data set.
type DataSet struct {
	Number            int                `xml:"numer,attr"`
	ReferenceFunction *ReferenceFunction `xml:"metaInformation>processInformation>referenceFunction"`
	Geography         *Geography         `xml:"metaInformation>processInformation>geography"`
	Time              *Time              `xml:"metaInformation>processInformation>timePeriod"`
	Exchanges         []Exchange         `xml:"flowData>exchange"`
}

// The ReferenceFunction element of a data set.
type ReferenceFunction struct {
	Name        string  `xml:"name,attr"`
	Amount      float64 `xml:"amount,attr"`
	Unit        string  `xml:"unit,attr"`
	Category    string  `xml:"category,attr"`
	SubCategory string  `xml:"subCategory,attr"`
	Comment     string  `xml:"generalComment,attr,omitempty"`
}

// The Geography element of a data set
type Geography struct {
	Location string `xml:"location,attr"`
	Text     string `xml:"text,attr,omitempty"`
}

// Time is the `timePeriod` element of the data set.
type Time struct {
	Text      string `xml:"text,attr,omitempty"`
	StartYear int    `xml:"startYear"`
	EndYear   int    `xml:"endYear"`
}

// An Exchange is an input or output of a flow.
type Exchange struct {
	Number      int     `xml:"numer,attr"`
	Name        string  `xml:"name,attr"`
	Amount      float64 `xml:"meanValue,attr"`
	Unit        string  `xml:"unit,attr"`
	Category    string  `xml:"category,attr"`
	SubCategory string  `xml:"subCategory,attr"`
	Location    string  `xml:"location,attr"`
	InputGroup  *int    `xml:"inputGroup"`
	OutputGroup *int    `xml:"outputGroup"`
}
