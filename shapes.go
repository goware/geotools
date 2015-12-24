package geoy

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"googlemaps.github.io/maps"
)

const (
	defaultSRID = 4326 // https://en.wikipedia.org/wiki/World_Geodetic_System
)

type pointEWKB struct {
	Endiness byte
	Type     uint32
	SRID     uint32
	X        float64
	Y        float64
}

// LatLnger defines a struct that can convert itself to cartesian coordinates.
type LatLnger interface {
	LatLng() []float64
}

// LatLng implements LatLnger
type LatLng maps.LatLng

// LatLng returns an array of [lat, lon]
func (l LatLng) LatLng() []float64 {
	return []float64{l.Lng, l.Lat}
}

// Point is a standard GeoJSON 2d Point with x,y coordinates
type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// NewPoint creates a GeoJSON 2d point.
func NewPoint(x, y float64) *Point {
	p := Point{Type: "point", Coordinates: []float64{x, y}}
	return &p
}

// LatLng implements LatLngr.
func (p Point) LatLng() []float64 {
	return []float64{p.Coordinates[1], p.Coordinates[0]}
}

// MarshalDB prepares the point to be stored.
func (p Point) MarshalDB() (interface{}, error) {
	return p.WKT(), nil
}

// UnmarshalDB converts an stored point into a Point struct.
func (p *Point) UnmarshalDB(v interface{}) error {
	s := string(v.([]byte))

	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	buf := bytes.NewReader(b)
	var ewkb pointEWKB
	err = binary.Read(buf, binary.LittleEndian, &ewkb)
	if err != nil {
		return err
	}
	p.Type = "point"
	p.Coordinates = []float64{ewkb.X, ewkb.Y}
	return nil
}

// WKT implements Geometry.
func (p Point) WKT() string {
	return fmt.Sprintf("SRID=%d;POINT(%0.6f %0.6f)", defaultSRID, p.Coordinates[0], p.Coordinates[1])
}

// String returns a text representation of the point.
func (p Point) String() string {
	return p.WKT()
}

// PointFromLatLng converts a latLng to a cartesian Point.
func PointFromLatLng(latlon LatLnger) *Point {
	l := latlon.LatLng()
	x, y := l[1], l[0]
	return NewPoint(x, y)
}

// LatLngFromPoint converts a Point to a LatLng.
func LatLngFromPoint(p Point) *LatLng {
	return &LatLng{Lat: p.Coordinates[1], Lng: p.Coordinates[0]}
}

// InstagramLocation is an object representing the location information returned by the Instagram API
type InstagramLocation struct {
	Latitude  float64
	Longitude float64
	ID        string
	Name      string
}

// LatLng implements LatLnger
func (l InstagramLocation) LatLng() []float64 {
	return []float64{l.Latitude, l.Longitude}
}

// FacebookLocation is an object representing the location information returned
// by the Facebook API
type FacebookLocation InstagramLocation

// TwitterLocation is an object representing the location information returned
// by the Twitter API (a GeoJSON Point)
type TwitterLocation Point

// Geometry defines WKT, which provides a text representation of a vector.
type Geometry interface {
	WKT() string
}

// Envelope is a GeoJSON like shape where coordinates contains [[left, top],
// [right, bottom]]
type Envelope struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

// NewEnvelope creates an envelope.
func NewEnvelope(left, top, right, bottom float64) *Envelope {
	e := Envelope{Type: "envelope", Coordinates: [][]float64{[]float64{left, top}, []float64{right, bottom}}}
	return &e
}

// WKT implements Geometry.
func (e Envelope) WKT() string {
	l, t := e.Coordinates[0][0], e.Coordinates[0][1]
	r, b := e.Coordinates[1][0], e.Coordinates[1][1]
	return fmt.Sprintf("POLYGON((%0.6f %0.6f, %0.6f %0.6f, %0.6f %0.6f, %0.6f %0.6f, %0.6f %0.6f))", l, t, l, b, r, b, r, t, l, t)
}

// MarshalDB implements db.Marshaler
func (e Envelope) MarshalDB() (interface{}, error) {
	return e.WKT(), nil
}

// Implements fmt.Stringer
func (e Envelope) String() string {
	return e.WKT()
}
