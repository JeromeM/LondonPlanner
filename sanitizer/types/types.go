package types

import (
	"encoding/xml"
)

type TransXChange struct {
	XMLName       xml.Name      `xml:"TransXChange"`
	StopPoints    StopPoints    `xml:"StopPoints"`
	RouteSections RouteSections `xml:"RouteSections"`
	Services      Services      `xml:"Services"`
}

type StopPoints struct {
	XMLName   xml.Name                `xml:"StopPoints"`
	StopPoint []AnnotatedStopPointRef `xml:"AnnotatedStopPointRef"`
}

type AnnotatedStopPointRef struct {
	XMLName xml.Name `xml:"AnnotatedStopPointRef"`
	Ref     string   `xml:"StopPointRef"`
	Name    string   `xml:"CommonName"`
}

type Services struct {
	XMLName xml.Name  `xml:"Services"`
	Service []Service `xml:"Service"`
}

type Service struct {
	XMLName   xml.Name `xml:"Service"`
	Line_name string   `xml:"Lines>Line>LineName"`
}

type RouteSections struct {
}

type Stations struct {
	Station Station
}
type Station struct {
	Name string
	Ref  string
}
