package components

import (
	componentsV1alpha1 "github.com/malijoe/daprdan/pkg/apis/components/v1alpha1"
	config "github.com/malijoe/daprdan/pkg/config/modes"
)

// StandaloneComponents loads components in a standalone environment.
type StandaloneComponents struct {
	componentsManifestLoader ManifestLoader[componentsV1alpha1.Component]
}

// newComponent creates a new component manifest.
func newComponent() componentsV1alpha1.Component {
	var comp componentsV1alpha1.Component
	comp.Spec = componentsV1alpha1.ComponentSpec{}
	return comp
}

// NewStandaloneComponents returns a new standalone loader.
func NewStandaloneComponents(configuration config.StandaloneConfig) *StandaloneComponents {
	return &StandaloneComponents{
		componentsManifestLoader: NewDiskManifestLoader(configuration.ComponentsPath, newComponent),
	}
}

// LoadComponents loads daprd components from a given directory.
func (s *StandaloneComponents) LoadComponents() ([]componentsV1alpha1.Component, error) {
	return s.componentsManifestLoader.Load()
}
