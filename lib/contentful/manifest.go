package contentful

import (
	"encoding/json"
	"io"
)

// Manifest represents the top level JSON object for a manifest file.
type Manifest struct {
	Queries []Query `json:"queries"`
}

// ManifestParser is the interface that wraps the Parse method that returns a manifest.
type ManifestParser interface {
	Parse(r io.Reader) (*Manifest, error)
}

type manifestParser struct {
}

// NewManifestParser constructs a struct that satisfies the ManifestParser interface.
func NewManifestParser() ManifestParser {
	return &manifestParser{}
}

// Parse parses the input Reader and constructs a Manifest.
// Parse expects the reader to contain a JSON string.
func (mp *manifestParser) Parse(r io.Reader) (*Manifest, error) {
	decoder := json.NewDecoder(r)
	var manifest Manifest
	if err := decoder.Decode(&manifest); err != nil {
		return nil, err
	}
	return &manifest, nil
}
