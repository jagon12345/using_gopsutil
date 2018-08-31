package systeminfo

import (
	"encoding/json"
	"log"
	"fmt"
	"os_monit/utils"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type CpuInfo struct {
	CpuCounts   int             `json:"cpucounts,-"`
	CpuTimes    []cpu.TimesStat `json:"-"`
	CpuPercent  []float64       `json:"-"`
	CpuInfoStat []cpu.InfoStat  `json:"cpuinfo,-"`
}

func HandleCpuInfo() (string, error) {
	info := GetCpuInfo()
	for i := range info.CpuInfoStat {
		// clear flags value
		info.CpuInfoStat[i].Flags = nil
		//enc := json.NewEncoder(os.Stdout)
		//enc.Encode(info)
	}
	jsonInfo, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(jsonInfo), nil

	//for _, cTimes := range(info.cpuTimes) {
	//	k, err := json.MarshalIndent(cTimes, "", "  ")
	//	if err != nil {return}
	//	fmt.Println(string(k))
	//}

	//for _, cPercent := range(info.cpuPercent) {
	//	k, err := json.MarshalIndent(cPercent, "", "  ")
	//	if err != nil {return}
	//	fmt.Println(string(k))
	//}

}

func GetCpuInfo() *CpuInfo {
	cpuCounts, _ := cpu.Counts(true)
	cpuPercent, _ := cpu.Percent(0, true)
	cpuInfoStat, _ := cpu.Info()
	cpuTimes, _ := cpu.Times(true)
	cpuInfo := &CpuInfo{
		CpuCounts:   cpuCounts,
		CpuTimes:    cpuTimes,
		CpuPercent:  cpuPercent,
		CpuInfoStat: cpuInfoStat,
	}
	return cpuInfo
}

type DiskInfo struct {
	AllPartition []disk.PartitionStat             `json:"partitions,-"`
	Usage        []disk.UsageStat                   `json:"diskusage,-"`
	IOCounters   [](map[string](disk.IOCountersStat)) `json:"iocounters,-"`
}

func HandleDiskInfo() (string, error) {
	info := GetDiskInfo()
	//newUsage := []disk.UsageStat{}
	//newIOCounters := [](map[string](disk.IOCountersStat)){}
	for i := range(info.AllPartition) {
		// remove device which is not ext4, and clear data for opts and fstype.
		if info.AllPartition[i].Fstype != "ext4" {
			info.AllPartition = append(info.AllPartition[:i], info.AllPartition[i+1:]...)
			continue
		}
		info.AllPartition[i].Opts = ""
		info.AllPartition[i].Fstype = ""
		u, _ := disk.Usage(info.AllPartition[i].Mountpoint)
		u.Total = utils.IBytes(u.Total)
		u.Used = utils.IBytes(u.Used)
		u.Free = utils.IBytes(u.Free)
		info.Usage = append(info.Usage, *u)
		c, _ := disk.IOCounters(info.AllPartition[i].Mountpoint)
		info.IOCounters = append(info.IOCounters, c)
	}
	k, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return "s", nil
	}
	fmt.Println(string(k))
	//}
	return "ss", nil
}

func GetDiskInfo() *DiskInfo {
	// only get physical device.
	allPartition, _ := disk.Partitions(false)
	// just examples.
	diskInfo := &DiskInfo{
		AllPartition: allPartition,
	}
	return diskInfo
}

//func DockerInfo() {
//	dockerInfo, _ := docker.GetDockerStat()
//	fmt.Println(dockerInfo)
//}

type HostInfo struct {
	BootTime            uint64         `json:"boottime,-"`
	KernelVersion       string         `json:"kernelversion,-"`
	PlatformInfo        [3]string         `json:"platforminfo,-"`
	Uptime              uint64         `json:"uptime,-"`
	Virtualization      [2]string         `json:"-"`
	SensorsTemperatures []host.TemperatureStat    `json:"-"`
	Users				[]host.UserStat
	HostInfo            *host.InfoStat `json:"hostinfo,-"`
}

func HandleHostInfo() (string, error) {
	info := GetHostInfo
	_ = info
	return "ss", nil
}

func GetHostInfo() *HostInfo {
	bootTime, _ := host.BootTime()
	kernelVersion, _ := host.KernelVersion()
	platform, family, version, _ := host.PlatformInformation()
	uptime, _ := host.Uptime()
	vsystem, vrole, _ := host.Virtualization()
	sensorsTemperatures, _ := host.SensorsTemperatures()
	users, _ := host.Users()
	info, _ := host.Info()

	hostInfo := &HostInfo{
		BootTime:            bootTime,
		KernelVersion:       kernelVersion,
		PlatformInfo:        [3]string{platform, family, version},
		Uptime:              uptime,
		Virtualization:      [2]string{vsystem, vrole},
		SensorsTemperatures: sensorsTemperatures,
		Users:               users,
		HostInfo:            info,
	}
	return hostInfo
}

type LoadInfo struct {
	LoadAvg  *load.AvgStat  `json:"loadavg,-"`
	LoadMisc *load.MiscStat `json:"-"`
}

func HandleLoadInfo() (string, error) {
	info := GetLoadInfo()
	_ = info
	return "ss", nil
}

func GetLoadInfo() *LoadInfo {
	avg, _ := load.Avg()
	misc, _ := load.Misc()
	loadInfo := &LoadInfo{
		LoadAvg:  avg,
		LoadMisc: misc,
	}
	return loadInfo

}

type MemInfo struct {
	SwapMemory    *mem.SwapMemoryStat    `json:"swapmemory,-"`
	VirtualMemory *mem.VirtualMemoryStat `json:"virtualmemory,-"`
}

func HandleMemInfo() (string, error) {
	info := GetMemInfo()
	_ = info
	return "ss", nil
}
func GetMemInfo() *MemInfo {
	swapMemory, _ := mem.SwapMemory()
	virtualMemory, _ := mem.VirtualMemory()
	memInfo := &MemInfo{
		SwapMemory:    swapMemory,
		VirtualMemory: virtualMemory,
	}
	return memInfo
}

type NetInfo struct {
	Pids           []int32                 `json:"-"`
	Connections    []net.ConnectionStat    `json:""`
	//ConnectionsPid []net.ConnectionStat    `json : ""`
	FilterCounters []net.FilterStat        `json:"-"`
	IOCounters     []net.IOCountersStat     `json:"-"`
	Interfaces     []net.InterfaceStat     `json:"-"`
	ProtoCounters  []net.ProtoCountersStat `json:"-"`
}

func HandleNetInfo() (string, error) {
	info := GetNetInfo()
	_ = info
	return "ss", nil
}
func GetNetInfo() *NetInfo {
	pids, _ := net.Pids()
	connections, _ := net.Connections("all")
	//connectionsPid, _ := net.ConnectionsPid("1")
	filterCounters, err := net.FilterCounters()
	if err != nil {
		return nil
	}
	ioCounters, _ := net.IOCounters(true)
	interfaces, _ := net.Interfaces()
	protoCounters, _ := net.ProtoCounters(nil)
	netInfo := &NetInfo{
		Pids:           pids,
		Connections:    connections,
		//ConnectionsPid: connectionsPid,
		FilterCounters: filterCounters,
		IOCounters:     ioCounters,
		Interfaces:     interfaces,
		ProtoCounters:  protoCounters,
	}
	return netInfo
}

//func Users() *users {

//}
