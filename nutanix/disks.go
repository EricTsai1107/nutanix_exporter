
package nutanix

import (
	"encoding/json"
)

type DiskResponse struct {
	Metadata	*DiskMetadata
	Entities	[]*DiskEntity
}

type DiskMetadata struct {
	GrandTotalEntites	float64	`json:"grand_total_entities"`
	TotalEntites		float64	`json:"total_entities"`
	Count			float64	`json:"count"`
}

type DiskEntity struct {
        Id              string  `json:"Id"`
        MountPath       string  `json:"mount_path"`
	HostName	string  `json:"node_name"`
        Stats           map[string]string       `json:"stats"`
        UsageStats      map[string]string       `json:"usage_stats"`
	StorageTierName string  `json:"storage_tier_name"`
	DiskStatus	string	`json:"disk_status"`
	Online		int	`json:"online"`
	DiskSize	int64	`json:"disk_size"`
	Location	string	`json:"location"`
}

func (n *Nutanix) GetDisks() []*DiskEntity {
	resp, _ := n.makeRequest("GET", "/disks/")
	data := json.NewDecoder(resp.Body)

	var d DiskResponse
	data.Decode(&d)

	return d.Entities
}
