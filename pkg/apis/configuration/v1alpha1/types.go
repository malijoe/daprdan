package v1alpha1

import (
	"strconv"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +kubebuilder:object:root=true

// Configuration describes a Daprd configuration setting.
type Configuration struct {
	metav1.TypeMeta `json:",inline"`
	//+optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec ConfigurationSpec `json:"spec,omitempty"`
}

// ConfigurationSpec is the spec for a configuration.
type ConfigurationSpec struct {
	// +optional
	AppHTTPPipelineSpec PipelineSpec `json:"appHttpPipeline,omitempty"`
	// +optional
	HTTPPipelineSpec PipelineSpec `json:"httpPipeline,omitempty"`
	// +optional
	TracingSpec TracingSpec `json:"tracing,omitempty"`
	// +kubebuilder:default={enabled:true}
	MetricSpec MetricSpec `json:"metric,omitempty"`
	// +optional
	MTLSSpec MTLSSpec `json:"mtls,omitempty"`
	// +optional
	Secrets SecretsSpec `json:"secret,omitempty"`
	// +optional
	AccessControlSpec AccessControlSpec `json:"accessControl,omitempty"`
	// +optional
	NameResolutionSpec NameResolutionSpec `json:"nameResolution,omitempty"`
	// +optinal
	Features []FeatureSpec `json:"features,omitempty"`
	// +optional
	APISpec APISpec `json:"api,omitempty"`
	// +optional
	ComponentsSpec ComponentsSpec `json:"components,omitempty"`
}

// APISpec describes the configuration for Daprd APIs
type APISpec struct {
	Allowed []APIAccessRule `json:"allowed,omitempty"`
}

// APIAccessRule describes an acccess rule for allowing a Daprd API to be enabled and accessible by an app.
type APIAccessRule struct {
	Component     string       `json:"component"`
	Version       string       `json:"version"`
	Configuration DynamicValue `json:"configuration"`
}

// NameResolutionSpec is the spec for name resolution configuration.
type NameResolutionSpec struct {
	Component     string       `json:"component"`
	Version       string       `json:"version"`
	Configuration DynamicValue `json:"configuration"`
}

// SecretsSpec is the spec for secrets configuration.
type SecretsSpec struct {
	Scopes []SecretsScope `json:"scopes"`
}

// SecretsScope defines the scope for secrets.
type SecretsScope struct {
	// +optional
	DefaultAccess string `json:"defaultAcces,omitempty"`
	StoreName     string `json:"storeName"`
	// +optional
	AllowedSecrets []string `json:"allowedSecrets,omitempty"`
	// +optional
	DeniedSecrets []string `json:"deniedSecrets,omitempty"`
}

// PipelineSpec defines the middleware pipeline.
type PipelineSpec struct {
	Handlers []HandlerSpec `json:"handlers"`
}

// HandlerSpec defines a request's handler.
type HandlerSpec struct {
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	SelectorSpec SelectorSpec `json:"selector,omitempty"`
}

// MTLSSpec defines mTLS configuration.
type MTLSSpec struct {
	Enabled bool `json:"enabled"`
	// +optional
	WorkloadCertTTL string `json:"workloadCertTTL"`
	// +optional
	AllowedClockSkew string `json:"allowedClockSkew"`
}

// SelectorSpec selects target services to which the handler is to be applied.
type SelectorSpec struct {
	Fields []SelectorField `json:"fields"`
}

// SelectorField defines a selctor's field.
type SelectorField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// TracingSpec defines distributed tracing configuration.
type TracingSpec struct {
	SamplingRate string `json:"samplingRate"`
	// +optional
	Stdout bool `json:"stdout"`
	// +optional
	Zipkin ZipkinSpec `json:"zipkin"`
	// +optional
	Otel OtelSpec `json:"otel"`
}

// OtelSpec defines Otel exporter configurations.
type OtelSpec struct {
	Protocol        string `json:"protocol" yaml:"protocol"`
	EndpointAddress string `json:"endpointAddress" yaml:"endpointAddress"`
	IsSecure        bool   `json:"isSecure" yaml:"isSecure"`
}

// ZipkinSpec defines Zipkin trace configuration.
type ZipkinSpec struct {
	EndpointAddress string `json:"endpointAddress"`
}

// MetricSpec defines metrics configuration.
type MetricSpec struct {
	Enabled bool `json:"enabled"`
}

// AppPolicySpec defines the policy data structure for each app.
type AppPolicySpec struct {
	AppName string `json:"appId" yaml:"appId"`
	// +optional
	DefaultAction string `json:"defaultAction" yaml:"defaultAction"`
	// +optional
	TrustDomain string `json:"trustDomain" yaml:"trustDomain"`
	// +optional
	Namespace string `json:"namespace" yaml:"namespace"`
	// +optional
	AppOperationActions []AppOperationAction `json:"operations" yaml:"operations"`
}

// AppOperationAction
type AppOperationAction struct {
	Operation string `json:"name" yaml:"name"`
	// +optional
	HTTPVerb []string `json:"httpVerb" yaml:"httpVerb"`
	Action   string   `json:"action" yaml:"action"`
}

// AccessControlSpec is the spec object in ConfigurationSpec.
type AccessControlSpec struct {
	// +optional
	DefaultAction string `json:"defaultAction" yaml:"defaultAction"`
	// +optional
	TrustDomain string `json:"trustDomain" yaml:"trustDomain"`
	// +optional
	AppPolicies []AppPolicySpec `json:"policies" yaml:"policies"`
}

// FeatureSpec defines the features that are enabled/disabled.
type FeatureSpec struct {
	Name    string `json:"name" yaml:"name"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}

// ComponentsSpec describes the configuration for Daprd components
type ComponentsSpec struct {
	// Denylist of component types that cannot be instantiated
	// +optional
	Deny []string `json:"deny,omitempty" yaml:"deny,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList is a list of Daprd event sources.
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Configuration `json:"items"`
}

// DynamicValue is a dynamic value struct for the component.metadata pair value.
type DynamicValue struct {
	v1.JSON `json:",inline"`
}

// String returns the string representation of the raw value.
// If the value is a string, it will be unquated as the string is guaranmteed to be a JSON serialized string.
func (d *DynamicValue) String() string {
	s := string(d.Raw)
	c, err := strconv.Unquote(s)
	if err == nil {
		s = c
	}
	return s
}
