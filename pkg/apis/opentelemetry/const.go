package opentelemetry

type (
	// ContextEntry represents a key in the context map
	ContextEntry string
)

const (
	// ContextInstance is the OpenTelemetryService CR (instance) that is the current target of the reconciliation
	ContextInstance ContextEntry = "__instance"

	// ContextLogger represents the context entry for the logger instance to be used for context-dependent log entries
	ContextLogger ContextEntry = "__logger"

	// CollectorConfigMapEntry represents the configuration file name for the collector
	CollectorConfigMapEntry = "collector.yaml"
)