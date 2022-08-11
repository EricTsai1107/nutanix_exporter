
package collector

//import "encoding/json"
import (
	"github.com/crlintsai/nutanix_exporter/nutanix"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/log"
)

type VmsStat struct {
	HelpText	string
	Labels		[]string
}

var (
	vmsNamespace string = "nutanix_vms"
	vmsLabels	  []string = []string{"vmname", "hostname"}
)

var vmsStats map[string]string = map[string]string {
	"hypervisor_avg_io_latency_usecs": "...",
	"num_read_iops": "...",
	"hypervisor_write_io_bandwidth_kBps": "...",
	"timespan_usecs": "...",
	"controller_num_read_iops": "...",
        "controller.storage_tier.ssd.usage_bytes": "...",
	"read_io_ppm": "...",
	"controller_num_iops": "...",
	"total_read_io_time_usecs": "...",
	"controller_total_read_io_time_usecs": "...",
	"hypervisor_num_io": "...",
	"controller_total_transformed_usage_bytes": "...",
	"hypervisor_cpu_usage_ppm": "...",
	"controller_num_write_io": "...",
	"avg_read_io_latency_usecs": "...",
	"guest.memory_swapped_in_bytes": "...",
	"controller_total_io_time_usecs": "...",
	"memory_usage_ppm": "...",
	"controller_total_read_io_size_kbytes": "...",
	"controller_num_seq_io": "...",
	"controller_read_io_ppm": "...",
	"controller_total_io_size_kbytes": "...",
	"hypervisor.cpu_ready_time_ppm": "...",
	"controller_num_io": "...",
	"hypervisor_avg_read_io_latency_usecs": "...",
	"num_write_iops": "...",
	"controller_num_random_io": "...",
	"num_iops": "...",
	"guest.memory_usage_ppm": "...",
	"hypervisor_num_read_io": "...",
	"hypervisor_total_read_io_time_usecs": "...",
	"controller_avg_io_latency_usecs": "...",
	"num_io": "...",
	"controller_num_read_io": "...",
	"hypervisor_num_write_io": "...",
	"controller_seq_io_ppm": "...",
	"guest.memory_usage_bytes": "...",
	"controller_read_io_bandwidth_kBps": "...",
	"controller_io_bandwidth_kBps": "...",
	"hypervisor_num_received_bytes": "...",
	"hypervisor_timespan_usecs": "...",
	"hypervisor_num_write_iops": "...",
	"total_read_io_size_kbytes": "...",
	"hypervisor_total_io_size_kbytes": "...",
	"avg_io_latency_usecs": "...",
	"hypervisor_num_read_iops": "...",
	"hypervisor_swap_in_rate_kBps": "...",
	"controller_write_io_bandwidth_kBps": "...",
	"controller_write_io_ppm": "...",
	"controller_user_bytes": "...",
	"hypervisor_avg_write_io_latency_usecs": "...",
	"hypervisor_num_transmitted_bytes": "...",
	"hypervisor_total_read_io_size_kbytes": "...",
	"read_io_bandwidth_kBps": "...",
	"guest.memory_swapped_out_bytes": "...",
	"hypervisor_memory_usage_ppm": "...",
	"hypervisor_num_iops": "...",
	"hypervisor_io_bandwidth_kBps": "...",
	"controller_num_write_iops": "...",
	"total_io_time_usecs": "...",
	"controller_random_io_ppm": "...",
	"controller.storage_tier.das-sata.usage_bytes": "...",
	"controller_avg_read_io_size_kbytes": "...",
	"hypervisor_swap_out_rate_kBps": "...",
	"total_transformed_usage_bytes": "...",
	"avg_write_io_latency_usecs": "...",
	"num_read_io": "...",
	"write_io_bandwidth_kBps": "...",
	"hypervisor_read_io_bandwidth_kBps": "...",
	"hypervisor_consumed_memory_bytes": "...",
	"random_io_ppm": "...",
	"total_untransformed_usage_bytes": "...",
	"hypervisor_total_io_time_usecs": "...",
	"num_random_io": "...",
	"controller_avg_write_io_size_kbytes": "...",
	"controller_avg_read_io_latency_usecs": "...",
	"num_write_io": "...",
	"total_io_size_kbytes": "...",
	"controller.storage_tier.cloud.usage_bytes": "...",
	"io_bandwidth_kBps": "...",
	"controller_timespan_usecs": "...",
	"num_seq_io": "...",
	"seq_io_ppm": "...",
	"write_io_ppm": "...",
	"controller_avg_write_io_latency_usecs": "...",
}

var vmsUsageStats map[string]string = map[string]string {
	"gpu_usage_ppm": "...",
	"framebuffer_usage_ppm": "...",
}

type VmsExporter struct {
	Stats		map[string]*prometheus.GaugeVec
	UsageStats	map[string]*prometheus.GaugeVec
}

func (e *VmsExporter) Describe(ch chan<- *prometheus.Desc) {
	e.Stats = make(map[string]*prometheus.GaugeVec)
	for k, h := range vmsStats {
		name := normalizeFQN(k)
		e.Stats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: vmsNamespace, Name: name, Help: h,}, vmsLabels, )
		e.Stats[k].Describe(ch)
	}

	e.UsageStats = make(map[string]*prometheus.GaugeVec)
	for k, h := range vmsUsageStats {
		name := normalizeFQN(k)
		e.UsageStats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: vmsNamespace, Name: name, Help: h,}, vmsLabels, )
		e.UsageStats[k].Describe(ch)
	}
}

func (e *VmsExporter) Collect(ch chan<- prometheus.Metric) {
	vms := nutanixApi.GetVms()
	for _, s := range vms {
		for i, k := range e.UsageStats {
			v, _ := strconv.ParseFloat(s.UsageStats[i], 64)
			g := k.WithLabelValues(s.Name, s.HostName).Set(v)
			g.Collect(ch)
		}
		for i, k := range e.Stats {
			v, _ := strconv.ParseFloat(s.Stats[i], 64)
			g := k.WithLabelValues(s.Name, s.HostName).Set(v)
			g.Collect(ch)
		}
	}
}

func NewVmsExporter(api *nutanix.Nutanix) *VmsExporter {
	nutanixApi = api
	return &VmsExporter{}
}

