
package nutanix

import (
	"encoding/json"
)

type HostResponse struct {
	Metadata	*HostMetadata
	Entities	[]HostEntity
}

type HostMetadata struct {

}

type HostEntity struct {
	Name		string  `json:"name"`
	CpuFrequency	int64	`json:"cpu_frequency_in_hz"`
	CpuCapacity	int64	`json:"cpu_capacity_in_hz"`
	MemoryCapacity	int64	`json:"memory_capacity_in_bytes"`
	NumCpuCores	int	`json:"num_cpu_cores`
	NumVms		int	`json:"num_vms"`
	BootTime	int64	`json:"boot_time_in_usecs"`
	Stats		map[string]string `json:"stats"`
	UsageStats	map[string]string `json:"usage_stats"`
}

func (n *Nutanix) GetHosts() []HostEntity {
	resp, _ := n.makeRequest("GET", "/hosts/")
	data := json.NewDecoder(resp.Body)

	var d HostResponse
	data.Decode(&d)

	return d.Entities
}
