// +build mock

package gmaps

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"googlemaps.github.io/maps"
)

const defaultMockFile = `mocks.gob`

type mockData []interface{}

var errNoMockData = errors.New(`No mock data.`)

var (
	mockFile string
	mockMap  map[string]mockData
	mockMu   sync.RWMutex
)

func init() {
	if mockFile = os.Getenv("MOCK_FILE"); mockFile == "" {
		mockFile = defaultMockFile
	}

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
	}
}

func readMock(key string) ([]interface{}, error) {
	mockMu.RLock()
	defer mockMu.RUnlock()
	if v, ok := mockMap[key]; ok {
		log.Printf("Reading mock resource with key %v", key)
		return v, nil
	}
	log.Printf("No such key %v", key)
	return nil, errors.New(`No such key`)
}

func writeMock(key string, data ...interface{}) error {
	mockMu.Lock()
	defer mockMu.Unlock()
	mockMap[key] = data
	log.Printf("Wrote key %v", key)

	// Don't know how to make this happen just once at the end of the program. I
	// don't think this is too bad, this mock helper was written for internal use
	// only.
	writeMockFile()

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
