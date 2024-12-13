package agent

import (
	"time"

	"github.com/sirupsen/logrus"
)

var clog = logrus.WithField("component", "config")

const (
	ListenPoll       = "poll"
	ListenWatch      = "watch"
	DeduperNone      = "none"
	DeduperFirstCome = "firstCome"
	DirectionIngress = "ingress"
	DirectionEgress  = "egress"
	DirectionBoth    = "both"

	IPTypeAny  = "any"
	IPTypeIPV4 = "ipv4"
	IPTypeIPV6 = "ipv6"

	IPIfaceExternal    = "external"
	IPIfaceLocal       = "local"
	IPIfaceNamedPrefix = "name:"
)

type FlowFilter struct {
	// FilterDirection is the direction of the flow filter.
	// Possible values are "Ingress" or "Egress".
	FilterDirection string `json:"direction,omitempty"`
	// FilterIPCIDR is the IP CIDR to filter flows.
	// Example: 10.10.10.0/24 or 100:100:100:100::/64, default is 0.0.0.0/0
	FilterIPCIDR string `json:"ip_cidr,omitempty"`
	// FilterProtocol is the protocol to filter flows.
	// supported protocols: TCP, UDP, SCTP, ICMP, ICMPv6
	FilterProtocol string `json:"protocol,omitempty"`
	// FilterSourcePort is the source port to filter flows.
	FilterSourcePort int32 `json:"source_port,omitempty"`
	// FilterDestinationPort is the destination port to filter flows.
	FilterDestinationPort int32 `json:"destination_port,omitempty"`
	// FilterPort is the port to filter flows, can be use for either source or destination port.
	FilterPort int32 `json:"port,omitempty"`
	// FilterSourcePortRange is the source port range to filter flows.
	// Example: 8000-8010
	FilterSourcePortRange string `json:"source_port_range,omitempty"`
	// FilterSourcePorts is two source ports to filter flows.
	// Example: 8000,8010
	FilterSourcePorts string `json:"source_ports,omitempty"`
	// FilterDestinationPortRange is the destination port range to filter flows.
	// Example: 8000-8010
	FilterDestinationPortRange string `json:"destination_port_range,omitempty"`
	// FilterDestinationPorts is two destination ports to filter flows.
	// Example: 8000,8010
	FilterDestinationPorts string `json:"destination_ports,omitempty"`
	// FilterPortRange is the port range to filter flows, can be used for either source or destination port.
	// Example: 8000-8010
	FilterPortRange string `json:"port_range,omitempty"`
	// FilterPorts is two ports option to filter flows, can be used for either source or destination port.
	// Example: 8000,8010
	FilterPorts string `json:"ports,omitempty"`
	// FilterICMPType is the ICMP type to filter flows.
	FilterICMPType int `json:"icmp_type,omitempty"`
	// FilterICMPCode is the ICMP code to filter flows.
	FilterICMPCode int `json:"icmp_code,omitempty"`
	// FilterPeerIP is the IP to filter flows.
	// Example: 10.10.10.10
	FilterPeerIP string `json:"peer_ip,omitempty"`
	// FilterAction is the action to filter flows.
	// Possible values are "Accept" or "Reject".
	FilterAction string `json:"action,omitempty"`
	// FilterTCPFlags is the TCP flags to filter flows.
	// possible values are: SYN, SYN-ACK, ACK, FIN, RST, PSH, URG, ECE, CWR, FIN-ACK, RST-ACK
	FilterTCPFlags string `json:"tcp_flags,omitempty"`
	// FilterDrops allow filtering flows with packet drops, default is false.
	FilterDrops bool `json:"drops,omitempty"`
	// FilterSample is the sample rate this matching flow will use
	FilterSample uint32 `json:"sample,omitempty"`
}

