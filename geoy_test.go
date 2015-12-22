package geoy

import (
	"encoding/json"
	"testing"
)

var (
	p *Point
	l *LatLng
)

func init() {
	p = &Point{Type: "point", Coordinates: []float64{-6.538, 53.339}}
	l = &LatLng{Lat: 53.339, Lng: -6.538}
	SetAPIKey("AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI")
}

func TestNewPoint(t *testing.T) {
	x := NewPoint(-6.538, 53.339)
	if x.Coordinates[0] != p.Coordinates[0] || x.Coordinates[1] != p.Coordinates[1] {
		t.Fail()
	}
}

func TestLatLngFromPoint(t *testing.T) {
	x := LatLngFromPoint(*NewPoint(-6.538, 53.339))
	if *x != *l {
		t.Fail()
	}
}

func TestPointToPlace(t *testing.T) {
	place, err := PointToPlace(*p)
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Logf("%v", place)
}

func TestStringToPlace(t *testing.T) {
	place, err := StringToPlace("Liberty Village")
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Logf("%v", place)
}

func TestStringToPlace2(t *testing.T) {
	place, err := StringToPlace("Hipo, Istanbul")
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Logf("%v", place)
}

func TestStringToPoint(t *testing.T) {
	x, err := StringToPoint("Liberty Village")
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Logf("%v", x)
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
	place, err := PointToPlace(v)
	if err != nil {
		t.Errorf("%s", err)
	}
	t.Logf("%v", place)
}
