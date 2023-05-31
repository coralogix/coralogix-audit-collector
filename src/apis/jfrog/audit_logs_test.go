package jfrog

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"testing"
)

func TestParseResponse(t *testing.T) {
	responses := []string{
		`2023-05-10T14:08:24.831Z|6e4a7f0c11e106e5|UNKNOWN|UNKNOWN|jffe@01f3csxdwrjzve1zd9chfy0zmm|jffe@01f3csxdwrjzve1zd9chfy0zmm|C|TKN|{"added":{"owner":"jffe@01f3csxdwrjzve1zd9chfy0zmm","created":"1683727704831","expirationTime":"1683728304831","subject":"jffe@01f3csxdwrjzve1zd9chfy0zmm","scope":"applied-permissions/admin","id":"81f00584-8eac-4ce1-a1d7-76b723a89d9c","type":"generic"}}`,
	}

	j := &Jfrog{}
	for _, response := range responses {
		fake, err := fakeData(response)
		if err != nil {
			t.Errorf("failed to create fake data: %s", err)
		}

		auditLog, err := j.parseResponse(fake)
		if err != nil {
			t.Errorf("failed to parse response: %s", err)
		}
		if auditLog == nil {
			t.Errorf("auditLog is nil")
		}
	}
}

func fakeData(data string) (io.Reader, error) {
	// Create a buffer to hold the compressed data
	var buf bytes.Buffer

	// Create a GZIP writer using the buffer
	gzipWriter := gzip.NewWriter(&buf)

	// Write dummy data to the GZIP writer
	dummyData := []byte(data)
	_, err := gzipWriter.Write(dummyData)
	if err != nil {
		return nil, fmt.Errorf("failed to write dummy data: %w", err)
	}

	// Close the GZIP writer
	err = gzipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close gzip writer: %w", err)
	}

	// Retrieve the compressed data from the buffer
	compressedData := buf.Bytes()
	return bytes.NewReader(compressedData), nil
}
