package gmaps

import (
	"encoding/gob"
	"errors"
	"fmt"
	"googlemaps.github.io/maps"
	"log"
	"os"
	"strings"
	"sync"
)

const mockFile = `mocks.gob`

type mockData []interface{}

var (
	mockMap map[string]mockData
	mockMu  sync.RWMutex
)

func init() {
	mockMap = make(map[string]mockData)

	gob.Register(&maps.PlaceDetailsResult{})
	gob.Register(&maps.QueryAutocompleteResponse{})
	gob.Register([]maps.GeocodingResult{})
	gob.Register(&maps.QueryAutocompletePrediction{})
	gob.Register([]maps.QueryAutocompletePrediction{})

	fp, err := os.Open(mockFile)
	if err == nil {
		dec := gob.NewDecoder(fp)
		err := dec.Decode(&mockMap)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("maps: %#v", mockMap)
	}
}

func readMock(key string) ([]interface{}, error) {
	mockMu.RLock()
	defer mockMu.RUnlock()
	log.Printf("Reading mock resource with key %v", key)
	if v, ok := mockMap[key]; ok {
		return v, nil
	}
	return nil, errors.New(`No such key`)
}

func writeMock(key string, data ...interface{}) error {
	mockMu.Lock()
	defer mockMu.Unlock()
	mockMap[key] = data
	return nil
}

func writeMockFile() error {
	fp, err := os.Create(mockFile)
	if err != nil {
		return err
	}
	defer fp.Close()

	enc := gob.NewEncoder(fp)
	return enc.Encode(mockMap)
}

func mockKey(fn string, args ...interface{}) string {
	chunks := make([]string, len(args)+1)
	chunks[0] = fn
	for i := range args {
		chunks[i+1] = fmt.Sprintf("%v", args[i])
	}
	return strings.Join(chunks, "/")
}
