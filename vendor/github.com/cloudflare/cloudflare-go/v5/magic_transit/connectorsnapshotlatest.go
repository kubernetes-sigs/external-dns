// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ConnectorSnapshotLatestService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorSnapshotLatestService] method instead.
type ConnectorSnapshotLatestService struct {
	Options []option.RequestOption
}

// NewConnectorSnapshotLatestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorSnapshotLatestService(opts ...option.RequestOption) (r *ConnectorSnapshotLatestService) {
	r = &ConnectorSnapshotLatestService{}
	r.Options = opts
	return
}

// Get latest Snapshots
func (r *ConnectorSnapshotLatestService) List(ctx context.Context, connectorID string, query ConnectorSnapshotLatestListParams, opts ...option.RequestOption) (res *ConnectorSnapshotLatestListResponse, err error) {
	var env ConnectorSnapshotLatestListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/snapshots/latest", query.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectorSnapshotLatestListResponse struct {
	Count float64                                   `json:"count,required"`
	Items []ConnectorSnapshotLatestListResponseItem `json:"items,required"`
	JSON  connectorSnapshotLatestListResponseJSON   `json:"-"`
}

// connectorSnapshotLatestListResponseJSON contains the JSON metadata for the
// struct [ConnectorSnapshotLatestListResponse]
type connectorSnapshotLatestListResponseJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseJSON) RawJSON() string {
	return r.raw
}

