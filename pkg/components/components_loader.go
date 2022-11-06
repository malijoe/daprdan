package components

import componentsV1alpha1 "github.com/malijoe/daprdan/pkg/apis/components/v1alpha1"

// ComponentLoader is an interface for returning Daprd components.
type ComponentLoader interface {
	LoadComponents() ([]componentsV1alpha1.Component, error)
}
