package api

import (
	"github.com/agalue/gominion/protobuf/ipc"
)

// Broker represents a broker implementation
type Broker interface {

	// Sends a Sink Message to OpenNMS
	Send(msg *ipc.SinkMessage) error
}

// SinkModule represents the Sink Module interface
type SinkModule interface {

	// Returns the ID of the Sink Module implementation
	GetID() string

	// Starts the Sink Module
	// Expected to be a blocking operation (run it within a GoRoutine)
	Start(config *MinionConfig, broker Broker)

	// Shutdown the Sink Module
	Stop()
}

// RPCModule represents the RPC Module interface
type RPCModule interface {

	// Returns the ID of the RPC Module implementation
	GetID() string

	// Executes an RPC request, and returns the response
	// The response tells if the operation was successful or not
	Execute(request *ipc.RpcRequestProto) *ipc.RpcResponseProto
}

// ServiceCollector represents the service collector interface
type ServiceCollector interface {

	// Returns the ID of the Service Collector implementation
	GetID() string

	// Executes the data collection operation from the request, and returns the collected metrics
	// The response tells if the operation was successful or not
	Collect(request *CollectorRequestDTO) CollectorResponseDTO
}

// ServiceDetector represents the service detector interface
type ServiceDetector interface {

	// Returns the ID of the Service Detector implementation
	GetID() string

	// Executes the detection operation from the request, and returns if it was successful or not
	Detect(request *DetectorRequestDTO) DetectResults
}

// ServiceMonitor represents the service monitor interface
type ServiceMonitor interface {

	// Returns the ID of the Service Monitor implementation
	GetID() string

	// Executes the polling operation from the request, and returns the status
	Poll(request *PollerRequestDTO) PollStatus
}
