package api

import (
	"testing"

	"gopkg.in/yaml.v2"
	"gotest.tools/assert"
)

func TestConfiguration(t *testing.T) {
	configYAML := `---
id: go-minion1
location: Test
brokerUrl: 10.0.0.100:8990
brokerProperties:
  tls-enabled: "true"
trapPort: 11162
syslogPort: 11514
listeners:
- name: Netflow-5
  port: 18877
  parser: Netflow5UdpParser
  properties:
    workers: "4"
- name: Netflow-9
  port: 14729
  parser: Netflow9UdpParser
`
	config := &MinionConfig{}
	err := yaml.Unmarshal([]byte(configYAML), config)
	assert.NilError(t, err)
	assert.NilError(t, config.IsValid())
	assert.Equal(t, "go-minion1", config.ID)
	assert.Equal(t, 2, len(config.Listeners))
	assert.Assert(t, config.BrokerProperties != nil)
	assert.Equal(t, "true", config.BrokerProperties["tls-enabled"])

	netflow := config.GetListener("Netflow-5")
	assert.Assert(t, netflow != nil)
	assert.Equal(t, 18877, netflow.Port)
	assert.Assert(t, netflow.Is("Netflow5UdpParser"))
	assert.Assert(t, netflow.Properties != nil)
	assert.Equal(t, "4", netflow.Properties["workers"])

	sflow := config.GetListener("SFlow")
	assert.Assert(t, sflow == nil)

	netflow = config.GetListenerByParser("Netflow9UdpParser")
	assert.Assert(t, netflow != nil)
	assert.Equal(t, 14729, netflow.Port)
	assert.Assert(t, netflow.Properties == nil)

	sflow = config.GetListenerByParser("SFlowUdpParser")
	assert.Assert(t, sflow == nil)

	listeners := []string{
		"Graphite,12003,ForwardParser",
		"NXOS,50000,NxosGrpcParser",
		"Wrong1,1000",
		"Wrong2,1001",
	}
	config.ParseListeners(listeners)
	assert.Equal(t, 4, len(config.Listeners))
	assert.Assert(t, config.GetListener("Graphite").Is("ForwardParser"))
}
