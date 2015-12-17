package geoy

import (
	"encoding/json"
	"github.com/pressly/geoy/gmaps"
)

var (
	gmapsApiClient *gmaps.MapsApiClient
)

const KM = 1000

func init() {
	gmapsApiClient = &gmaps.MapsApiClient{}
}

// Set the Google Maps api key before calling any other methods
func SetApiKey(key string) {
	gmapsApiClient.Key = key
}

// Convert a Point to a Place object using the Google API
func PointToPlace(p LatLnger) (*Place, error) {
	l := p.LatLng()
	places, err := gmapsApiClient.ReverseGeocode(l[0], l[1])
	if err != nil {
		return nil, err
	}
	gPlace := places[0]
	return gPlaceToPlace(gPlace), nil
}

// Convert a string place name/address to a Place object using the Google API
// While the API may return many possible place results this method simply picks the first one
func StringToPlace(s string) (*Place, error) {
	predictions, err := gmapsApiClient.Autocomplete(s)
	if err != nil {
		return nil, err
	}
	pid := predictions[0].PlaceId
	gPlace, err := gmapsApiClient.Details(pid)
	if err != nil {
		return nil, err
	}
	return gPlaceToPlace(*gPlace), nil
}

// Convert a string place name/address to a Point (using StringToPlace)
func StringToPoint(s string) (*Point, error) {
	p, err := StringToPlace(s)
	if err != nil {
		return nil, err
	}
	return p.Location, err
}

func gPlaceToPlace(gPlace gmaps.Place) *Place {
	place := Place{
		Name:          gPlace.Name,
		Location:      PointFromLatLng(gPlace.Geometry.Location),
		AddressString: gPlace.FormattedAddress,
	}
	place.AddressComponents = make([]AddressComponent, len(gPlace.AddressComponents))
	for i, c := range gPlace.AddressComponents {
		place.AddressComponents[i].Name = c.Name
		place.AddressComponents[i].Type = c.Types[0]
	}
	if gPlace.Geometry.Viewport != nil {
		northEast := gPlace.Geometry.Viewport.Northeast
		southWest := gPlace.Geometry.Viewport.Southwest
		place.BoundingBox = NewEnvelope(southWest.Lng, northEast.Lat, northEast.Lng, southWest.Lat)
	}
	return &place
}

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

type AddressComponent struct {
	Name string
	Type string
}
