
package nutanix

import (
	"encoding/json"
)

type VmResponse struct {
	Metadata	*VmMetadata
	Entities	[]*VmEntity
}

type VmMetadata struct {
	GrandTotalEntites	float64	`json:"grand_total_entities"`
	TotalEntites		float64	`json:"total_entities"`
	Count			float64	`json:"count"`
}

type VmEntity struct {
        Id              string  `json:"vmId"`
        Name            string  `json:"vmName"`
	HostName	string  `json:"hostName"`
        Stats           map[string]string       `json:"stats"`
        UsageStats      map[string]string       `json:"usage_stats"`
	NumVCpus	float64	`json:"numVCpus"`
	MemoryMb	float64	`json:"memoryCapacityInBytes"`
	MemoryCapMb	float64 `json:"memoryReservedCapacityInBytes"`
	DiskMb		float64 `json:"diskCapacityInBytes"`
	PowerState	string	`json:"powerState"`
        ClusterName     string
}

func (n *Nutanix) GetVms() []*VmEntity {
	resp, _ := n.makeRequest("GET", "/vms/")
	data := json.NewDecoder(resp.Body)

	var d VmResponse
	data.Decode(&d)

	return d.Entities
}
