package geoy

import (
	"github.com/pressly/geoy/gmaps"
	"golang.org/x/net/context"
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
func PlaceDetails(ctx context.Context, placeId string) (*Place, error) {
	place, err := mapsClient().Details(ctx, placeId)
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
func PointToPlace(ctx context.Context, p LatLnger) (*Place, error) {
	l := p.LatLng()
	places, err := mapsClient().ReverseGeocode(ctx, l[0], l[1])
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
func StringToPlace(ctx context.Context, s string) (*Place, error) {
	predictions, err := mapsClient().Autocomplete(ctx, s)
	if err != nil {
		return nil, err
	}
	placeID := predictions[0].PlaceID
	placeDetails, err := mapsClient().Details(ctx, placeID)
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
func StringToPoint(ctx context.Context, s string) (*Point, error) {
	p, err := StringToPlace(ctx, s)
	if err != nil {
		return nil, err
	}
	return p.Location, err
}
