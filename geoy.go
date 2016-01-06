package geoy

import (
	"github.com/pressly/geoy/gmaps"
)

var (
	defaultMapsClient *gmaps.MapsApiClient
)

func mapsClient() *gmaps.MapsApiClient {
	if defaultMapsClient == nil {
		panic("Maps client was not initialized. Missing call to SetAPIKey()?")
	}
	return defaultMapsClient
}

// SetAPIKey sets the Google Maps API key.
func SetAPIKey(key string) (err error) {
	defaultMapsClient, err = gmaps.NewMapsClient(key)
	return err
}

// PlaceDetails returns the details of a place given its placeId.
func PlaceDetails(placeId string) (*Place, error) {
	place, err := mapsClient().Details(placeId)
	if err != nil {
		return nil, err
	}
	res := result{
		PlaceID:           place.PlaceID,
		AddressComponents: place.AddressComponents,
		Geometry:          place.Geometry,
		FormattedAddress:  place.FormattedAddress,
	}
	return res.toPlace(), nil
}

// PointToPlace lookups a coordinate and returns the place that corresponds to it.
func PointToPlace(p LatLnger) (*Place, error) {
	l := p.LatLng()
	places, err := mapsClient().ReverseGeocode(l[0], l[1])
	if err != nil {
		return nil, err
	}
	res := result{
		PlaceID:           places[0].PlaceID,
		AddressComponents: places[0].AddressComponents,
		Geometry:          places[0].Geometry,
		FormattedAddress:  places[0].FormattedAddress,
	}
	return res.toPlace(), nil
}

// StringToPlace converts a string place name/address to a Place object. While
// the API may return many possible place results this method simply picks the
// first one
func StringToPlace(s string) (*Place, error) {
	predictions, err := mapsClient().Autocomplete(s)
	if err != nil {
		return nil, err
	}
	placeID := predictions[0].PlaceID
	placeDetails, err := mapsClient().Details(placeID)
	if err != nil {
		return nil, err
	}
	res := result{
		PlaceID:           placeDetails.PlaceID,
		AddressComponents: placeDetails.AddressComponents,
		Geometry:          placeDetails.Geometry,
		FormattedAddress:  placeDetails.FormattedAddress,
	}
	return res.toPlace(), nil
}

// StringToPoint converts a string place name/address to a Point (using
// StringToPlace)
func StringToPoint(s string) (*Point, error) {
	p, err := StringToPlace(s)
	if err != nil {
		return nil, err
	}
	return p.Location, err
}
