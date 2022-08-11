
package collector

import (
	"github.com/crlintsai/nutanix-exporter/nutanix"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type DisksStat struct {
	HelpText	string
	Labels		[]string
}

var (
	disksNamespace  string = "nutanix_disk"
	disksLabels	[]string = []string{"mount_path", "node_name", "location", "disk_status"}
)

var disksStats map[string]string = map[string]string {
	"hypervisor_avg_io_latency_usecs": "...",
	"num_read_iops": "...",
	"hypervisor_write_io_bandwidth_kBps": "...",
	"timespan_usecs": "...",
	"controller_num_read_iops": "...",
	"read_io_ppm": "...",
	"controller_num_iops": "...",
	"total_read_io_time_usecs": "...",
	"controller_total_read_io_time_usecs": "...",
	"hypervisor_num_io": "...",
	"controller_total_transformed_usage_bytes": "...",
	"controller_num_write_io": "...",
	"avg_read_io_latency_usecs": "...",
	"controller_total_io_time_usecs": "...",
	"controller_total_read_io_size_kbytes": "...",
	"controller_num_seq_io": "...",
	"controller_read_io_ppm": "...",
	"controller_total_io_size_kbytes": "...",
	"controller_num_io": "...",
	"hypervisor_avg_read_io_latency_usecs": "...",
	"num_write_iops": "...",
	"controller_num_random_io": "...",
	"num_iops": "...",
	"hypervisor_num_read_io": "...",
	"hypervisor_total_read_io_time_usecs": "...",
	"controller_avg_io_latency_usecs": "...",
	"num_io": "...",
	"controller_num_read_io": "...",
	"hypervisor_num_write_io": "...",
	"controller_seq_io_ppm": "...",
	"controller_read_io_bandwidth_kBps": "...",
	"controller_io_bandwidth_kBps": "...",
	"hypervisor_timespan_usecs": "...",
	"hypervisor_num_write_iops": "...",
	"total_read_io_size_kbytes": "...",
	"hypervisor_total_io_size_kbytes": "...",
	"avg_io_latency_usecs": "...",
	"hypervisor_num_read_iops": "...",
	"controller_write_io_bandwidth_kBps": "...",
	"controller_write_io_ppm": "...",
	"hypervisor_avg_write_io_latency_usecs": "...",
	"hypervisor_total_read_io_size_kbytes": "...",
	"read_io_bandwidth_kBps": "...",
	"hypervisor_num_iops": "...",
	"hypervisor_io_bandwidth_kBps": "...",
	"controller_num_write_iops": "...",
	"total_io_time_usecs": "...",
	"controller_random_io_ppm": "...",
	"controller_avg_read_io_size_kbytes": "...",
	"total_transformed_usage_bytes": "...",
	"avg_write_io_latency_usecs": "...",
	"num_read_io": "...",
	"write_io_bandwidth_kBps": "...",
	"hypervisor_read_io_bandwidth_kBps": "...",
	"random_io_ppm": "...",
	"total_untransformed_usage_bytes": "...",
	"hypervisor_total_io_time_usecs": "...",
	"num_random_io": "...",
	"controller_avg_write_io_size_kbytes": "...",
	"controller_avg_read_io_latency_usecs": "...",
	"num_write_io": "...",
	"total_io_size_kbytes": "...",
	"io_bandwidth_kBps": "...",
	"controller_timespan_usecs": "...",
	"num_seq_io": "...",
	"seq_io_ppm": "...",
	"write_io_ppm": "...",
	"controller_avg_write_io_latency_usecs": "...",
}

var disksUsageStats map[string]string = map[string]string {
	"storage.logical_usage_bytes": "...",
	"storage.capacity_bytes": "...",
	"storage.free_bytes": "...",
	"storage.usage_bytes": "...",
}

type DisksExporter struct {
        Online          *prometheus.GaugeVec
        DiskSize        *prometheus.GaugeVec
	Stats		map[string]*prometheus.GaugeVec
	UsageStats	map[string]*prometheus.GaugeVec
}


func (e *DisksExporter) Describe(ch chan<- *prometheus.Desc) {
	e.Online = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: disksNamespace, Name: "online", Help: "Online Status of Disk",}, disksLabels, )
        e.Online.Describe(ch)

        e.DiskSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: disksNamespace, Name: "disksize", Help: "Disk Size of Disk",}, disksLabels, )
        e.DiskSize.Describe(ch)


	e.Stats = make(map[string]*prometheus.GaugeVec)
	for k, h := range disksStats {
		name := normalizeFQN(k)
		e.Stats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: disksNamespace, Name: name, Help: h,}, disksLabels, )
		e.Stats[k].Describe(ch)
	}

	e.UsageStats = make(map[string]*prometheus.GaugeVec)
	for k, h := range disksUsageStats {
		name := normalizeFQN(k)
		e.UsageStats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: disksNamespace, Name: name, Help: h,}, disksLabels, )
		e.UsageStats[k].Describe(ch)
	}
}

func (e *DisksExporter) Collect(ch chan<- prometheus.Metric) {
	disks := nutanixApi.GetDisks()

	for _, s := range disks {
                {
                        g := e.Online.WithLabelValues(s.MountPath, s.Name, s.Location, s.DiskStatus)
                        g.Set(int(s.Online))
                        g.Collect(ch)

                        g = e.DiskSize.WithLabelValues(s.MountPath, s.Name, s.Location, s.DiskStatus)
                        g.Set(int(s.DiskSize))
                        g.Collect(ch)
                }

		for i, k := range e.UsageStats {
			v, _ := strconv.ParseFloat(s.UsageStats[i], 64)
			g := k.WithLabelValues(s.MountPath, s.HostName, s.Location, s.DiskStatus)
			g.Set(v)
			g.Collect(ch)
		}
		for i, k := range e.Stats {
			v, _ := strconv.ParseFloat(s.Stats[i], 64)
			g := k.WithLabelValues(s.MountPath, s.HostName, s.Location, s.DiskStatus)
			g.Set(v)
			g.Collect(ch)
		}
	}
}

func NewDisksExporter(api *nutanix.Nutanix) *DisksExporter {
	nutanixApi = api
	return &DisksExporter{}
}

