package components

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"

	"github.com/malijoe/daprdan/utils"

	"github.com/ghodss/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const yamlSeparator = "\n---"

// manifestLoader loads a specific manifest kind from a folder.
type DiskManifestLoader[T kubernetesManifest] struct {
	zv   func() T
	kind string
	path string
}

func NewDiskManifestLoader[T kubernetesManifest](path string, zeroValue func() T) DiskManifestLoader[T] {
	return DiskManifestLoader[T]{
		path: path,
		kind: zeroValue().Kind(),
		zv:   zeroValue,
	}
}

func (m DiskManifestLoader[T]) Load() ([]T, error) {
	files, err := os.ReadDir(m.path)
	if err != nil {
		return nil, err
	}

	manifests := make([]T, 0)

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			if !utils.IsYaml(fileName) {
				log.Warnf("A non-YAML %s file %s was detected, it will not be loaded", m.kind, fileName)
				continue
			}
			fileManifests := m.loadManifestsFromFile(filepath.Join(m.path, fileName))
			manifests = append(manifests, fileManifests...)
		}
	}

	return manifests, nil
}

func (m DiskManifestLoader[T]) loadManifestsFromFile(manifestPath string) []T {
	var errors []error

	manifests := make([]T, 0)
	b, err := os.ReadFile(manifestPath)
	if err != nil {
		log.Warnf("daprd load %s error when reading file %s: %s", m.kind, manifestPath, err)
		return manifests
	}
	manifests, errors = m.decodeYaml(b)
	for _, err := range errors {
		log.Warnf("daprd load %s error when parsing manifests yaml resource in %s : %s", m.kind, manifestPath, err)
	}
	return manifests
}

type typeInfo struct {
	metav1.TypeMeta `json:",inline"`
}

func (m DiskManifestLoader[T]) decodeYaml(b []byte) ([]T, []error) {
	list := make([]T, 0)
	errors := []error{}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Split(splitYamlDoc)

	for {
		if !scanner.Scan() {
			err := scanner.Err()
			if err != nil {
				errors = append(errors, err)
				continue
			}
			break
		}

		scannerBytes := scanner.Bytes()
		var ti typeInfo
		if err := yaml.Unmarshal(scannerBytes, &ti); err != nil {
			errors = append(errors, err)
			continue
		}

		if ti.Kind != m.kind {
			continue
		}

		manifest := m.zv()
		if err := yaml.Unmarshal(scannerBytes, &manifest); err != nil {
			errors = append(errors, err)
			continue
		}

		list = append(list, manifest)
	}

	return list, errors
}

func splitYamlDoc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	sep := len([]byte(yamlSeparator))
	if i := bytes.Index(data, []byte(yamlSeparator)); i >= 0 {
		i += sep
		after := data[i:]

		if len(after) == 0 {
			if atEOF {
				return len(data), data[:len(data)-sep], nil
			}
			return 0, nil, nil
		}
		if j := bytes.IndexByte(after, '\n'); j >= 0 {
			return i + j + 1, data[0 : i-sep], nil
		}
		return 0, nil, nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
