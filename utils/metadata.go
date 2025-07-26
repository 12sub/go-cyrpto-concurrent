package utils

import (
	"fmt"
	"os"
	"time"
	"gopkg.in/yaml.v3"
)

type Metadata struct {
	OriginalFilename	string	`yaml:"original_filename"`
	Scheme	string		`yaml:"scheme"`
	KeyDerivation	string	`yaml:"key_derivation"`
	Salt		string		`yaml:"salt,omitempty"`
	Timestamp 	time.Time	`yaml:"timestamp"`	
}

func WriteMetadataFile(path string, meta Metadata) error {
	yamlBytes, err := yaml.Marshal(&meta)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %v", err)
	}
	metaPath := path + ".meta.yaml"
	return os.WriteFile(metaPath, yamlBytes, 0644)
}

// Read Metadata file
func LoadMetadataFile(path string) (Metadata, error) {
	var meta Metadata
	data, err := os.ReadFile(path + ".meta.yaml")
	if err != nil {
		return meta, err
	}
	err = yaml.Unmarshal(data, &meta)
	return meta, err
}