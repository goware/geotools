package geoy

import (
	"encoding/json"

	"googlemaps.github.io/maps"
)

// Address represents a normalized address.
type Address struct {
	Street      string `json:"street,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	Formatted   string `json:"formatted"`
}

// Place represents a physical location.
type Place struct {
	PlaceID     string    `json:"place_id"`
	Name        string    `json:"name"`
	Address     Address   `json:"address"`
	Location    *Point    `json:"location"`
	BoundingBox *Envelope `json:"bounding_box"`
}

func (p Place) String() string {
	b, _ := json.MarshalIndent(p, "", "\t")
	return string(b)
}

func toPlace(r *result) *Place {
	place := Place{
		PlaceID:  r.PlaceID,
		Name:     gPlaceName(r.AddressComponents),
		Location: gLatLngToPoint(r.Geometry.Location),
	}

	place.Address = Address{
		Formatted: r.FormattedAddress,
	}

	for _, c := range r.AddressComponents {
		if matchAnyType([]string{"street_address", "route"}, c.Types) {
			place.Address.Street = c.LongName
		}
		if matchAnyType([]string{"house_number", "street_number"}, c.Types) {
			place.Address.HouseNumber = c.LongName
		}
		if matchAnyType([]string{"sublocality", "locality", "postal_town"}, c.Types) {
			place.Address.City = c.LongName
		}
		if matchAnyType([]string{"administrative_area_level_1"}, c.Types) {
			place.Address.State = c.LongName
		}
		if matchAnyType([]string{"country"}, c.Types) {
			place.Address.Country = c.LongName
		}
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

func matchAnyType(types []string, cmp []string) bool {
	for _, t := range types {
		for _, c := range cmp {
			if c == t {
				return true
			}
		}
	}
	return false
}
