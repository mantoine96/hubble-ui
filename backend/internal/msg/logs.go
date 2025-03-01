package msg

const (
	FlowsError                 = "flow error: %v\n"
	BackendToUIStreamIsClosed  = "stream (ui client <-> ui backend) is closed\n"
	NoEventsRequired           = "no events requested, aborted.\n"
	LinkBetweenServicesChanged = "hubble-ui.ServiceLink changed: %s\n"
	ServiceChanged             = "hubble-ui.Service changed: %s"
	SkippingZeroIdentitySvc    = "%s svc identity == 0. Skipping.\n"
	ZeroIdentityInSourceOrDest = "invalid (zero) identity in source / dest services\n"
	PrintZeroIdentityFlowJSON  = "zero identity flow json: %v\n"

	ServerSetupGRPCClientError     = "failed to setup grpc client: %v\n"
	ServerSetupK8sClientsetError   = "failed to create k8s clientset: %v\n"
	ServerSetupNoRelayAddrError    = "hubble-relay addr is empty, flows fetching is impossible\n"
	ServerSetupRelayClientReady    = "hubble-relay grpc client created (hubble-relay addr: %s)\n"
	ServerSetupRunError            = "failed to run server: %v\n"
	ServerSetupUsingDefRelayAddr   = "hubble-relay addr is set to default (%s)\n"
	ServerSetupUsingDefPort        = "server port is set to default (%s)\n"
	ServerSetupListeningAt         = "listening at: %s\n"
	ServerSetupConfigInitError     = "failed to initialize hubble-ui backend: %s\n"
	ServerSetupTLSAllowInsecureDef = "TLSAllowInsecure option set to: %v\n"
	ServerSetupTLSInitWithNCACerts = "initialized with %d CA certificates\n"
	ServerSetupTLSInitState        = "initialized with TLS %s\n"
	ServerSetupRelayClientError    = "failed to create a hubble-relay client from user context %v\n"

	SendConnStateError           = "failed to send connection state (%s) to the client: %v\n"
	SendK8sUnavailableNotifError = "failed to send k8s unavailable notification: %v\n"
	SendK8sConnectedNotifError   = "failed to send k8s connected notification: %v\n"
	SendFlowsError               = "failed to send bunch of flows to the client: %v\n"
	SendNamespaceEventError      = "failed to send namespaces event to the client: %v\n"
	SendServiceEventError        = "failed to send service event to the client: %v\n"
	SendLinkEventError           = "failed to send link event to the client: %v\n"
	SendHubbleStatusError        = "failed to send hubble status update to the client: %v\n"

	GetFlowsConnectingToRelay      = "fetching hubble flows: connecting to hubble-relay (attempt #%d)\n"
	GetFlowsConnectionAttemptError = "fetching hubble flows: connecting to hubble-relay (attempt #%d) failed: %v\n"
	GetFlowsConnectedToRelay       = "fetching hubble flows: connection to hubble-relay established\n"
	GetFlowsUIStreamisClosed       = "fetching hubble flows: stream (ui backend <-> hubble-relay) is closed\n"
	GetFlowsLastParseError         = "failed to convert GET_FLOWS_LAST to int, falling back to default value (10000): %v\n"
	GetFlowsSinceParseError        = "failed to parse GET_FLOWS_SINCE: %v\n"
	GetFlowsSinceUseError          = "failed to use GET_FLOWS_SINCE: %v\n"

	HubbleStatusCheckerIsRunning      = "running hubble status checker\n"
	HubbleStatusCriticalError         = "hubble status checker: critical error: %v\n"
	HubbleStatusRelayUnavailableError = "hubble status checker: failed to connect to hubble-relay: %v\n"
	HubbleStatusCheckerIsStopped      = "hubble status checker: stopped\n"
	HubbleStatusCheckerRelayConnected = "hubble status checker: connection to hubble-relay established\n"

	NSWatcherK8sTimeoutError     = "fetching namespaces failed, k8s timeout error: %v\n"
	NSWatcherK8sUnavailableError = "fetching namespaces failed, k8s api is unavailable: %v\n"
	NSWatcherUnknownError        = "fetching namespaces unknown error: %v\n"
)
