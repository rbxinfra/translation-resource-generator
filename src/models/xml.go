package models

import (
	"encoding/xml"
	"io"
)

// Custom Xml marshaler for specific types

type StringMap map[string]string

type stringMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// UnmarshalXML implements the xml.Unmarshaler interface for StringMap.
func (m *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringMap{}
	for {
		var e stringMapEntry
		err := d.Decode(&e)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		(*m)[e.XMLName.Local] = e.Value
	}
}

func (m StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for k, v := range m {
		e.Encode(stringMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}
	return nil
}

type StringsMap map[string]TranslationResources

type stringsMapEntry struct {
	XMLName xml.Name

	EnglishString string `xml:"en_us"`
	Translations  StringMap
}

// UnmarshalXML implements the xml.Unmarshaler interface for StringsMap.
func (m *StringsMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringsMap{}
	for {
		var e stringsMapEntry
		err := d.Decode(&e)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		(*m)[e.XMLName.Local] = TranslationResources{
			EnglishString: e.EnglishString,
			Translations:  e.Translations,
		}
	}
}

func (m StringsMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for k, v := range m {
		e.Encode(stringsMapEntry{
			XMLName:       xml.Name{Local: k},
			EnglishString: v.EnglishString,
			Translations:  v.Translations,
		})
	}
	return nil
}

type TranslationResourcesMap map[string]StringsMap

type translationResourcesMapEntry struct {
	XMLName xml.Name
	StringsMap
}

// UnmarshalXML implements the xml.Unmarshaler interface for TranslationResourcesMap.
func (m *TranslationResourcesMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = TranslationResourcesMap{}
	for {
		var e translationResourcesMapEntry
		err := d.Decode(&e)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		(*m)[e.XMLName.Local] = e.StringsMap
	}
}

func (m TranslationResourcesMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for k, v := range m {
		e.Encode(translationResourcesMapEntry{
			XMLName:    xml.Name{Local: k},
			StringsMap: v,
		})
	}
	return nil
}
