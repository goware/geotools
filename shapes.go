package geoy

type LatLoner interface {
	LatLon() []float64
}

type LatLon struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// LatLon returns an array of [lat, lon]
func (l LatLon) LatLon() []float64 {
	return []float64{l.Lon, l.Lat}
}

// Point is a standard GeoJSON 2d Point with x,y coordinates
type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func NewPoint(x, y float64) *Point {
	p := Point{Type: "point", Coordinates: []float64{x, y}}
	return &p
}

// LatLon returns an array of [lat, lon]
func (p Point) LatLon() []float64 {
	return []float64{p.Coordinates[1], p.Coordinates[0]}
}

// Envelope is a GeoJSON like shape where coordinates contains [[left, top], [right, bottom]]
type Envelope struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

func NewEnvelope(left, top, right, bottom float64) *Envelope {
	e := Envelope{Type: "envelope", Coordinates: [][]float64{[]float64{left, top}, []float64{right, bottom}}}
	return &e
}

// PointFromLatLon converts a latLon to a xy Point
func PointFromLatLon(latlon LatLoner) *Point {
	l := latlon.LatLon()
	x, y := l[1], l[0]
	return NewPoint(x, y)
}

// LatLonFromPoint converts a Point to a LatLon
func LatLonFromPoint(p Point) *LatLon {
	return &LatLon{Lat: p.Coordinates[1], Lon: p.Coordinates[0]}
}

// InstagramLocation is an object representing the location information returned by the Instagram API
type InstagramLocation struct {
	Latitude  float64
	Longitude float64
	Id        string
	Name      string
}

func (l InstagramLocation) LatLon() []float64 {
	return []float64{l.Latitude, l.Longitude}
}

// FacebookLocation is an object representing the location information returned by the Facebook API
type FacebookLocation InstagramLocation

// TwitterLocation is an object representing the location information returned by the Twitter API (a GeoJSON Point)
type TwitterLocation Point
