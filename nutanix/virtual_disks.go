
package nutanix

import (
	"encoding/json"
)

type VirtualDiskResponse struct {
	Metadata	*VirtualDiskMetadata
	Entities	[]*VirtualDiskEntity
}

type VirtualDiskMetadata struct {
	GrandTotalEntites	float64	`json:"grand_total_entities"`
	TotalEntites		float64	`json:"total_entities"`
	Count			float64	`json:"count"`
}

type VirtualDiskEntity struct {
        Id              string  `json:"virtual_disk_id"`
        HostName            string  `json:"attached_vmname"`
        Stats           map[string]string       `json:"stats"`
	DiskMb		float64 `json:"disk_capacity_in_bytes"`
	NutanixNFSFilePath	string	`json:"nutanix_nfsfile_path"`
}

func (n *Nutanix) GetVirtualDisks() []*VirtualDiskEntity {
	resp, _ := n.makeRequest("GET", "/virtual_disks/")
	data := json.NewDecoder(resp.Body)

	var d VirtualDiskResponse
	data.Decode(&d)

	return d.Entities
}
