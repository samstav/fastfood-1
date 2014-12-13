package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	tempPackEnvVar    = "FASTFOOD_TEMPLATE_PACK"
	cookbookTemplates = "cookbook"
)

type Manifest struct {
	Providers map[string]struct {
		Name          string `json:"name"`
		Manifest      string `json:"manifest"`
		Help          string `json:"help"`
		templatesPath string
	}

	Cookbook struct {
		Directories   []string `json:"directories"`
		Files         []string `json:"files"`
		TemplatesPath string   `json:"templates_path"`
	}
}

func NewManifest(path string) (Manifest, error) {

	var manifest Manifest

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return manifest, errors.New(
			fmt.Sprintf("reading manifest %s: %v", path, err),
		)
	}

	err = json.Unmarshal(f, &manifest)
	if err != nil {
		return manifest, errors.New(
			fmt.Sprintf("parsing manifest %s: %v", path, err),
		)
	}

	if manifest.Cookbook.TemplatesPath == "" {
		manifest.Cookbook.TemplatesPath = cookbookTemplates
	}

	return manifest, nil

}

func DefaultTempPack() string {
	packEnv := os.Getenv(tempPackEnvVar)
	if packEnv == "" {
		return path.Join(os.Getenv("HOME"), "fastfood")
	} else {
		return packEnv
	}
}
