
package nutanix

import (
	"encoding/json"
)

type ClusterResponse struct {
        Metadata        *ClusterMetadata
        Entities        []Cluster
}

type ClusterMetadata struct {

}

type Cluster struct {
	Id				string  `json:"id"`
	Uuid				string  `json:"uuid"`
	Name				string  `json:"name"`
	NumNodes			int	`json:"num_nodes"`
	SsdPinningPercentageLimit	int	`json:"ssd_pinning_percentage_limit"`
	RackableUnits			[]RackableUnits	`json:"rackable_units`
	Stats				map[string]string `json:"stats"`
	UsageStats			map[string]string `json:"usage_stats"`
}

type RackableUnits struct {
	Id			int   `json:"id"`
	RackableUnitUuid	string `json:"rackable_unit_uuid"`
	Model			string `json:"model"`
	ModelName		string `json:"model_name"`
	Serial			string `json:"serial"`
	Positions		[]string  `json:"positions`
	Nodes			[]int  `json:"nodes"`
	NodeUUids		[]string `json:"node_uuids"`
}

func (n *Nutanix) GetCluster() []Cluster {
	resp, _ := n.makeRequest("GET", "/clusters/")
	data := json.NewDecoder(resp.Body)

	var d ClusterResponse
	data.Decode(&d)

	return d.Entities
}
