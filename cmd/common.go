package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	tempPackEnvVar     = "FASTFOOD_TEMPLATE_PACK"
	cookbookPathEnvVar = "COOKBOOKS"
	cookbookTemplates  = "cookbook"
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

func (m *Manifest) Help() string {
	var providersHelp []string

	for name, provider := range m.Providers {
		var help string
		if provider.Help == "" {
			help = "NO HELP FOUND"
		} else {
			help = provider.Help
		}

		providersHelp = append(
			providersHelp,
			fmt.Sprintf("  %-15s - %s", name, help),
		)
	}

	return fmt.Sprintf(`
Available Providers:

%s
`, strings.Join(providersHelp, "\n\n"))
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

func DefaultCookbookPath() string {
	cookbookPath := os.Getenv(cookbookPathEnvVar)

	return cookbookPath
}

// Translates key:value strings into a map
func MapArgs(args []string) map[string]string {
	var argMap map[string]string
	argMap = make(map[string]string)

	for _, arg := range args {
		if strings.Contains(arg, ":") {
			// Split at the first : in an arg
			splitArg := strings.SplitN(arg, ":", 2)

			argMap[splitArg[0]] = splitArg[1]
		}
	}

	return argMap
}
