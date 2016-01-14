package geoy

import (
	"encoding/json"
	"fmt"
	"sort"

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

func (a Address) String() string {
	return a.Formatted
}

// Place represents a physical location.
type Place struct {
	PlaceID     string    `json:"place_id"`
	Name        string    `json:"name"`
	Address     Address   `json:"address"`
	Location    *Point    `json:"location"`
	BoundingBox *Envelope `json:"bounding_box"`
}

type orderedAddressComponents []maps.AddressComponent

func (c *orderedAddressComponents) Len() int {
	ac := []maps.AddressComponent(*c)
	return len(ac)
}

func (c *orderedAddressComponents) Less(i, j int) bool {
	ac := []maps.AddressComponent(*c)
	return ac[i].LongName < ac[j].LongName
}

func (c *orderedAddressComponents) Swap(i, j int) {
	ac := []maps.AddressComponent(*c)
	ac[i], ac[j] = ac[j], ac[i]
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

	components := orderedAddressComponents(r.AddressComponents)

	// Making sure address components are always in the same order.
	sort.Sort(&components)

	for _, c := range components {
		dest := func() *string {
			// Mapping address types to fields. See https://github.com/pressly/geoy/issues/9.
			switch true {
			case matchAnyType([]string{"street_address", "route", "premise", "subpremise"}, c.Types):
				return &place.Address.Street
			case matchAnyType([]string{"house_number", "street_number"}, c.Types):
				return &place.Address.HouseNumber
			case matchAnyType([]string{"sublocality", "locality", "postal_town"}, c.Types):
				return &place.Address.City
			case matchAnyType([]string{"administrative_area_level_1"}, c.Types):
				return &place.Address.State
			case matchAnyType([]string{"country", "colloquial_area"}, c.Types):
				return &place.Address.Country
			}
			return nil
		}()

		if dest != nil {
			if *dest == "" {
				*dest = c.LongName
			} else {
				*dest = fmt.Sprintf("%s, %s", *dest, c.LongName)
			}
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
