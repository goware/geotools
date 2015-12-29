package geoy

import (
	"encoding/json"
	"googlemaps.github.io/maps"
)

// AddressComponent represents a part of an address.
type AddressComponent struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Place represents a physical location.
type Place struct {
	PlaceID           string             `json:"place_id"`
	Name              string             `json:"name"`
	AddressComponents []AddressComponent `json:"address_components"`
	AddressString     string             `json:"address_string"`
	Location          *Point             `json:"location"`
	BoundingBox       *Envelope          `json:"bounding_box"`
}

func (p Place) String() string {
	b, _ := json.MarshalIndent(p, "", "\t")
	return string(b)
}

func (r result) toPlace() *Place {
	place := Place{
		PlaceID:       r.PlaceID,
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
	PlaceID           string
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
