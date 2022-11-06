package components

const yamlSeparator = "\n---"

// manifestLoader loads a specific manifest kind from a folder.
type DiskManifestLoader[T kubernetesManifest] struct {
	zv func() T
	kind string
	path string
}

