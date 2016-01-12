// +build mock

package gmaps

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const defaultMockFile = `testdata.gob`

type mockData []interface{}

var errNoMockData = errors.New(`No mock data.`)

var (
	mockFile string
	mockMap  map[string]mockData
	mockMu   sync.RWMutex
)

func loadMockFile() {
	if mockFile = os.Getenv("MOCK_FILE"); mockFile == "" {
		mockFile = defaultMockFile
	}

	mockMap = make(map[string]mockData)

	fp, err := os.Open(mockFile)
	if err == nil {
		dec := gob.NewDecoder(fp)
		err := dec.Decode(&mockMap)
		if err != nil && err != io.EOF {
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
	if err := writeMockFile(); err != nil {
		log.Printf("Error writing mock file: %q", err)
	}

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
