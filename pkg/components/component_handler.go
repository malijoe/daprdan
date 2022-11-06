package components

import componentsV1alpha1 "github.com/malijoe/daprdan/pkg/apis/components/v1alpha1"

// ComponentHandler is an interface for reacting on component changes.
type ComponentHandler interface {
	OnComponentUpdate(component componentsV1alpha1.Component)
}
