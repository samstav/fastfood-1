package fastfood

import "testing"

func TestNewManifest(t *testing.T) {
	sampleFile := "tests/templatepack/manifest.json"

	manifest, err := NewManifest(sampleFile)

	if err != nil {
		t.Errorf("Did not expect error %v", err)
	}

	if len(manifest.Providers) != 2 {
		t.Errorf("Expected the length of the commands array to be 2")
	}

	dbCmdExists := false
	for _, provider := range manifest.Providers {
		if provider.Name == "database" {
			dbCmdExists = true
		}
	}

	if !dbCmdExists {
		t.Errorf("Expected one of the parsed commands to match the name 'db'")
	}

	if len(manifest.Cookbook.Files) == 0 {
		t.Errorf("Expected more than 0 cookbook files")
	}
}
