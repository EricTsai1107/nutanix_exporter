
package collector

import (
	"github.com/erictsai1107/nutanix_exporter/nutanix"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type VirtualDisksStat struct {
	HelpText	string
	Labels		[]string
}

var (
	virtualdisksNamespace  string = "nutanix_virtual_disk"
	virtualdisksLabels	[]string = []string{"host_name", "id", "NFS_file_path"}
)

var virtualdisksStats map[string]string = map[string]string {
	"hypervisor_avg_io_latency_usecs": "EMPTY",
        "hypervisor_write_io_bandwidth_kBps": "EMPTY",
        "controller.random_ops_ppm": "EMPTY",
        "controller.storage_tier.ssd.usage_bytes": "EMPTY",
        "read_io_ppm": "EMPTY",
        "controller.frontend_read_latency_histogram_1000us": "EMPTY",
        "controller_num_iops": "EMPTY",
        "controller.frontend_write_ops": "EMPTY",
        "controller.frontend_write_latency_histogram_10000us": "EMPTY",
        "controller.read_size_histogram_1024kB": "EMPTY",
        "total_read_io_time_usecs": "EMPTY",
        "controller_total_read_io_time_usecs": "EMPTY",
        "controller.wss_3600s_write_MB": "EMPTY",
        "controller.frontend_read_latency_histogram_50000us": "EMPTY",
        "controller.frontend_read_latency_histogram_2000us": "EMPTY",
        "controller_num_write_io": "EMPTY",
        "controller.read_source_cache_ssd_bytes": "EMPTY",
        "controller.read_source_oplog_bytes": "EMPTY",
        "controller.read_source_cache_dram_bytes": "EMPTY",
        "controller.random_read_ops": "EMPTY",
        "controller_total_io_time_usecs": "EMPTY",
        "controller_num_seq_io": "EMPTY",
        "controller_total_io_size_kbytes": "EMPTY",
        "controller.wss_120s_write_MB": "EMPTY",
        "controller.read_source_block_store_bytes": "EMPTY",
        "controller_num_io": "EMPTY",
        "controller.read_source_estore_zero_bytes": "EMPTY",
        "controller_num_random_io": "EMPTY",
        "hypervisor_num_read_io": "EMPTY",
        "hypervisor_total_read_io_time_usecs": "EMPTY",
        "num_io": "EMPTY",
        "hypervisor_num_write_io": "EMPTY",
        "controller.write_size_histogram_32kB": "EMPTY",
        "controller.frontend_read_latency_histogram_20000us": "EMPTY",
        "controller.read_size_histogram_32kB": "EMPTY",
        "hypervisor_num_write_iops": "EMPTY",
        "avg_io_latency_usecs": "EMPTY",
        "controller_write_io_ppm": "EMPTY",
        "controller.read_source_estore_ssd_bytes": "EMPTY",
        "hypervisor_total_read_io_size_kbytes": "EMPTY",
        "controller_num_write_iops": "EMPTY",
        "total_io_time_usecs": "EMPTY",
        "controller.wss_3600s_read_MB": "EMPTY",
        "controller.summary_read_source_ssd_bytes_per_sec": "EMPTY",
        "controller.write_size_histogram_16kB": "EMPTY",
        "total_transformed_usage_bytes": "EMPTY",
        "avg_write_io_latency_usecs": "EMPTY",
        "controller.cse_target_90_percent_write_MB": "EMPTY",
        "num_read_io": "EMPTY",
        "hypervisor_read_io_bandwidth_kBps": "EMPTY",
        "hypervisor_total_io_time_usecs": "EMPTY",
        "num_random_io": "EMPTY",
        "controller.write_dest_estore_bytes": "EMPTY",
        "controller.frontend_write_latency_histogram_5000us": "EMPTY",
        "controller.storage_tier.das-sata.pinned_usage_bytes": "EMPTY",
        "num_write_io": "EMPTY",
        "controller.frontend_write_latency_histogram_2000us": "EMPTY",
        "io_bandwidth_kBps": "EMPTY",
        "controller.write_size_histogram_512kB": "EMPTY",
        "controller.read_size_histogram_16kB": "EMPTY",
        "write_io_ppm": "EMPTY",
        "controller_avg_write_io_latency_usecs": "EMPTY",
        "controller.frontend_read_latency_histogram_100000us": "EMPTY",
        "num_read_iops": "EMPTY",
        "controller.summary_read_source_hdd_bytes_per_sec": "EMPTY",
        "controller.read_source_extent_cache_bytes": "EMPTY",
        "timespan_usecs": "EMPTY",
        "controller_num_read_iops": "EMPTY",
        "controller.frontend_read_latency_histogram_10000us": "EMPTY",
        "controller.write_size_histogram_64kB": "EMPTY",
        "controller.frontend_write_latency_histogram_0us": "EMPTY",
        "controller.frontend_write_latency_histogram_100000us": "EMPTY",
        "hypervisor_num_io": "EMPTY",
        "controller_total_transformed_usage_bytes": "EMPTY",
        "avg_read_io_latency_usecs": "EMPTY",
        "controller_total_read_io_size_kbytes": "EMPTY",
        "controller_read_io_ppm": "EMPTY",
        "controller.frontend_ops": "EMPTY",
        "controller.wss_120s_read_MB": "EMPTY",
        "controller.read_size_histogram_512kB": "EMPTY",
        "hypervisor_avg_read_io_latency_usecs": "EMPTY",
        "controller.write_size_histogram_1024kB": "EMPTY",
        "controller.write_dest_block_store_bytes": "EMPTY",
        "controller.read_size_histogram_4kB": "EMPTY",
        "num_write_iops": "EMPTY",
        "controller.random_ops_per_sec": "EMPTY",
        "num_iops": "EMPTY",
        "controller.storage_tier.cloud.pinned_usage_bytes": "EMPTY",
        "controller_avg_io_latency_usecs": "EMPTY",
        "controller.read_size_histogram_8kB": "EMPTY",
        "controller_num_read_io": "EMPTY",
        "controller_seq_io_ppm": "EMPTY",
        "controller_read_io_bandwidth_kBps": "EMPTY",
        "controller_io_bandwidth_kBps": "EMPTY",
        "controller.read_size_histogram_0kB": "EMPTY",
        "controller.random_ops": "EMPTY",
        "hypervisor_timespan_usecs": "EMPTY",
        "total_read_io_size_kbytes": "EMPTY",
        "hypervisor_total_io_size_kbytes": "EMPTY",
        "controller.frontend_ops_per_sec": "EMPTY",
        "controller.write_dest_oplog_bytes": "EMPTY",
        "controller.frontend_write_latency_histogram_1000us": "EMPTY",
        "hypervisor_num_read_iops": "EMPTY",
        "controller.summary_read_source_cache_bytes_per_sec": "EMPTY",
        "controller_write_io_bandwidth_kBps": "EMPTY",
        "controller_user_bytes": "EMPTY",
        "hypervisor_avg_write_io_latency_usecs": "EMPTY",
        "controller.storage_tier.ssd.pinned_usage_bytes": "EMPTY",
        "read_io_bandwidth_kBps": "EMPTY",
        "controller.frontend_read_ops": "EMPTY",
        "hypervisor_num_iops": "EMPTY",
        "hypervisor_io_bandwidth_kBps": "EMPTY",
        "controller.wss_120s_union_MB": "EMPTY",
        "controller.read_source_estore_hdd_bytes": "EMPTY",
        "controller_random_io_ppm": "EMPTY",
        "controller.cse_target_90_percent_read_MB": "EMPTY",
        "controller.storage_tier.das-sata.usage_bytes": "EMPTY",
        "controller.frontend_read_latency_histogram_5000us": "EMPTY",
        "controller_avg_read_io_size_kbytes": "EMPTY",
        "write_io_bandwidth_kBps": "EMPTY",
        "controller.random_read_ops_per_sec": "EMPTY",
        "controller.read_size_histogram_64kB": "EMPTY",
        "controller.wss_3600s_union_MB": "EMPTY",
        "random_io_ppm": "EMPTY",
        "total_untransformed_usage_bytes": "EMPTY",
        "controller.frontend_read_latency_histogram_0us": "EMPTY",
        "controller.random_write_ops": "EMPTY",
        "controller_avg_write_io_size_kbytes": "EMPTY",
        "controller_avg_read_io_latency_usecs": "EMPTY",
        "total_io_size_kbytes": "EMPTY",
        "controller.storage_tier.cloud.usage_bytes": "EMPTY",
        "controller.frontend_write_latency_histogram_50000us": "EMPTY",
        "controller.write_size_histogram_8kB": "EMPTY",
        "controller_timespan_usecs": "EMPTY",
        "num_seq_io": "EMPTY",
        "controller.write_size_histogram_4kB": "EMPTY",
        "seq_io_ppm": "EMPTY",
        "controller.write_size_histogram_0kB": "EMPTY",
}

type VirtualDisksExporter struct {
        DiskMb          *prometheus.GaugeVec
	Stats		map[string]*prometheus.GaugeVec
}


func (e *VirtualDisksExporter) Describe(ch chan<- *prometheus.Desc) {
	e.DiskMb = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: virtualdisksNamespace, Name: "disk_mb", Help: "Disk Size of Virtual Disk",}, virtualdisksLabels, )
        e.DiskMb.Describe(ch)


	e.Stats = make(map[string]*prometheus.GaugeVec)
	for k, h := range virtualdisksStats {
		name := normalizeFQN(k)
		e.Stats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: virtualdisksNamespace, Name: name, Help: h,}, virtualdisksLabels, )
		e.Stats[k].Describe(ch)
	}
}

func (e *VirtualDisksExporter) Collect(ch chan<- prometheus.Metric) {
	virtualdisks := nutanixApi.GetVirtualDisks()

	for _, s := range virtualdisks {
	        {
		        g := e.DiskMb.WithLabelValues(s.HostName, s.Id, s.NutanixNFSFilePath)
			g.Set(float64(s.DiskMb))
	                g.Collect(ch)
		}


		for i, k := range e.Stats {
			v, _ := strconv.ParseFloat(s.Stats[i], 64)
			g := k.WithLabelValues(s.HostName, s.Id, s.NutanixNFSFilePath)
			g.Set(v)
			g.Collect(ch)
		}
	}
}

func NewVirtualDisksExporter(api *nutanix.Nutanix) *VirtualDisksExporter {
	nutanixApi = api
	return &VirtualDisksExporter{}
}

