package catalog

import (
	"fmt"

	mnt "github.com/moby/sys/mountinfo"
	term "github.com/mt1976/crt"
	supt "github.com/mt1976/mockterm/support"
	mem "github.com/shirou/gopsutil/mem"
	cpu "github.com/shirou/gopsutil/v3/cpu"
	dsk "github.com/shirou/gopsutil/v3/disk"
	hst "github.com/shirou/gopsutil/v3/host"
)

// Catalogs the resources available on a given system.

var debugMode bool = false

func Run(t term.ViewPort, debug bool, path string) {

	debugMode = debug
	hostname := t.Helpers.GetHostName()
	outputFilename := "catalog_" + hostname + "_" + supt.GetTimeStamp() + ".info"

	info := New()
	//X := T

	info.breakData(t, "Cataloging system resources")
	t.Print("Output file : " + outputFilename)
	t.Break()
	info.storeData(t, "Hostname", t.Helpers.GetHostName())
	info.storeData(t, "Machine Name", t.Helpers.GetSytemInfo())
	info.storeData(t, "Username", t.Helpers.GetUsername())
	info.storeData(t, "Current Path", path)

	m, _ := mem.VirtualMemory()
	info.storeData(t, "Total Memory", t.Formatters.Human(m.Total))
	info.storeData(t, "Available Memory", t.Formatters.Human(m.Available))
	info.storeData(t, "Used Memory", t.Formatters.Human(m.Used))
	info.storeData(t, "Used Percent", fmt.Sprintf("%f", m.UsedPercent))
	info.storeData(t, "Free Memory", t.Formatters.Human(m.Free))
	cpuData, _ := cpu.Info()
	//range through each cpu and store the cpu details

	for _, cpu := range cpuData {
		t.Break()
		info.storeData(t, fmt.Sprintf("CPU %d", cpu.CPU), t.Formatters.Human(cpu.CPU))
		info.storeData(t, fmt.Sprintf("CPU %d Model", cpu.CPU), t.Formatters.Human(cpu.ModelName))
		info.storeData(t, fmt.Sprintf("CPU %d Cores", cpu.CPU), t.Formatters.Human(cpu.Cores))
		info.storeData(t, fmt.Sprintf("CPU %d Mhz", cpu.CPU), t.Formatters.Human(cpu.Mhz))
		info.storeData(t, fmt.Sprintf("CPU %d Cache Size", cpu.CPU), t.Formatters.Human(cpu.CacheSize))
		info.storeData(t, fmt.Sprintf("CPU %d Flags", cpu.CPU), t.Formatters.Human(cpu.Flags))
		info.storeData(t, fmt.Sprintf("CPU %d Stepping", cpu.CPU), t.Formatters.Human(cpu.Stepping))
		info.storeData(t, fmt.Sprintf("CPU %d Vendor ID", cpu.CPU), t.Formatters.Human(cpu.VendorID))
		info.storeData(t, fmt.Sprintf("CPU %d Family", cpu.CPU), t.Formatters.Human(cpu.Family))
		info.storeData(t, fmt.Sprintf("CPU %d Model", cpu.CPU), t.Formatters.Human(cpu.Model))
		info.storeData(t, fmt.Sprintf("CPU %d Physical ID", cpu.CPU), t.Formatters.Human(cpu.PhysicalID))
		info.storeData(t, fmt.Sprintf("CPU %d Core ID", cpu.CPU), t.Formatters.Human(cpu.CoreID))
		info.storeData(t, fmt.Sprintf("CPU %d Microcode", cpu.CPU), t.Formatters.Human(cpu.Microcode))
		info.storeData(t, fmt.Sprintf("CPU %d Model Name", cpu.CPU), t.Formatters.Human(cpu.ModelName))
	}

	ht, _ := hst.Info()
	//range through each host and print host info
	t.Break()
	info.storeData(t, fmt.Sprintf("Host %s", "ID"), t.Formatters.Human(ht.HostID))
	info.storeData(t, fmt.Sprintf("Host %s Hostname", ""), t.Formatters.Human(ht.Hostname))
	info.storeData(t, fmt.Sprintf("Host %s Uptime", ""), t.Formatters.Human(ht.Uptime))
	info.storeData(t, fmt.Sprintf("Host %s Boot Time", ""), t.Formatters.Human(ht.BootTime))
	info.storeData(t, fmt.Sprintf("Host %s Procs", ""), t.Formatters.Human(ht.Procs))
	info.storeData(t, fmt.Sprintf("Host %s OS", ""), t.Formatters.Human(ht.OS))
	info.storeData(t, fmt.Sprintf("Host %s Platform", ""), t.Formatters.Human(ht.Platform))
	info.storeData(t, fmt.Sprintf("Host %s Platform Family", ""), t.Formatters.Human(ht.PlatformFamily))
	info.storeData(t, fmt.Sprintf("Host %s Platform Version", ""), t.Formatters.Human(ht.PlatformVersion))
	info.storeData(t, fmt.Sprintf("Host %s Kernel Version", ""), t.Formatters.Human(ht.KernelVersion))
	info.storeData(t, fmt.Sprintf("Host %s Virtualization System", ""), t.Formatters.Human(ht.VirtualizationSystem))
	info.storeData(t, fmt.Sprintf("Host %s Virtualization Role", ""), t.Formatters.Human(ht.VirtualizationRole))
	//info.storeData(t, T,  "Host Info", T.Human( h))

	v, _ := mnt.GetMounts(nil)
	//zz := 0
	for zz, v := range v {
		//info.storeData(t, T,  fmt.Sprintf("Mount %d", zz), T.Human( v))
		t.Break()
		info.storeData(t, fmt.Sprintf("Mount %d ID", zz), t.Formatters.Human(v.ID))
		info.storeData(t, fmt.Sprintf("Mount %d Major", zz), t.Formatters.Human(v.Major))
		info.storeData(t, fmt.Sprintf("Mount %d Minor", zz), t.Formatters.Human(v.Minor))
		info.storeData(t, fmt.Sprintf("Mount %d Root", zz), t.Formatters.Human(v.Root))
		info.storeData(t, fmt.Sprintf("Mount %d Parent", zz), t.Formatters.Human(v.Parent))
		info.storeData(t, fmt.Sprintf("Mount %d Mountpoint", zz), t.Formatters.Human(v.Mountpoint))
		info.storeData(t, fmt.Sprintf("Mount %d Options", zz), t.Formatters.Human(v.Options))
		info.storeData(t, fmt.Sprintf("Mount %d Optional", zz), t.Formatters.Human(v.Optional))
		info.storeData(t, fmt.Sprintf("Mount %d FSType", zz), t.Formatters.Human(v.FSType))
		info.storeData(t, fmt.Sprintf("Mount %d Source", zz), t.Formatters.Human(v.Source))
		info.storeData(t, fmt.Sprintf("Mount %d VFSOptions", zz), t.Formatters.Human(v.VFSOptions))

		usage, _ := dsk.Usage(v.Mountpoint)
		//info.storeData(t, T,  fmt.Sprintf("Disk Usage %d", zz), T.Human( usage))
		info.storeData(t, fmt.Sprintf("Mount %d Total", zz), t.Formatters.Human(usage.Total))
		info.storeData(t, fmt.Sprintf("Mount %d Free", zz), t.Formatters.Human(usage.Free))
		info.storeData(t, fmt.Sprintf("Mount %d Used", zz), t.Formatters.Human(usage.Used))
		info.storeData(t, fmt.Sprintf("Mount %d UsedPercent", zz), t.Formatters.Human(usage.UsedPercent))
		info.storeData(t, fmt.Sprintf("Mount %d InodesTotal", zz), t.Formatters.Human(usage.InodesTotal))
		info.storeData(t, fmt.Sprintf("Mount %d InodesUsed", zz), t.Formatters.Human(usage.InodesUsed))
		info.storeData(t, fmt.Sprintf("Mount %d InodesFree", zz), t.Formatters.Human(usage.InodesFree))
		info.storeData(t, fmt.Sprintf("Mount %d InodesUsedPercent", zz), t.Formatters.Human(usage.InodesUsedPercent))

	}
	if !debugMode {
		// Open output file
		file, err := supt.OpenFile(t, outputFilename)
		if err != nil {
			return
		}
		defer file.Close()
		err = supt.WriteStringSliceToFile(t, file, info.data)
		if err != nil {
			return
		}
	}
}
