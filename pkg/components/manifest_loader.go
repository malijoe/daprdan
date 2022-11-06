package components

type kubernetesManifest interface {
	Kind() string
}

// ManifestLoader loads manifest-like files.
type ManifestLoader[T kubernetesManifest] interface {
	Load() ([]T, error)
}