// Snapshot
type ConnectorSnapshotLatestListResponseItem struct {
	// Count of failures to reclaim space
	CountReclaimFailures float64 `json:"count_reclaim_failures,required"`
	// Count of reclaimed paths
	CountReclaimedPaths float64 `json:"count_reclaimed_paths,required"`
	// Count of failed snapshot recordings
	CountRecordFailed float64 `json:"count_record_failed,required"`
	// Count of failed snapshot transmissions
	CountTransmitFailures float64 `json:"count_transmit_failures,required"`
	// Time the Snapshot was recorded (seconds since the Unix epoch)
	T float64 `json:"t,required"`
	// Version
	V string `json:"v,required"`
	// Count of processors/cores
	CPUCount float64 `json:"cpu_count"`
	// Percentage of time over a 10 second window that tasks were stalled
	CPUPressure10s float64 `json:"cpu_pressure_10s"`
	// Percentage of time over a 5 minute window that tasks were stalled
	CPUPressure300s float64 `json:"cpu_pressure_300s"`
	// Percentage of time over a 1 minute window that tasks were stalled
	CPUPressure60s float64 `json:"cpu_pressure_60s"`
	// Total stall time (microseconds)
	CPUPressureTotalUs float64 `json:"cpu_pressure_total_us"`
	// Time spent running a virtual CPU or guest OS (milliseconds)
	CPUTimeGuestMs float64 `json:"cpu_time_guest_ms"`
	// Time spent running a niced guest (milliseconds)
	CPUTimeGuestNiceMs float64 `json:"cpu_time_guest_nice_ms"`
	// Time spent in idle state (milliseconds)
	CPUTimeIdleMs float64 `json:"cpu_time_idle_ms"`
	// Time spent wait for I/O to complete (milliseconds)
	CPUTimeIowaitMs float64 `json:"cpu_time_iowait_ms"`
	// Time spent servicing interrupts (milliseconds)
	CPUTimeIrqMs float64 `json:"cpu_time_irq_ms"`
	// Time spent in low-priority user mode (milliseconds)
	CPUTimeNiceMs float64 `json:"cpu_time_nice_ms"`
	// Time spent servicing softirqs (milliseconds)
	CPUTimeSoftirqMs float64 `json:"cpu_time_softirq_ms"`
	// Time stolen (milliseconds)
	CPUTimeStealMs float64 `json:"cpu_time_steal_ms"`
	// Time spent in system mode (milliseconds)
	CPUTimeSystemMs float64 `json:"cpu_time_system_ms"`
	// Time spent in user mode (milliseconds)
	CPUTimeUserMs float64                                             `json:"cpu_time_user_ms"`
	DHCPLeases    []ConnectorSnapshotLatestListResponseItemsDHCPLease `json:"dhcp_leases"`
	Disks         []ConnectorSnapshotLatestListResponseItemsDisk      `json:"disks"`
	// Name of high availability state
	HaState string `json:"ha_state"`
	// Numeric value associated with high availability state (0 = disabled, 1 = active,
	// 2 = standby, 3 = stopped, 4 = fault)
	HaValue    float64                                             `json:"ha_value"`
	Interfaces []ConnectorSnapshotLatestListResponseItemsInterface `json:"interfaces"`
	// Percentage of time over a 10 second window that all tasks were stalled
	IoPressureFull10s float64 `json:"io_pressure_full_10s"`
	// Percentage of time over a 5 minute window that all tasks were stalled
	IoPressureFull300s float64 `json:"io_pressure_full_300s"`
	// Percentage of time over a 1 minute window that all tasks were stalled
	IoPressureFull60s float64 `json:"io_pressure_full_60s"`
	// Total stall time (microseconds)
	IoPressureFullTotalUs float64 `json:"io_pressure_full_total_us"`
	// Percentage of time over a 10 second window that some tasks were stalled
	IoPressureSome10s float64 `json:"io_pressure_some_10s"`
	// Percentage of time over a 3 minute window that some tasks were stalled
	IoPressureSome300s float64 `json:"io_pressure_some_300s"`
	// Percentage of time over a 1 minute window that some tasks were stalled
	IoPressureSome60s float64 `json:"io_pressure_some_60s"`
	// Total stall time (microseconds)
	IoPressureSomeTotalUs float64 `json:"io_pressure_some_total_us"`
	// Boot time (seconds since Unix epoch)
	KernelBtime float64 `json:"kernel_btime"`
	// Number of context switches that the system underwent
	KernelCtxt float64 `json:"kernel_ctxt"`
	// Number of forks since boot
	KernelProcesses float64 `json:"kernel_processes"`
	// Number of processes blocked waiting for I/O
	KernelProcessesBlocked float64 `json:"kernel_processes_blocked"`
	// Number of processes in runnable state
	KernelProcessesRunning float64 `json:"kernel_processes_running"`
	// The fifteen-minute load average
	LoadAverage15m float64 `json:"load_average_15m"`
	// The one-minute load average
	LoadAverage1m float64 `json:"load_average_1m"`
	// The five-minute load average
	LoadAverage5m float64 `json:"load_average_5m"`
	// Number of currently runnable kernel scheduling entities
	LoadAverageCur float64 `json:"load_average_cur"`
	// Number of kernel scheduling entities that currently exist on the system
	LoadAverageMax float64 `json:"load_average_max"`
	// Memory that has been used more recently
	MemoryActiveBytes float64 `json:"memory_active_bytes"`
	// Non-file backed huge pages mapped into user-space page tables
	MemoryAnonHugepagesBytes float64 `json:"memory_anon_hugepages_bytes"`
	// Non-file backed pages mapped into user-space page tables
	MemoryAnonPagesBytes float64 `json:"memory_anon_pages_bytes"`
	// Estimate of how much memory is available for starting new applications
	MemoryAvailableBytes float64 `json:"memory_available_bytes"`
	// Memory used for block device bounce buffers
	MemoryBounceBytes float64 `json:"memory_bounce_bytes"`
	// Relatively temporary storage for raw disk blocks
	MemoryBuffersBytes float64 `json:"memory_buffers_bytes"`
	// In-memory cache for files read from the disk
	MemoryCachedBytes float64 `json:"memory_cached_bytes"`
	// Free CMA (Contiguous Memory Allocator) pages
	MemoryCmaFreeBytes float64 `json:"memory_cma_free_bytes"`
	// Total CMA (Contiguous Memory Allocator) pages
	MemoryCmaTotalBytes float64 `json:"memory_cma_total_bytes"`
	// Total amount of memory currently available to be allocated on the system
	MemoryCommitLimitBytes float64 `json:"memory_commit_limit_bytes"`
	// Amount of memory presently allocated on the system
	MemoryCommittedAsBytes float64 `json:"memory_committed_as_bytes"`
	// Memory which is waiting to get written back to the disk
	MemoryDirtyBytes float64 `json:"memory_dirty_bytes"`
	// The sum of LowFree and HighFree
	MemoryFreeBytes float64 `json:"memory_free_bytes"`
	// Amount of free highmem
	MemoryHighFreeBytes float64 `json:"memory_high_free_bytes"`
	// Total amount of highmem
	MemoryHighTotalBytes float64 `json:"memory_high_total_bytes"`
	// The number of huge pages in the pool that are not yet allocated
	MemoryHugepagesFree float64 `json:"memory_hugepages_free"`
	// Number of huge pages for which a commitment has been made, but no allocation has
	// yet been made
	MemoryHugepagesRsvd float64 `json:"memory_hugepages_rsvd"`
	// Number of huge pages in the pool above the threshold
	MemoryHugepagesSurp float64 `json:"memory_hugepages_surp"`
	// The size of the pool of huge pages
	MemoryHugepagesTotal float64 `json:"memory_hugepages_total"`
	// The size of huge pages
	MemoryHugepagesizeBytes float64 `json:"memory_hugepagesize_bytes"`
	// Memory which has been less recently used
	MemoryInactiveBytes float64 `json:"memory_inactive_bytes"`
	// Kernel allocations that the kernel will attempt to reclaim under memory pressure
	MemoryKReclaimableBytes float64 `json:"memory_k_reclaimable_bytes"`
	// Amount of memory allocated to kernel stacks
	MemoryKernelStackBytes float64 `json:"memory_kernel_stack_bytes"`
	// Amount of free lowmem
	MemoryLowFreeBytes float64 `json:"memory_low_free_bytes"`
	// Total amount of lowmem
	MemoryLowTotalBytes float64 `json:"memory_low_total_bytes"`
	// Files which have been mapped into memory
	MemoryMappedBytes float64 `json:"memory_mapped_bytes"`
	// Amount of memory dedicated to the lowest level of page tables
	MemoryPageTablesBytes float64 `json:"memory_page_tables_bytes"`
	// Memory allocated to the per-cpu alloctor used to back per-cpu allocations
	MemoryPerCPUBytes float64 `json:"memory_per_cpu_bytes"`
	// Percentage of time over a 10 second window that all tasks were stalled
	MemoryPressureFull10s float64 `json:"memory_pressure_full_10s"`
	// Percentage of time over a 5 minute window that all tasks were stalled
	MemoryPressureFull300s float64 `json:"memory_pressure_full_300s"`
	// Percentage of time over a 1 minute window that all tasks were stalled
	MemoryPressureFull60s float64 `json:"memory_pressure_full_60s"`
	// Total stall time (microseconds)
	MemoryPressureFullTotalUs float64 `json:"memory_pressure_full_total_us"`
	// Percentage of time over a 10 second window that some tasks were stalled
	MemoryPressureSome10s float64 `json:"memory_pressure_some_10s"`
	// Percentage of time over a 5 minute window that some tasks were stalled
	MemoryPressureSome300s float64 `json:"memory_pressure_some_300s"`
	// Percentage of time over a 1 minute window that some tasks were stalled
	MemoryPressureSome60s float64 `json:"memory_pressure_some_60s"`
	// Total stall time (microseconds)
	MemoryPressureSomeTotalUs float64 `json:"memory_pressure_some_total_us"`
	// Part of slab that can be reclaimed on memory pressure
	MemorySReclaimableBytes float64 `json:"memory_s_reclaimable_bytes"`
	// Part of slab that cannot be reclaimed on memory pressure
	MemorySUnreclaimBytes float64 `json:"memory_s_unreclaim_bytes"`
	// Amount of memory dedicated to the lowest level of page tables
	MemorySecondaryPageTablesBytes float64 `json:"memory_secondary_page_tables_bytes"`
	// Amount of memory consumed by tmpfs
	MemoryShmemBytes float64 `json:"memory_shmem_bytes"`
	// Memory used by shmem and tmpfs, allocated with huge pages
	MemoryShmemHugepagesBytes float64 `json:"memory_shmem_hugepages_bytes"`
	// Shared memory mapped into user space with huge pages
	MemoryShmemPmdMappedBytes float64 `json:"memory_shmem_pmd_mapped_bytes"`
	// In-kernel data structures cache
	MemorySlabBytes float64 `json:"memory_slab_bytes"`
	// Memory swapped out and back in while still in swap file
	MemorySwapCachedBytes float64 `json:"memory_swap_cached_bytes"`
	// Amount of swap space that is currently unused
	MemorySwapFreeBytes float64 `json:"memory_swap_free_bytes"`
	// Total amount of swap space available
	MemorySwapTotalBytes float64 `json:"memory_swap_total_bytes"`
	// Total usable RAM
	MemoryTotalBytes float64 `json:"memory_total_bytes"`
	// Largest contiguous block of vmalloc area which is free
	MemoryVmallocChunkBytes float64 `json:"memory_vmalloc_chunk_bytes"`
	// Total size of vmalloc memory area
	MemoryVmallocTotalBytes float64 `json:"memory_vmalloc_total_bytes"`
	// Amount of vmalloc area which is used
	MemoryVmallocUsedBytes float64 `json:"memory_vmalloc_used_bytes"`
	// Memory which is actively being written back to the disk
	MemoryWritebackBytes float64 `json:"memory_writeback_bytes"`
	// Memory used by FUSE for temporary writeback buffers
	MemoryWritebackTmpBytes float64 `json:"memory_writeback_tmp_bytes"`
	// Memory consumed by the zswap backend, compressed
	MemoryZSwapBytes float64 `json:"memory_z_swap_bytes"`
	// Amount of anonymous memory stored in zswap, uncompressed
	MemoryZSwappedBytes float64                                          `json:"memory_z_swapped_bytes"`
	Mounts              []ConnectorSnapshotLatestListResponseItemsMount  `json:"mounts"`
	Netdevs             []ConnectorSnapshotLatestListResponseItemsNetdev `json:"netdevs"`
	// Number of ICMP Address Mask Reply messages received
	SnmpIcmpInAddrMaskReps float64 `json:"snmp_icmp_in_addr_mask_reps"`
	// Number of ICMP Address Mask Request messages received
	SnmpIcmpInAddrMasks float64 `json:"snmp_icmp_in_addr_masks"`
	// Number of ICMP messages received with bad checksums
	SnmpIcmpInCsumErrors float64 `json:"snmp_icmp_in_csum_errors"`
	// Number of ICMP Destination Unreachable messages received
	SnmpIcmpInDestUnreachs float64 `json:"snmp_icmp_in_dest_unreachs"`
	// Number of ICMP Echo Reply messages received
	SnmpIcmpInEchoReps float64 `json:"snmp_icmp_in_echo_reps"`
	// Number of ICMP Echo (request) messages received
	SnmpIcmpInEchos float64 `json:"snmp_icmp_in_echos"`
	// Number of ICMP messages received with ICMP-specific errors
	SnmpIcmpInErrors float64 `json:"snmp_icmp_in_errors"`
	// Number of ICMP messages received
	SnmpIcmpInMsgs float64 `json:"snmp_icmp_in_msgs"`
	// Number of ICMP Parameter Problem messages received
	SnmpIcmpInParmProbs float64 `json:"snmp_icmp_in_parm_probs"`
	// Number of ICMP Redirect messages received
	SnmpIcmpInRedirects float64 `json:"snmp_icmp_in_redirects"`
	// Number of ICMP Source Quench messages received
	SnmpIcmpInSrcQuenchs float64 `json:"snmp_icmp_in_src_quenchs"`
	// Number of ICMP Time Exceeded messages received
	SnmpIcmpInTimeExcds float64 `json:"snmp_icmp_in_time_excds"`
	// Number of ICMP Address Mask Request messages received
	SnmpIcmpInTimestampReps float64 `json:"snmp_icmp_in_timestamp_reps"`
	// Number of ICMP Timestamp (request) messages received
	SnmpIcmpInTimestamps float64 `json:"snmp_icmp_in_timestamps"`
	// Number of ICMP Address Mask Reply messages sent
	SnmpIcmpOutAddrMaskReps float64 `json:"snmp_icmp_out_addr_mask_reps"`
	// Number of ICMP Address Mask Request messages sent
	SnmpIcmpOutAddrMasks float64 `json:"snmp_icmp_out_addr_masks"`
	// Number of ICMP Destination Unreachable messages sent
	SnmpIcmpOutDestUnreachs float64 `json:"snmp_icmp_out_dest_unreachs"`
	// Number of ICMP Echo Reply messages sent
	SnmpIcmpOutEchoReps float64 `json:"snmp_icmp_out_echo_reps"`
	// Number of ICMP Echo (request) messages sent
	SnmpIcmpOutEchos float64 `json:"snmp_icmp_out_echos"`
	// Number of ICMP messages which this entity did not send due to ICMP-specific
	// errors
	SnmpIcmpOutErrors float64 `json:"snmp_icmp_out_errors"`
	// Number of ICMP messages attempted to send
	SnmpIcmpOutMsgs float64 `json:"snmp_icmp_out_msgs"`
	// Number of ICMP Parameter Problem messages sent
	SnmpIcmpOutParmProbs float64 `json:"snmp_icmp_out_parm_probs"`
	// Number of ICMP Redirect messages sent
	SnmpIcmpOutRedirects float64 `json:"snmp_icmp_out_redirects"`
	// Number of ICMP Source Quench messages sent
	SnmpIcmpOutSrcQuenchs float64 `json:"snmp_icmp_out_src_quenchs"`
	// Number of ICMP Time Exceeded messages sent
	SnmpIcmpOutTimeExcds float64 `json:"snmp_icmp_out_time_excds"`
	// Number of ICMP Timestamp Reply messages sent
	SnmpIcmpOutTimestampReps float64 `json:"snmp_icmp_out_timestamp_reps"`
	// Number of ICMP Timestamp (request) messages sent
	SnmpIcmpOutTimestamps float64 `json:"snmp_icmp_out_timestamps"`
	// Default value of the Time-To-Live field of the IP header
	SnmpIPDefaultTTL float64 `json:"snmp_ip_default_ttl"`
	// Number of datagrams forwarded to their final destination
	SnmpIPForwDatagrams float64 `json:"snmp_ip_forw_datagrams"`
	// Set when acting as an IP gateway
	SnmpIPForwardingEnabled bool `json:"snmp_ip_forwarding_enabled"`
	// Number of datagrams generated by fragmentation
	SnmpIPFragCreates float64 `json:"snmp_ip_frag_creates"`
	// Number of datagrams discarded because fragmentation failed
	SnmpIPFragFails float64 `json:"snmp_ip_frag_fails"`
	// Number of datagrams successfully fragmented
	SnmpIPFragOks float64 `json:"snmp_ip_frag_oks"`
	// Number of input datagrams discarded due to errors in the IP address
	SnmpIPInAddrErrors float64 `json:"snmp_ip_in_addr_errors"`
	// Number of input datagrams successfully delivered to IP user-protocols
	SnmpIPInDelivers float64 `json:"snmp_ip_in_delivers"`
	// Number of input datagrams otherwise discarded
	SnmpIPInDiscards float64 `json:"snmp_ip_in_discards"`
	// Number of input datagrams discarded due to errors in the IP header
	SnmpIPInHdrErrors float64 `json:"snmp_ip_in_hdr_errors"`
	// Number of input datagrams received from interfaces
	SnmpIPInReceives float64 `json:"snmp_ip_in_receives"`
	// Number of input datagrams discarded due unknown or unsupported protocol
	SnmpIPInUnknownProtos float64 `json:"snmp_ip_in_unknown_protos"`
	// Number of output datagrams otherwise discarded
	SnmpIPOutDiscards float64 `json:"snmp_ip_out_discards"`
	// Number of output datagrams discarded because no route matched
	SnmpIPOutNoRoutes float64 `json:"snmp_ip_out_no_routes"`
	// Number of datagrams supplied for transmission
	SnmpIPOutRequests float64 `json:"snmp_ip_out_requests"`
	// Number of failures detected by the reassembly algorithm
	SnmpIPReasmFails float64 `json:"snmp_ip_reasm_fails"`
	// Number of datagrams successfully reassembled
	SnmpIPReasmOks float64 `json:"snmp_ip_reasm_oks"`
	// Number of fragments received which needed to be reassembled
	SnmpIPReasmReqds float64 `json:"snmp_ip_reasm_reqds"`
	// Number of seconds fragments are held while awaiting reassembly
	SnmpIPReasmTimeout float64 `json:"snmp_ip_reasm_timeout"`
	// Number of times TCP transitions to SYN-SENT from CLOSED
	SnmpTCPActiveOpens float64 `json:"snmp_tcp_active_opens"`
	// Number of times TCP transitions to CLOSED from SYN-SENT or SYN-RCVD, plus
	// transitions to LISTEN from SYN-RCVD
	SnmpTCPAttemptFails float64 `json:"snmp_tcp_attempt_fails"`
	// Number of TCP connections in ESTABLISHED or CLOSE-WAIT
	SnmpTCPCurrEstab float64 `json:"snmp_tcp_curr_estab"`
	// Number of times TCP transitions to CLOSED from ESTABLISHED or CLOSE-WAIT
	SnmpTCPEstabResets float64 `json:"snmp_tcp_estab_resets"`
	// Number of TCP segments received with checksum errors
	SnmpTCPInCsumErrors float64 `json:"snmp_tcp_in_csum_errors"`
	// Number of TCP segments received in error
	SnmpTCPInErrs float64 `json:"snmp_tcp_in_errs"`
	// Number of TCP segments received
	SnmpTCPInSegs float64 `json:"snmp_tcp_in_segs"`
	// Limit on the total number of TCP connections
	SnmpTCPMaxConn float64 `json:"snmp_tcp_max_conn"`
	// Number of TCP segments sent with RST flag
	SnmpTCPOutRsts float64 `json:"snmp_tcp_out_rsts"`
	// Number of TCP segments sent
	SnmpTCPOutSegs float64 `json:"snmp_tcp_out_segs"`
	// Number of times TCP transitions to SYN-RCVD from LISTEN
	SnmpTCPPassiveOpens float64 `json:"snmp_tcp_passive_opens"`
	// Number of TCP segments retransmitted
	SnmpTCPRetransSegs float64 `json:"snmp_tcp_retrans_segs"`
	// Maximum value permitted by a TCP implementation for the retransmission timeout
	// (milliseconds)
	SnmpTCPRtoMax float64 `json:"snmp_tcp_rto_max"`
	// Minimum value permitted by a TCP implementation for the retransmission timeout
	// (milliseconds)
	SnmpTCPRtoMin float64 `json:"snmp_tcp_rto_min"`
	// Number of UDP datagrams delivered to UDP applications
	SnmpUdpInDatagrams float64 `json:"snmp_udp_in_datagrams"`
	// Number of UDP datagrams failed to be delivered for reasons other than lack of
	// application at the destination port
	SnmpUdpInErrors float64 `json:"snmp_udp_in_errors"`
	// Number of UDP datagrams received for which there was not application at the
	// destination port
	SnmpUdpNoPorts float64 `json:"snmp_udp_no_ports"`
	// Number of UDP datagrams sent
	SnmpUdpOutDatagrams float64 `json:"snmp_udp_out_datagrams"`
	// Boottime of the system (seconds since the Unix epoch)
	SystemBootTimeS float64                                           `json:"system_boot_time_s"`
	Thermals        []ConnectorSnapshotLatestListResponseItemsThermal `json:"thermals"`
	Tunnels         []ConnectorSnapshotLatestListResponseItemsTunnel  `json:"tunnels"`
	// Sum of how much time each core has spent idle
	UptimeIdleMs float64 `json:"uptime_idle_ms"`
	// Uptime of the system, including time spent in suspend
	UptimeTotalMs float64                                     `json:"uptime_total_ms"`
	JSON          connectorSnapshotLatestListResponseItemJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemJSON contains the JSON metadata for the
// struct [ConnectorSnapshotLatestListResponseItem]
type connectorSnapshotLatestListResponseItemJSON struct {
	CountReclaimFailures           apijson.Field
	CountReclaimedPaths            apijson.Field
	CountRecordFailed              apijson.Field
	CountTransmitFailures          apijson.Field
	T                              apijson.Field
	V                              apijson.Field
	CPUCount                       apijson.Field
	CPUPressure10s                 apijson.Field
	CPUPressure300s                apijson.Field
	CPUPressure60s                 apijson.Field
	CPUPressureTotalUs             apijson.Field
	CPUTimeGuestMs                 apijson.Field
	CPUTimeGuestNiceMs             apijson.Field
	CPUTimeIdleMs                  apijson.Field
	CPUTimeIowaitMs                apijson.Field
	CPUTimeIrqMs                   apijson.Field
	CPUTimeNiceMs                  apijson.Field
	CPUTimeSoftirqMs               apijson.Field
	CPUTimeStealMs                 apijson.Field
	CPUTimeSystemMs                apijson.Field
	CPUTimeUserMs                  apijson.Field
	DHCPLeases                     apijson.Field
	Disks                          apijson.Field
	HaState                        apijson.Field
	HaValue                        apijson.Field
	Interfaces                     apijson.Field
	IoPressureFull10s              apijson.Field
	IoPressureFull300s             apijson.Field
	IoPressureFull60s              apijson.Field
	IoPressureFullTotalUs          apijson.Field
	IoPressureSome10s              apijson.Field
	IoPressureSome300s             apijson.Field
	IoPressureSome60s              apijson.Field
	IoPressureSomeTotalUs          apijson.Field
	KernelBtime                    apijson.Field
	KernelCtxt                     apijson.Field
	KernelProcesses                apijson.Field
	KernelProcessesBlocked         apijson.Field
	KernelProcessesRunning         apijson.Field
	LoadAverage15m                 apijson.Field
	LoadAverage1m                  apijson.Field
	LoadAverage5m                  apijson.Field
	LoadAverageCur                 apijson.Field
	LoadAverageMax                 apijson.Field
	MemoryActiveBytes              apijson.Field
	MemoryAnonHugepagesBytes       apijson.Field
	MemoryAnonPagesBytes           apijson.Field
	MemoryAvailableBytes           apijson.Field
	MemoryBounceBytes              apijson.Field
	MemoryBuffersBytes             apijson.Field
	MemoryCachedBytes              apijson.Field
	MemoryCmaFreeBytes             apijson.Field
	MemoryCmaTotalBytes            apijson.Field
	MemoryCommitLimitBytes         apijson.Field
	MemoryCommittedAsBytes         apijson.Field
	MemoryDirtyBytes               apijson.Field
	MemoryFreeBytes                apijson.Field
	MemoryHighFreeBytes            apijson.Field
	MemoryHighTotalBytes           apijson.Field
	MemoryHugepagesFree            apijson.Field
	MemoryHugepagesRsvd            apijson.Field
	MemoryHugepagesSurp            apijson.Field
	MemoryHugepagesTotal           apijson.Field
	MemoryHugepagesizeBytes        apijson.Field
	MemoryInactiveBytes            apijson.Field
	MemoryKReclaimableBytes        apijson.Field
	MemoryKernelStackBytes         apijson.Field
	MemoryLowFreeBytes             apijson.Field
	MemoryLowTotalBytes            apijson.Field
	MemoryMappedBytes              apijson.Field
	MemoryPageTablesBytes          apijson.Field
	MemoryPerCPUBytes              apijson.Field
	MemoryPressureFull10s          apijson.Field
	MemoryPressureFull300s         apijson.Field
	MemoryPressureFull60s          apijson.Field
	MemoryPressureFullTotalUs      apijson.Field
	MemoryPressureSome10s          apijson.Field
	MemoryPressureSome300s         apijson.Field
	MemoryPressureSome60s          apijson.Field
	MemoryPressureSomeTotalUs      apijson.Field
	MemorySReclaimableBytes        apijson.Field
	MemorySUnreclaimBytes          apijson.Field
	MemorySecondaryPageTablesBytes apijson.Field
	MemoryShmemBytes               apijson.Field
	MemoryShmemHugepagesBytes      apijson.Field
	MemoryShmemPmdMappedBytes      apijson.Field
	MemorySlabBytes                apijson.Field
	MemorySwapCachedBytes          apijson.Field
	MemorySwapFreeBytes            apijson.Field
	MemorySwapTotalBytes           apijson.Field
	MemoryTotalBytes               apijson.Field
	MemoryVmallocChunkBytes        apijson.Field
	MemoryVmallocTotalBytes        apijson.Field
	MemoryVmallocUsedBytes         apijson.Field
	MemoryWritebackBytes           apijson.Field
	MemoryWritebackTmpBytes        apijson.Field
	MemoryZSwapBytes               apijson.Field
	MemoryZSwappedBytes            apijson.Field
	Mounts                         apijson.Field
	Netdevs                        apijson.Field
	SnmpIcmpInAddrMaskReps         apijson.Field
	SnmpIcmpInAddrMasks            apijson.Field
	SnmpIcmpInCsumErrors           apijson.Field
	SnmpIcmpInDestUnreachs         apijson.Field
	SnmpIcmpInEchoReps             apijson.Field
	SnmpIcmpInEchos                apijson.Field
	SnmpIcmpInErrors               apijson.Field
	SnmpIcmpInMsgs                 apijson.Field
	SnmpIcmpInParmProbs            apijson.Field
	SnmpIcmpInRedirects            apijson.Field
	SnmpIcmpInSrcQuenchs           apijson.Field
	SnmpIcmpInTimeExcds            apijson.Field
	SnmpIcmpInTimestampReps        apijson.Field
	SnmpIcmpInTimestamps           apijson.Field
	SnmpIcmpOutAddrMaskReps        apijson.Field
	SnmpIcmpOutAddrMasks           apijson.Field
	SnmpIcmpOutDestUnreachs        apijson.Field
	SnmpIcmpOutEchoReps            apijson.Field
	SnmpIcmpOutEchos               apijson.Field
	SnmpIcmpOutErrors              apijson.Field
	SnmpIcmpOutMsgs                apijson.Field
	SnmpIcmpOutParmProbs           apijson.Field
	SnmpIcmpOutRedirects           apijson.Field
	SnmpIcmpOutSrcQuenchs          apijson.Field
	SnmpIcmpOutTimeExcds           apijson.Field
	SnmpIcmpOutTimestampReps       apijson.Field
	SnmpIcmpOutTimestamps          apijson.Field
	SnmpIPDefaultTTL               apijson.Field
	SnmpIPForwDatagrams            apijson.Field
	SnmpIPForwardingEnabled        apijson.Field
	SnmpIPFragCreates              apijson.Field
	SnmpIPFragFails                apijson.Field
	SnmpIPFragOks                  apijson.Field
	SnmpIPInAddrErrors             apijson.Field
	SnmpIPInDelivers               apijson.Field
	SnmpIPInDiscards               apijson.Field
	SnmpIPInHdrErrors              apijson.Field
	SnmpIPInReceives               apijson.Field
	SnmpIPInUnknownProtos          apijson.Field
	SnmpIPOutDiscards              apijson.Field
	SnmpIPOutNoRoutes              apijson.Field
	SnmpIPOutRequests              apijson.Field
	SnmpIPReasmFails               apijson.Field
	SnmpIPReasmOks                 apijson.Field
	SnmpIPReasmReqds               apijson.Field
	SnmpIPReasmTimeout             apijson.Field
	SnmpTCPActiveOpens             apijson.Field
	SnmpTCPAttemptFails            apijson.Field
	SnmpTCPCurrEstab               apijson.Field
	SnmpTCPEstabResets             apijson.Field
	SnmpTCPInCsumErrors            apijson.Field
	SnmpTCPInErrs                  apijson.Field
	SnmpTCPInSegs                  apijson.Field
	SnmpTCPMaxConn                 apijson.Field
	SnmpTCPOutRsts                 apijson.Field
	SnmpTCPOutSegs                 apijson.Field
	SnmpTCPPassiveOpens            apijson.Field
	SnmpTCPRetransSegs             apijson.Field
	SnmpTCPRtoMax                  apijson.Field
	SnmpTCPRtoMin                  apijson.Field
	SnmpUdpInDatagrams             apijson.Field
	SnmpUdpInErrors                apijson.Field
	SnmpUdpNoPorts                 apijson.Field
	SnmpUdpOutDatagrams            apijson.Field
	SystemBootTimeS                apijson.Field
	Thermals                       apijson.Field
	Tunnels                        apijson.Field
	UptimeIdleMs                   apijson.Field
	UptimeTotalMs                  apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Snapshot DHCP lease
type ConnectorSnapshotLatestListResponseItemsDHCPLease struct {
	// Client ID of the device the IP Address was leased to
	ClientID string `json:"client_id,required"`
	// Expiry time of the DHCP lease (seconds since the Unix epoch)
	ExpiryTime float64 `json:"expiry_time,required"`
	// Hostname of the device the IP Address was leased to
	Hostname string `json:"hostname,required"`
	// Name of the network interface
	InterfaceName string `json:"interface_name,required"`
	// IP Address that was leased
	IPAddress string `json:"ip_address,required"`
	// MAC Address of the device the IP Address was leased to
	MacAddress string `json:"mac_address,required"`
	// Connector identifier
	ConnectorID string                                                `json:"connector_id"`
	JSON        connectorSnapshotLatestListResponseItemsDHCPLeaseJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsDHCPLeaseJSON contains the JSON metadata
// for the struct [ConnectorSnapshotLatestListResponseItemsDHCPLease]
type connectorSnapshotLatestListResponseItemsDHCPLeaseJSON struct {
	ClientID      apijson.Field
	ExpiryTime    apijson.Field
	Hostname      apijson.Field
	InterfaceName apijson.Field
	IPAddress     apijson.Field
	MacAddress    apijson.Field
	ConnectorID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsDHCPLease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsDHCPLeaseJSON) RawJSON() string {
	return r.raw
}

// Snapshot Disk
type ConnectorSnapshotLatestListResponseItemsDisk struct {
	// I/Os currently in progress
	InProgress float64 `json:"in_progress,required"`
	// Device major number
	Major float64 `json:"major,required"`
	// Reads merged
	Merged float64 `json:"merged,required"`
	// Device minor number
	Minor float64 `json:"minor,required"`
	// Device name
	Name string `json:"name,required"`
	// Reads completed successfully
	Reads float64 `json:"reads,required"`
	// Sectors read successfully
	SectorsRead float64 `json:"sectors_read,required"`
	// Sectors written successfully
	SectorsWritten float64 `json:"sectors_written,required"`
	// Time spent doing I/Os (milliseconds)
	TimeInProgressMs float64 `json:"time_in_progress_ms,required"`
	// Time spent reading (milliseconds)
	TimeReadingMs float64 `json:"time_reading_ms,required"`
	// Time spent writing (milliseconds)
	TimeWritingMs float64 `json:"time_writing_ms,required"`
	// Weighted time spent doing I/Os (milliseconds)
	WeightedTimeInProgressMs float64 `json:"weighted_time_in_progress_ms,required"`
	// Writes completed
	Writes float64 `json:"writes,required"`
	// Writes merged
	WritesMerged float64 `json:"writes_merged,required"`
	// Connector identifier
	ConnectorID string `json:"connector_id"`
	// Discards completed successfully
	Discards float64 `json:"discards"`
	// Discards merged
	DiscardsMerged float64 `json:"discards_merged"`
	// Flushes completed successfully
	Flushes float64 `json:"flushes"`
	// Sectors discarded
	SectorsDiscarded float64 `json:"sectors_discarded"`
	// Time spent discarding (milliseconds)
	TimeDiscardingMs float64 `json:"time_discarding_ms"`
	// Time spent flushing (milliseconds)
	TimeFlushingMs float64                                          `json:"time_flushing_ms"`
	JSON           connectorSnapshotLatestListResponseItemsDiskJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsDiskJSON contains the JSON metadata for
// the struct [ConnectorSnapshotLatestListResponseItemsDisk]
type connectorSnapshotLatestListResponseItemsDiskJSON struct {
	InProgress               apijson.Field
	Major                    apijson.Field
	Merged                   apijson.Field
	Minor                    apijson.Field
	Name                     apijson.Field
	Reads                    apijson.Field
	SectorsRead              apijson.Field
	SectorsWritten           apijson.Field
	TimeInProgressMs         apijson.Field
	TimeReadingMs            apijson.Field
	TimeWritingMs            apijson.Field
	WeightedTimeInProgressMs apijson.Field
	Writes                   apijson.Field
	WritesMerged             apijson.Field
	ConnectorID              apijson.Field
	Discards                 apijson.Field
	DiscardsMerged           apijson.Field
	Flushes                  apijson.Field
	SectorsDiscarded         apijson.Field
	TimeDiscardingMs         apijson.Field
	TimeFlushingMs           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsDisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsDiskJSON) RawJSON() string {
	return r.raw
}

// Snapshot Interface
type ConnectorSnapshotLatestListResponseItemsInterface struct {
	// Name of the network interface
	Name string `json:"name,required"`
	// UP/DOWN state of the network interface
	Operstate string `json:"operstate,required"`
	// Connector identifier
	ConnectorID string                                                        `json:"connector_id"`
	IPAddresses []ConnectorSnapshotLatestListResponseItemsInterfacesIPAddress `json:"ip_addresses"`
	// Speed of the network interface (bits per second)
	Speed float64                                               `json:"speed"`
	JSON  connectorSnapshotLatestListResponseItemsInterfaceJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsInterfaceJSON contains the JSON metadata
// for the struct [ConnectorSnapshotLatestListResponseItemsInterface]
type connectorSnapshotLatestListResponseItemsInterfaceJSON struct {
	Name        apijson.Field
	Operstate   apijson.Field
	ConnectorID apijson.Field
	IPAddresses apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsInterface) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsInterfaceJSON) RawJSON() string {
	return r.raw
}

// Snapshot Interface Address
type ConnectorSnapshotLatestListResponseItemsInterfacesIPAddress struct {
	// Name of the network interface
	InterfaceName string `json:"interface_name,required"`
	// IP address of the network interface
	IPAddress string `json:"ip_address,required"`
	// Connector identifier
	ConnectorID string                                                          `json:"connector_id"`
	JSON        connectorSnapshotLatestListResponseItemsInterfacesIPAddressJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsInterfacesIPAddressJSON contains the
// JSON metadata for the struct
// [ConnectorSnapshotLatestListResponseItemsInterfacesIPAddress]
type connectorSnapshotLatestListResponseItemsInterfacesIPAddressJSON struct {
	InterfaceName apijson.Field
	IPAddress     apijson.Field
	ConnectorID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsInterfacesIPAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsInterfacesIPAddressJSON) RawJSON() string {
	return r.raw
}

// Snapshot Mount
type ConnectorSnapshotLatestListResponseItemsMount struct {
	// File system on disk (EXT4, NTFS, etc.)
	FileSystem string `json:"file_system,required"`
	// Kind of disk (HDD, SSD, etc.)
	Kind string `json:"kind,required"`
	// Path where disk is mounted
	MountPoint string `json:"mount_point,required"`
	// Name of the disk mount
	Name string `json:"name,required"`
	// Available disk size (bytes)
	AvailableBytes float64 `json:"available_bytes"`
	// Connector identifier
	ConnectorID string `json:"connector_id"`
	// Determines whether the disk is read-only
	IsReadOnly bool `json:"is_read_only"`
	// Determines whether the disk is removable
	IsRemovable bool `json:"is_removable"`
	// Total disk size (bytes)
	TotalBytes float64                                           `json:"total_bytes"`
	JSON       connectorSnapshotLatestListResponseItemsMountJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsMountJSON contains the JSON metadata for
// the struct [ConnectorSnapshotLatestListResponseItemsMount]
type connectorSnapshotLatestListResponseItemsMountJSON struct {
	FileSystem     apijson.Field
	Kind           apijson.Field
	MountPoint     apijson.Field
	Name           apijson.Field
	AvailableBytes apijson.Field
	ConnectorID    apijson.Field
	IsReadOnly     apijson.Field
	IsRemovable    apijson.Field
	TotalBytes     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsMount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsMountJSON) RawJSON() string {
	return r.raw
}

// Snapshot Netdev
type ConnectorSnapshotLatestListResponseItemsNetdev struct {
	// Name of the network device
	Name string `json:"name,required"`
	// Total bytes received
	RecvBytes float64 `json:"recv_bytes,required"`
	// Compressed packets received
	RecvCompressed float64 `json:"recv_compressed,required"`
	// Packets dropped
	RecvDrop float64 `json:"recv_drop,required"`
	// Bad packets received
	RecvErrs float64 `json:"recv_errs,required"`
	// FIFO overruns
	RecvFifo float64 `json:"recv_fifo,required"`
	// Frame alignment errors
	RecvFrame float64 `json:"recv_frame,required"`
	// Multicast packets received
	RecvMulticast float64 `json:"recv_multicast,required"`
	// Total packets received
	RecvPackets float64 `json:"recv_packets,required"`
	// Total bytes transmitted
	SentBytes float64 `json:"sent_bytes,required"`
	// Number of packets not sent due to carrier errors
	SentCarrier float64 `json:"sent_carrier,required"`
	// Number of collisions
	SentColls float64 `json:"sent_colls,required"`
	// Number of compressed packets transmitted
	SentCompressed float64 `json:"sent_compressed,required"`
	// Number of packets dropped during transmission
	SentDrop float64 `json:"sent_drop,required"`
	// Number of transmission errors
	SentErrs float64 `json:"sent_errs,required"`
	// FIFO overruns
	SentFifo float64 `json:"sent_fifo,required"`
	// Total packets transmitted
	SentPackets float64 `json:"sent_packets,required"`
	// Connector identifier
	ConnectorID string                                             `json:"connector_id"`
	JSON        connectorSnapshotLatestListResponseItemsNetdevJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsNetdevJSON contains the JSON metadata
// for the struct [ConnectorSnapshotLatestListResponseItemsNetdev]
type connectorSnapshotLatestListResponseItemsNetdevJSON struct {
	Name           apijson.Field
	RecvBytes      apijson.Field
	RecvCompressed apijson.Field
	RecvDrop       apijson.Field
	RecvErrs       apijson.Field
	RecvFifo       apijson.Field
	RecvFrame      apijson.Field
	RecvMulticast  apijson.Field
	RecvPackets    apijson.Field
	SentBytes      apijson.Field
	SentCarrier    apijson.Field
	SentColls      apijson.Field
	SentCompressed apijson.Field
	SentDrop       apijson.Field
	SentErrs       apijson.Field
	SentFifo       apijson.Field
	SentPackets    apijson.Field
	ConnectorID    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsNetdev) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsNetdevJSON) RawJSON() string {
	return r.raw
}

// Snapshot Thermal
type ConnectorSnapshotLatestListResponseItemsThermal struct {
	// Sensor identifier for the component
	Label string `json:"label,required"`
	// Connector identifier
	ConnectorID string `json:"connector_id"`
	// Critical failure temperature of the component (degrees Celsius)
	CriticalCelcius float64 `json:"critical_celcius"`
	// Current temperature of the component (degrees Celsius)
	CurrentCelcius float64 `json:"current_celcius"`
	// Maximum temperature of the component (degrees Celsius)
	MaxCelcius float64                                             `json:"max_celcius"`
	JSON       connectorSnapshotLatestListResponseItemsThermalJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsThermalJSON contains the JSON metadata
// for the struct [ConnectorSnapshotLatestListResponseItemsThermal]
type connectorSnapshotLatestListResponseItemsThermalJSON struct {
	Label           apijson.Field
	ConnectorID     apijson.Field
	CriticalCelcius apijson.Field
	CurrentCelcius  apijson.Field
	MaxCelcius      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsThermal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsThermalJSON) RawJSON() string {
	return r.raw
}

// Snapshot Tunnels
type ConnectorSnapshotLatestListResponseItemsTunnel struct {
	// Name of tunnel health state (unknown, healthy, degraded, down)
	HealthState string `json:"health_state,required"`
	// Numeric value associated with tunnel state (0 = unknown, 1 = healthy, 2 =
	// degraded, 3 = down)
	HealthValue float64 `json:"health_value,required"`
	// The tunnel interface name (i.e. xfrm1, xfrm3.99, etc.)
	InterfaceName string `json:"interface_name,required"`
	// Tunnel identifier
	TunnelID string `json:"tunnel_id,required"`
	// Connector identifier
	ConnectorID string                                             `json:"connector_id"`
	JSON        connectorSnapshotLatestListResponseItemsTunnelJSON `json:"-"`
}

// connectorSnapshotLatestListResponseItemsTunnelJSON contains the JSON metadata
// for the struct [ConnectorSnapshotLatestListResponseItemsTunnel]
type connectorSnapshotLatestListResponseItemsTunnelJSON struct {
	HealthState   apijson.Field
	HealthValue   apijson.Field
	InterfaceName apijson.Field
	TunnelID      apijson.Field
	ConnectorID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseItemsTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseItemsTunnelJSON) RawJSON() string {
	return r.raw
}

type ConnectorSnapshotLatestListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectorSnapshotLatestListResponseEnvelope struct {
	Result   ConnectorSnapshotLatestListResponse                   `json:"result,required"`
	Success  bool                                                  `json:"success,required"`
	Errors   []ConnectorSnapshotLatestListResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorSnapshotLatestListResponseEnvelopeMessages `json:"messages"`
	JSON     connectorSnapshotLatestListResponseEnvelopeJSON       `json:"-"`
}

// connectorSnapshotLatestListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ConnectorSnapshotLatestListResponseEnvelope]
type connectorSnapshotLatestListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorSnapshotLatestListResponseEnvelopeErrors struct {
	Code    float64                                               `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    connectorSnapshotLatestListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorSnapshotLatestListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ConnectorSnapshotLatestListResponseEnvelopeErrors]
type connectorSnapshotLatestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorSnapshotLatestListResponseEnvelopeMessages struct {
	Code    float64                                                 `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    connectorSnapshotLatestListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorSnapshotLatestListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ConnectorSnapshotLatestListResponseEnvelopeMessages]
type connectorSnapshotLatestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorSnapshotLatestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorSnapshotLatestListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
