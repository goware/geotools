package geotools

import (
	"github.com/pressly/geotools/gmaps"
	"golang.org/x/net/context"
)

var (
	defaultMapsClient gmaps.MapsApiClient
)

func mapsClient() gmaps.MapsApiClient {
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

// PlaceDetails returns the details of a place given its placeID.
func PlaceDetails(ctx context.Context, placeID string) (*Place, error) {
	place, err := mapsClient().Details(ctx, placeID)
	if err != nil {
		return nil, err
	}
	res := result{
		PlaceID:           place.PlaceID,
		AddressComponents: place.AddressComponents,
		Geometry:          place.Geometry,
		FormattedAddress:  place.FormattedAddress,
	}
	return toPlace(&res), nil
}

// LookupCoordinates lookups a coordinate and returns all its associated
// places.
func LookupCoordinates(ctx context.Context, p LatLnger) ([]*Place, error) {
	l := p.LatLng()
	places, err := mapsClient().ReverseGeocode(ctx, l[0], l[1])
	if err != nil {
		return nil, err
	}

	res := make([]*Place, len(places))
	for i := range places {
		res[i] = toPlace(&result{
			PlaceID:           places[i].PlaceID,
			AddressComponents: places[i].AddressComponents,
			Geometry:          places[i].Geometry,
			FormattedAddress:  places[i].FormattedAddress,
		})
	}

	return res, nil
}

// LookupName converts a string place name/address into
func LookupName(ctx context.Context, s string) ([]*Place, error) {
	places, err := mapsClient().TextSearch(ctx, s)
	if err != nil {
		return nil, err
	}
	res := make([]*Place, len(places))
	for i := range places {
		res[i] = toPlace(&result{
			PlaceID: places[i].PlaceID,
			//AddressComponents: places[i].AddressComponents,
			Geometry:         places[i].Geometry,
			FormattedAddress: places[i].FormattedAddress,
		})
	}
	return res, nil
}
