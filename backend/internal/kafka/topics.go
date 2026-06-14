package kafka

const (
	TopicMetricsRaw        = "marsadyn.metrics.raw"
	TopicLogsRaw           = "marsadyn.logs.raw"
	TopicTracesRaw         = "marsadyn.traces.raw"
	TopicMetricsValidated  = "marsadyn.metrics.validated"
	TopicLogsValidated     = "marsadyn.logs.validated"
	TopicTracesValidated   = "marsadyn.traces.validated"
	TopicDeadLetter        = "marsadyn.deadletter"
)

type TopicConfig struct {
	Name              string
	Partitions        int
	ReplicationFactor int
}

var DefaultTopics = map[string]TopicConfig{
	TopicMetricsRaw: {
		Name:              TopicMetricsRaw,
		Partitions:        3,
		ReplicationFactor: 1,
	},
	TopicLogsRaw: {
		Name:              TopicLogsRaw,
		Partitions:        3,
		ReplicationFactor: 1,
	},
	TopicTracesRaw: {
		Name:              TopicTracesRaw,
		Partitions:        3,
		ReplicationFactor: 1,
	},
	TopicMetricsValidated: {
		Name:              TopicMetricsValidated,
		Partitions:        3,
		ReplicationFactor: 1,
	},
	TopicLogsValidated: {
		Name:              TopicLogsValidated,
		Partitions:        3,
		ReplicationFactor: 1,
	},
	TopicTracesValidated: {
		Name:              TopicTracesValidated,
		Partitions:        3,
		ReplicationFactor: 1,
	},
	TopicDeadLetter: {
		Name:              TopicDeadLetter,
		Partitions:        3,
		ReplicationFactor: 1,
	},
}

func GetTopicConfig(topicName string) (TopicConfig, bool) {
	config, ok := DefaultTopics[topicName]
	return config, ok
}
