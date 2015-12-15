package gmaps

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type autocompleteResponse struct {
	Status      string
	Predictions []prediction
}

type prediction struct {
	Description string
	PlaceId     string `json:"place_id"`
}

type detailsResponse struct {
	Status string
	Result Place
}

type geocodeResponse struct {
	Status  string
	Results []Place
}

type Place struct {
	FormattedAddress  string `json:"formatted_address"`
	Geometry          *geometry
	Name              string
	AddressComponents []struct {
		Name  string   `json:"long_name"`
		Types []string `json:"types"`
	} `json:"address_components"`
}

type geometry struct {
	Location latlng
	Viewport *viewport
}

type viewport struct {
	Northeast latlng
	Southwest latlng
}

type latlng struct {
	Lat float64
	Lng float64
}

func (l latlng) LatLon() []float64 {
	return []float64{l.Lat, l.Lng}
}

func (l *latlng) String() string {
	return fmt.Sprintf("%0.6f,%0.6f", l.Lat, l.Lng)
}

type MapsApiClient struct {
	Key string
}

func (c *MapsApiClient) Autocomplete(input string) ([]prediction, error) {
	q := url.Values{}
	q.Set("input", input)
	data := autocompleteResponse{}
	err := c.callMethod("place/autocomplete", q, &data)
	if err != nil {
		return nil, err
	}
	if data.Status != "OK" {
		return nil, errors.New(data.Status)
	}
	return data.Predictions, nil
}

func (c *MapsApiClient) Details(placeid string) (*Place, error) {
	q := url.Values{}
	q.Set("placeid", placeid)
	data := detailsResponse{}
	err := c.callMethod("place/details", q, &data)
	if err != nil {
		return nil, err
	}
	if data.Status != "OK" {
		return nil, errors.New(data.Status)
	}
	return &data.Result, nil
}

func (c *MapsApiClient) ReverseGeocode(lat, lng float64) ([]Place, error) {
	q := url.Values{}
	ll := &latlng{lat, lng}
	q.Set("latlng", ll.String())
	data := geocodeResponse{}
	err := c.callMethod("geocode", q, &data)
	if err != nil {
		return nil, err
	}
	if data.Status != "OK" {
		return nil, errors.New(data.Status)
	}
	return data.Results, nil
}

func (c *MapsApiClient) callMethod(method string, params url.Values, resultContainer interface{}) error {
	if c.Key != "" {
		params.Set("key", c.Key)
	} else {
		panic(errors.New("Maps api key must be set with set."))
	}

	u := "https://maps.googleapis.com/maps/api/" + method + "/json?"
	res, err := http.Get(u + params.Encode())
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, resultContainer)
	if err != nil {
		log.Printf("Unexpected body: %s", body)
	}
	return err
}
