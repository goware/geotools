package geoy

import (
	"encoding/json"
	"googlemaps.github.io/maps"
)

// AddressComponent represents a part of an address.
type AddressComponent struct {
	Name string
	Type string
}

// Place represents a physical location.
type Place struct {
	Name              string
	AddressComponents []AddressComponent
	AddressString     string
	Location          *Point
	BoundingBox       *Envelope
}

func (p Place) String() string {
	b, _ := json.MarshalIndent(p, "", "\t")
	return string(b)
}

func (r result) toPlace() *Place {
	place := Place{
		Name:          gPlaceName(r.AddressComponents),
		Location:      gLatLngToPoint(r.Geometry.Location),
		AddressString: r.FormattedAddress,
	}
	place.AddressComponents = make([]AddressComponent, len(r.AddressComponents))
	for i, c := range r.AddressComponents {
		place.AddressComponents[i].Name = c.LongName
		place.AddressComponents[i].Type = c.Types[0]
	}

	ne := r.Geometry.Viewport.NorthEast
	sw := r.Geometry.Viewport.SouthWest

	place.BoundingBox = NewEnvelope(sw.Lng, ne.Lat, ne.Lng, sw.Lat)

	return &place
}

type result struct {
	AddressComponents []maps.AddressComponent
	Geometry          maps.AddressGeometry
	FormattedAddress  string
}

func gLatLngToPoint(ll maps.LatLng) *Point {
	return NewPoint(ll.Lng, ll.Lat)
}

func gPlaceName(ac []maps.AddressComponent) string {
	if len(ac) > 0 {
		return ac[0].LongName
	}
	return ""
}
