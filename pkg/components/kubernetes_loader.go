package components

import (
	"context"
	"encoding/json"
	"time"

	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	"github.com/malijoe/kit/logger"

	componentsV1alpha1 "github.com/malijoe/daprdan/pkg/apis/components/v1alpha1"
	config "github.com/malijoe/daprdan/pkg/config/modes"
	operatorv1pb "github.com/malijoe/daprdan/pkg/proto/operator/v1"
)

var log = logger.NewLogger("daprd.runtime.components")

const (
	operatorCallTimeout = time.Second * 5
	operatorMaxRetries  = 100
)

// KubernetesComponents loads components in a kubernetes environment.
type KubernetesComponents struct {
	config    config.KubernetesConfig
	client    operatorv1pb.OperatorClient
	namespace string
	podName   string
}

func NewKubernetesComponents(configuration config.KubernetesConfig, namespace string, operatorClient operatorv1pb.OperatorClient, podName string) *KubernetesComponents {
	return &KubernetesComponents{
		config:    configuration,
		client:    operatorClient,
		namespace: namespace,
		podName:   podName,
	}
}

// LoadComponents returns components from a given control plane address.
func (k *KubernetesComponents) LoadComponents() ([]componentsV1alpha1.Component, error) {
	resp, err := k.client.ListComponents(context.Background(), &operatorv1pb.ListComponentsRequest{
		Namespace: k.namespace,
		PodName:   k.podName,
	}, grpcRetry.WithMax(operatorMaxRetries), grpcRetry.WithPerRetryTimeout(operatorCallTimeout))
	if err != nil {
		return nil, err
	}
	comps := resp.GetComponents()

	components := []componentsV1alpha1.Component{}
	for _, c := range comps {
		var component componentsV1alpha1.Component
		component.Spec = componentsV1alpha1.ComponentSpec{}
		err := json.Unmarshal(c, &component)
		if err != nil {
			log.Warnf("error deserializing component: %s", err)
			continue
		}
		components = append(components, component)
	}
	return components, nil
}
