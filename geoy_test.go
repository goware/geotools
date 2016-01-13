package geoy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

var (
	p = &Point{Type: "point", Coordinates: []float64{-6.538, 53.339}}
	l = &LatLng{Lat: 53.339, Lng: -6.538}
)

const (
	apiKey = "AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI"
)

func init() {
	SetAPIKey(apiKey)
}

func TestNewPoint(t *testing.T) {
	q := NewPoint(-6.538, 53.339)
	assert.Equal(t, p.Coordinates[0], q.Coordinates[0])
	assert.Equal(t, p.Coordinates[1], q.Coordinates[1])
}

func TestLatLngFromPoint(t *testing.T) {
	x := LatLngFromPoint(*NewPoint(-6.538, 53.339))
	assert.Equal(t, *l, *x)
}

func TestLookupCoordinates(t *testing.T) {
	res, err := LookupCoordinates(context.Background(), *p)
	assert.NoError(t, err)
	assert.True(t, len(res) > 0)

	first := res[0]
	assert.Equal(t, "Eio0IE1haW4gU3QsIENlbGJyaWRnZSwgQ28uIEtpbGRhcmUsIElyZWxhbmQ", first.PlaceID)

	assert.Equal(t, "Main Street", first.Address.Street)
	assert.Equal(t, "4", first.Address.HouseNumber)
	assert.Equal(t, "Celbridge", first.Address.City)
	assert.Equal(t, "Kildare", first.Address.State)
	assert.Equal(t, "Ireland", first.Address.Country)
	assert.Equal(t, "4 Main St, Celbridge, Co. Kildare, Ireland", first.Address.Formatted)

	t.Logf("%v", res)
}

func TestLookupName(t *testing.T) {
	res, err := LookupName(context.Background(), "Liberty Village")
	assert.NoError(t, err)
	assert.True(t, len(res) > 0)

	first := res[0]
	assert.Equal(t, "Liberty, MO, USA", first.Address.Formatted)

	t.Logf("%v", res)
}

func TestLookupName2(t *testing.T) {
	res, err := LookupName(context.Background(), "Hipo, Istanbul")
	assert.NoError(t, err)
	assert.True(t, len(res) > 0)
	assert.Equal(t, "ChIJZ2wOIhK3yhQRyAefRIkiPfU", res[0].PlaceID)
	t.Logf("%v", res)
}

func TestStringToPoint(t *testing.T) {
	q, err := LookupName(context.Background(), "Liberty Village")
	assert.NoError(t, err)
	assert.True(t, len(q) > 0)
	point := q[0].Location
	assert.Equal(t, "SRID=4326;POINT(-94.419118 39.246114)", point.String())
	t.Logf("%v", point)
}

func TestPlaceDetails(t *testing.T) {
	place, err := PlaceDetails(context.Background(), "ChIJrTLr-GyuEmsRBfy61i59si0")
	assert.NoError(t, err)
	assert.True(t, place != nil)
	assert.Equal(t, "Sydney", place.Name)

	assert.Equal(t, "Sydney", place.Address.City)
	assert.Equal(t, "New South Wales", place.Address.State)
	assert.Equal(t, "Australia", place.Address.Country)
	assert.Equal(t, "32 The Promenade, King Street Wharf 5, Sydney NSW 2000, Australia", place.Address.Formatted)

	t.Logf("%v", place)
}

func TestInstagramToPlace(t *testing.T) {
	b := []byte(`{
        "latitude": 37.780885099999999,
        "id": "514276",
        "longitude": -122.3948632,
        "name": "Instagram"
    }`)
	var v InstagramLocation
	json.Unmarshal(b, &v)
	res, err := LookupCoordinates(context.Background(), v)
	assert.NoError(t, err)
	assert.True(t, len(res) > 0)

	first := res[0]
	assert.Equal(t, first.PlaceID, "Ei8xNzMtMTk5IFMgUGFyayBTdCwgU2FuIEZyYW5jaXNjbywgQ0EgOTQxMDcsIFVTQQ")

	assert.Equal(t, "South Park Street", first.Address.Street)
	assert.Equal(t, "173-199", first.Address.HouseNumber)
	assert.Equal(t, "San Francisco", first.Address.City)
	assert.Equal(t, "California", first.Address.State)
	assert.Equal(t, "United States", first.Address.Country)
	assert.Equal(t, "173-199 S Park St, San Francisco, CA 94107, USA", first.Address.Formatted)

	t.Logf("%v", res)
}
