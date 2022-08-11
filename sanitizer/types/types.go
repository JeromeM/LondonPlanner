package types

import (
	"encoding/xml"
)

type TransXChange struct {
	XMLName       xml.Name      `xml:"TransXChange"`
	StopPoints    StopPoints    `xml:"StopPoints"`
	RouteSections RouteSections `xml:"RouteSections"`
}

type StopPoints struct {
	XMLName   xml.Name                `xml:"StopPoints"`
	StopPoint []AnnotatedStopPointRef `xml:"AnnotatedStopPointRef"`
}

type RouteSections struct {
}

type AnnotatedStopPointRef struct {
	XMLName xml.Name `xml:"AnnotatedStopPointRef"`
	Ref     string   `xml:"StopPointRef"`
	Name    string   `xml:"CommonName"`
}

type Stations struct {
	Station Station
}
type Station struct {
	Name string
	Ref  string
}
