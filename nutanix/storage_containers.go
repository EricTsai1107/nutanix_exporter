
package nutanix

import "encoding/json"

type StorageEntity struct {
	Id				string  `json:"id"`
	Name				string  `json:"name"`
	MaxCapacity			uint64	`json:"max_capacity"`
	Stats				map[string]string	`json:"stats"`
	UsageStats			map[string]string	`json:"usage_stats"`
}

type StorageResponse struct {
	Metadata	*NutanixMetadata
	Entities	[]StorageEntity
}

func (n *Nutanix) GetStorageContainers() []StorageEntity {
	resp, _ := n.makeRequest("GET", "/storage_containers/")
	data := json.NewDecoder(resp.Body)

	var d StorageResponse
	data.Decode(&d)

	return d.Entities
}
