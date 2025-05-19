package models

import "encoding/xml"

// ResxData represents the data element from a ResX file.
type ResxData struct {
	XMLName xml.Name `xml:"data"`

	// Represents the name attribute
	NameAttr string `xml:"name,attr"`

	// Represents the space setting attribute
	XmlSpaceAttr string `xml:"xml:space,attr,omitempty"`

	// Gets the value element.
	Value string `xml:"value"`

	// Gets the comment element.
	Comment string `xml:"comment"`
}

type ResxRoot struct {
	XMLName xml.Name    `xml:"root"`

	// Data entries
	Data    []*ResxData `xml:"data"`
}