type Config struct {
	// AgentIP allows overriding the reported Agent IP address on each flow.
	AgentIP string `env:"AGENT_IP"`
	// AgentIPIface specifies which interface should the agent pick the IP address from in order to
	// report it in the AgentIP field on each flow. Accepted values are: external (default), local,
	// or name:<interface name> (e.g. name:eth0).
	// If the AgentIP configuration property is set, this property has no effect.
	AgentIPIface string `env:"AGENT_IP_IFACE" envDefault:"external"`
	// AgentIPType specifies which type of IP address (IPv4 or IPv6 or any) should the agent report
	// in the AgentID field of each flow. Accepted values are: any (default), ipv4, ipv6.
	// If the AgentIP configuration property is set, this property has no effect.
	AgentIPType string `env:"AGENT_IP_TYPE" envDefault:"any"`
	// Export selects the exporter protocol.
	// Accepted values for Flows are: grpc (default), kafka, ipfix+udp, ipfix+tcp or direct-flp.
	// Accepted values for Packets are: grpc (default) or direct-flp
	Export string `env:"EXPORT" envDefault:"grpc"`
	// Host is the host name or IP of the flow or packet collector, when the EXPORT variable is
	// set to "grpc"
	TargetHost string `env:"TARGET_HOST"`
	// Port is the port the flow or packet collector, when the EXPORT variable is set to "grpc"
	TargetPort int `env:"TARGET_PORT"`
	// GRPCMessageMaxFlows specifies the limit, in number of flows, of each GRPC message. Messages
	// larger than that number will be split and submitted sequentially.
	GRPCMessageMaxFlows int `env:"GRPC_MESSAGE_MAX_FLOWS" envDefault:"10000"`
	// Interfaces contains the interface names from where flows will be collected. If empty, the agent
	// will fetch all the interfaces in the system, excepting the ones listed in ExcludeInterfaces.
	// If an entry is enclosed by slashes (e.g. `/br-/`), it will match as regular expression,
	// otherwise it will be matched as a case-sensitive string.
	Interfaces []string `env:"INTERFACES" envSeparator:","`
	// ExcludeInterfaces contains the interface names that will be excluded from flow tracing. Default:
	// "lo" (loopback).
	// If an entry is enclosed by slashes (e.g. `/br-/`), it will match as regular expression,
	// otherwise it will be matched as a case-sensitive string.
	ExcludeInterfaces []string `env:"EXCLUDE_INTERFACES" envSeparator:"," envDefault:"lo"`
	// BuffersLength establishes the length of communication channels between the different processing
	// stages
	BuffersLength int `env:"BUFFERS_LENGTH" envDefault:"50"`
	// InterfaceIPs is a list of CIDR-notation IPs/Subnets where any interface containing an IP in the given ranges
	// should be listened on. This allows users to specify interfaces without knowing the OS-assigned interface names.
	// Exclusive with Interfaces/ExcludeInterfaces.
	InterfaceIPs []string `env:"INTERFACE_IPS" envSeparator:","`
	// ExporterBufferLength establishes the length of the buffer of flow batches (not individual flows)
	// that can be accumulated before the Kafka or GRPC exporter. When this buffer is full (e.g.
	// because the Kafka or GRPC endpoint is slow), incoming flow batches will be dropped. If unset,
	// its value is the same as the BUFFERS_LENGTH property.
	ExporterBufferLength int `env:"EXPORTER_BUFFER_LENGTH"`
	// CacheMaxFlows specifies how many flows can be accumulated in the accounting cache before
	// being flushed for its later export
	CacheMaxFlows int `env:"CACHE_MAX_FLOWS" envDefault:"5000"`
	// CacheActiveTimeout specifies the maximum duration that flows are kept in the accounting
	// cache before being flushed for its later export
	CacheActiveTimeout time.Duration `env:"CACHE_ACTIVE_TIMEOUT" envDefault:"5s"`
	// Deduper specifies the deduper type. Accepted values are "none" (disabled) and "firstCome".
	// When enabled, it will detect duplicate flows (flows that have been detected e.g. through
	// both the physical and a virtual interface).
	// "firstCome" will forward only flows from the first interface the flows are received from.
	Deduper string `env:"DEDUPER" envDefault:"none"`
	// DeduperFCExpiry specifies the expiry duration of the flows "firstCome" deduplicator. After
	// a flow hasn't been received for that expiry time, the deduplicator forgets it. That means
	// that a flow from a connection that has been inactive during that period could be forwarded
	// again from a different interface.
	// If the value is not set, it will default to 2 * CacheActiveTimeout
	DeduperFCExpiry time.Duration `env:"DEDUPER_FC_EXPIRY"`
	// DeduperJustMark will just mark duplicates (boolean field) instead of dropping them.
	DeduperJustMark bool `env:"DEDUPER_JUST_MARK" envDefault:"false"`
	// DeduperMerge will merge duplicated flows and generate list of interfaces and direction pairs
	DeduperMerge bool `env:"DEDUPER_MERGE" envDefault:"true"`
	// Direction allows selecting which flows to trace according to its direction. Accepted values
	// are "ingress", "egress" or "both" (default).
	Direction string `env:"DIRECTION" envDefault:"both"`
	// Logger level. From more to less verbose: trace, debug, info, warn, error, fatal, panic.
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	// Sampling holds the rate at which packets should be sampled and sent to the target collector.
	// E.g. if set to 100, one out of 100 packets, on average, will be sent to the target collector.
	Sampling int `env:"SAMPLING" envDefault:"0"`
	// ListenInterfaces specifies the mechanism used by the agent to listen for added or removed
	// network interfaces. Accepted values are "watch" (default) or "poll".
	// If the value is "watch", interfaces are traced immediately after they are created. This is
	// the recommended setting for most configurations. "poll" value is a fallback mechanism that
	// periodically queries the current network interfaces (frequency specified by ListenPollPeriod).
	ListenInterfaces string `env:"LISTEN_INTERFACES" envDefault:"watch"`
	// ListenPollPeriod specifies the periodicity to query the network interfaces when the
	// ListenInterfaces value is set to "poll".
	ListenPollPeriod time.Duration `env:"LISTEN_POLL_PERIOD" envDefault:"10s"`
	// KafkaBrokers is a comma-separated list of tha addresses of the brokers of the Kafka cluster
	// that this agent is configured to send messages to.
	KafkaBrokers []string `env:"KAFKA_BROKERS" envSeparator:","`
	// KafkaTopic is the name of the topic where the flows' processor will receive the flows from.
	KafkaTopic string `env:"KAFKA_TOPIC" envDefault:"network-flows"`
	// KafkaBatchMessages sets the limit on how many messages will be buffered before being sent to a
	// partition.
	KafkaBatchMessages int `env:"KAFKA_BATCH_MESSAGES" envDefault:"1000"`
	// KafkaBatchSize sets the limit, in bytes, of the maximum size of a request before being sent
	// to a partition.
	KafkaBatchSize int `env:"KAFKA_BATCH_SIZE" envDefault:"1048576"`
	// KafkaAsync. If it's true, the message writing process will never block. It also means that
	// errors are ignored since the caller will not receive the returned value.
	KafkaAsync bool `env:"KAFKA_ASYNC" envDefault:"true"`
	// KafkaCompression sets the compression codec to be used to compress messages. The accepted
	// values are: none (default), gzip, snappy, lz4, zstd.
	KafkaCompression string `env:"KAFKA_COMPRESSION" envDefault:"none"`
	// KafkaEnableTLS set true to enable TLS
	KafkaEnableTLS bool `env:"KAFKA_ENABLE_TLS" envDefault:"false"`
	// KafkaTLSInsecureSkipVerify skips server certificate verification in TLS connections
	KafkaTLSInsecureSkipVerify bool `env:"KAFKA_TLS_INSECURE_SKIP_VERIFY" envDefault:"false"`
	// KafkaTLSCACertPath is the path to the Kafka server certificate for TLS connections
	KafkaTLSCACertPath string `env:"KAFKA_TLS_CA_CERT_PATH"`
	// KafkaTLSUserCertPath is the path to the user (client) certificate for mTLS connections
	KafkaTLSUserCertPath string `env:"KAFKA_TLS_USER_CERT_PATH"`
	// KafkaTLSUserKeyPath is the path to the user (client) private key for mTLS connections
	KafkaTLSUserKeyPath string `env:"KAFKA_TLS_USER_KEY_PATH"`
	// KafkaEnableSASL set true to enable SASL auth
	KafkaEnableSASL bool `env:"KAFKA_ENABLE_SASL" envDefault:"false"`
	// KafkaSASLType type of SASL mechanism: plain or scramSHA512
	KafkaSASLType string `env:"KAFKA_SASL_TYPE" envDefault:"plain"`
	// KafkaSASLClientIDPath is the path to the client ID (username) for SASL auth
	KafkaSASLClientIDPath string `env:"KAFKA_SASL_CLIENT_ID_PATH"`
	// KafkaSASLClientSecretPath is the path to the client secret (password) for SASL auth
	KafkaSASLClientSecretPath string `env:"KAFKA_SASL_CLIENT_SECRET_PATH"`
	// ProfilePort sets the listening port for Go's Pprof tool. If it is not set, profile is disabled
	ProfilePort int `env:"PROFILE_PORT"`
	// Flowlogs-pipeline configuration as YAML or JSON, used when export is "direct-flp". Cf https://github.com/netobserv/flowlogs-pipeline
	// The "ingest" stage must be omitted from this configuration, since it is handled internally by the agent. The first stage should follow "preset-ingester".
	// E.g: {"pipeline":[{"name": "writer","follows": "preset-ingester"}],"parameters":[{"name": "writer","write": {"type": "stdout"}}]}.
	FLPConfig string `env:"FLP_CONFIG"`
	// Enable RTT calculations for the flows, default is false (disabled), set to true to enable.
	// This feature requires the flows agent to attach at both Ingress and Egress hookpoints.
	// If both Ingress and Egress are not enabled then this feature will not be enabled even if set to true via env.
	EnableRTT bool `env:"ENABLE_RTT" envDefault:"false"`
	// ForceGC enables forcing golang garbage collection run at the end of every map eviction, default is true
	ForceGC bool `env:"FORCE_GARBAGE_COLLECTION" envDefault:"true"`
	// EnablePktDrops enable Packet drops eBPF hook to account for dropped flows
	EnablePktDrops bool `env:"ENABLE_PKT_DROPS" envDefault:"false"`
	// EnableDNSTracking enable DNS tracking eBPF hook to track dns query/response flows
	EnableDNSTracking bool `env:"ENABLE_DNS_TRACKING" envDefault:"false"`
	// DNSTrackingPort used to define which port the DNS service is mapped to at the pod level,
	// so we can track DNS at the pod level
	DNSTrackingPort uint16 `env:"DNS_TRACKING_PORT" envDefault:"53"`
	// StaleEntriesEvictTimeout specifies the maximum duration that stale entries are kept
	// before being deleted, default is 5 seconds.
	StaleEntriesEvictTimeout time.Duration `env:"STALE_ENTRIES_EVICT_TIMEOUT" envDefault:"5s"`
	// EnablePCA enables Packet Capture Agent (PCA). By default, PCA is off.
	EnablePCA bool `env:"ENABLE_PCA" envDefault:"false"`
	// MetricsEnable enables http server to collect ebpf agent metrics, default is false.
	MetricsEnable bool `env:"METRICS_ENABLE" envDefault:"false"`
	// MetricsServerAddress is the address of the server that collects ebpf agent metrics.
	MetricsServerAddress string `env:"METRICS_SERVER_ADDRESS"`
	// MetricsPort is the port of the server that collects ebpf agent metrics.
	MetricsPort int `env:"METRICS_SERVER_PORT" envDefault:"9090"`
	// MetricsTLSCertPath is the path to the server certificate for TLS connections
	MetricsTLSCertPath string `env:"METRICS_TLS_CERT_PATH"`
	// MetricsTLSKeyPath is the path to the server private key for TLS connections
	MetricsTLSKeyPath string `env:"METRICS_TLS_KEY_PATH"`
	// MetricsPrefix is the prefix of the metrics that are sent to the server.
	MetricsPrefix string `env:"METRICS_PREFIX" envDefault:"ebpf_agent_"`
	// EnableFlowFilter enables flow filter, default is false.
	EnableFlowFilter bool `env:"ENABLE_FLOW_FILTER" envDefault:"false"`
	// FlowFilterRules list of flow filter rules
	FlowFilterRules string `env:"FLOW_FILTER_RULES"`
	// EnableNetworkEventsMonitoring enables monitoring network plugin events, default is false.
	EnableNetworkEventsMonitoring bool `env:"ENABLE_NETWORK_EVENTS_MONITORING" envDefault:"false"`
	// NetworkEventsMonitoringGroupID to allow ebpf hook to process samples for specific groupID and ignore the rest
	NetworkEventsMonitoringGroupID int `env:"NETWORK_EVENTS_MONITORING_GROUP_ID" envDefault:"10"`
	// EnablePktTranslationTracking allow tracking packets after translation - for example, NAT, default is false.
	EnablePktTranslationTracking bool `env:"ENABLE_PKT_TRANSLATION" envDefault:"false"`
	// EbpfProgramManagerMode is enabled when eBPF manager is handling netobserv ebpf programs life cycle, default is false.
	EbpfProgramManagerMode bool `env:"EBPF_PROGRAM_MANAGER_MODE" envDefault:"false"`
	// BpfManBpfFSPath user configurable ebpf manager mount path
	BpfManBpfFSPath string `env:"BPFMAN_BPF_FS_PATH" envDefault:"/run/netobserv/maps"`

	/* Deprecated configs are listed below this line
	 * See manageDeprecatedConfigs function for details
	 */

	// Deprecated FlowsTargetHost replaced by TargetHost
	FlowsTargetHost string `env:"FLOWS_TARGET_HOST"`
	// Deprecated FlowsTargetPort replaced by TargetPort
	FlowsTargetPort int `env:"FLOWS_TARGET_PORT"`
	// Deprecated PCAServerPort replaced by TargetPort
	PCAServerPort int `env:"PCA_SERVER_PORT"`
}

func manageDeprecatedConfigs(cfg *Config) {
	if len(cfg.FlowsTargetHost) != 0 {
		clog.Infof("Using deprecated FlowsTargetHost %s", cfg.FlowsTargetHost)
		cfg.TargetHost = cfg.FlowsTargetHost
	}

	if cfg.FlowsTargetPort != 0 {
		clog.Infof("Using deprecated FlowsTargetPort %d", cfg.FlowsTargetPort)
		cfg.TargetPort = cfg.FlowsTargetPort
	} else if cfg.PCAServerPort != 0 {
		clog.Infof("Using deprecated PCAServerPort %d", cfg.PCAServerPort)
		cfg.TargetPort = cfg.PCAServerPort
	}
}