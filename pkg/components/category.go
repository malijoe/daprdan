package components

type Category string

// supported components category
const (
	CategoryBindings       Category = "bindings"
	CategoryPubSub         Category = "pubsub"
	CategorySecretStore    Category = "secretstores"
	CategoryStateStore     Category = "state"
	CategoryMiddleware     Category = "middleware"
	CategoryConfiguration  Category = "configuration"
	CategoryLock           Category = "lock"
	CategoryNameResolution Category = "nameresolution"
)
