package catalog

import (
	"fmt"
	"os"

	mnt "github.com/moby/sys/mountinfo"
	term "github.com/mt1976/crt"
	lang "github.com/mt1976/mockterm/language"
	supt "github.com/mt1976/mockterm/support"
	mem "github.com/shirou/gopsutil/mem"
	cpu "github.com/shirou/gopsutil/v3/cpu"
	dsk "github.com/shirou/gopsutil/v3/disk"
	hst "github.com/shirou/gopsutil/v3/host"
)

// Catalogs the resources available on a given system.

var debugMode bool = false
var results = []string{}

func Run(t *term.ViewPort) {

	p := t.NewPage("Cataloging System Resources")

	debugMode = false
	hostname := t.Helpers.GetHostName()
	outputFilename := "data/catalogs/catalog_" + hostname + "_" + supt.GetTimeStamp() + ".info"
	path, _ := os.Getwd()
	info := New()

	p.AddFieldValuePair("Hostname", t.Helpers.GetHostName())
	storeData(p, info, "Hostname", t.Helpers.GetHostName())

	p.AddFieldValuePair("Machine Name", t.Helpers.GetSytemInfo())
	storeData(p, info, "Machine Name", t.Helpers.GetSytemInfo())

	p.AddFieldValuePair("Username", t.Helpers.GetUsername())
	storeData(p, info, "Username", t.Helpers.GetUsername())

	p.AddFieldValuePair("Current Path", path)
	storeData(p, info, "Current Path", path)
	p.AddBlankRow()

	p.AddFieldValuePair("Output file", outputFilename)

	ok, err := p.Display_Confirmation("Do you want to continue with the cataloging process")
	if err != nil {
		p.Error(err, "unable to get user response")
	}
	if !ok {
		//fmt.Printf("%s Exiting\n", PFY)
		p.Info(lang.TxtQuittingMessage)
		return
	}
	m, _ := mem.VirtualMemory()
	storeData(p, info, "Total Memory", t.Formatters.Human(m.Total))
	storeData(p, info, "Available Memory", t.Formatters.Human(m.Available))
	storeData(p, info, "Used Memory", t.Formatters.Human(m.Used))
	storeData(p, info, "Used Percent", fmt.Sprintf("%f", m.UsedPercent))
	storeData(p, info, "Free Memory", t.Formatters.Human(m.Free))
	cpuData, _ := cpu.Info()
	//range through each cpu and store the cpu details
	p.AddBlankRow()
	p.AddColumnsTitle("CPU", "Model", "Cores", "Mhz", "Model Name")
	for _, cpu := range cpuData {
		t.Break()
		storeData(p, info, fmt.Sprintf("CPU %d", cpu.CPU), t.Formatters.Human(cpu.CPU))
		storeData(p, info, fmt.Sprintf("CPU %d Model", cpu.CPU), t.Formatters.Human(cpu.ModelName))
		storeData(p, info, fmt.Sprintf("CPU %d Cores", cpu.CPU), t.Formatters.Human(cpu.Cores))
		storeData(p, info, fmt.Sprintf("CPU %d Mhz", cpu.CPU), t.Formatters.Human(cpu.Mhz))
		storeData(p, info, fmt.Sprintf("CPU %d Cache Size", cpu.CPU), t.Formatters.Human(cpu.CacheSize))
		storeData(p, info, fmt.Sprintf("CPU %d Flags", cpu.CPU), t.Formatters.Human(cpu.Flags))
		storeData(p, info, fmt.Sprintf("CPU %d Stepping", cpu.CPU), t.Formatters.Human(cpu.Stepping))
		storeData(p, info, fmt.Sprintf("CPU %d Vendor ID", cpu.CPU), t.Formatters.Human(cpu.VendorID))
		storeData(p, info, fmt.Sprintf("CPU %d Family", cpu.CPU), t.Formatters.Human(cpu.Family))
		storeData(p, info, fmt.Sprintf("CPU %d Model", cpu.CPU), t.Formatters.Human(cpu.Model))
		storeData(p, info, fmt.Sprintf("CPU %d Physical ID", cpu.CPU), t.Formatters.Human(cpu.PhysicalID))
		storeData(p, info, fmt.Sprintf("CPU %d Core ID", cpu.CPU), t.Formatters.Human(cpu.CoreID))
		storeData(p, info, fmt.Sprintf("CPU %d Microcode", cpu.CPU), t.Formatters.Human(cpu.Microcode))
		storeData(p, info, fmt.Sprintf("CPU %d Model Name", cpu.CPU), t.Formatters.Human(cpu.ModelName))
		p.AddColumns(t.Formatters.Human(cpu.CPU), t.Formatters.Human(cpu.ModelName), t.Formatters.Human(cpu.Cores), t.Formatters.Human(cpu.Mhz), t.Formatters.Human(cpu.ModelName))
	}
	p.AddBlankRow()

	ht, _ := hst.Info()
	//range through each host and print host info
	t.Break()
	p.AddColumnsTitle("ID", "Hostname", "Uptime", "Procs", "OS-Platform", "Platform Family", "Platform Version", "Kernel Version")
	storeData(p, info, fmt.Sprintf("Host %s", "ID"), t.Formatters.Human(ht.HostID))
	storeData(p, info, fmt.Sprintf("Host %s Hostname", ""), t.Formatters.Human(ht.Hostname))
	storeData(p, info, fmt.Sprintf("Host %s Uptime", ""), t.Formatters.Human(ht.Uptime))
	storeData(p, info, fmt.Sprintf("Host %s Boot Time", ""), t.Formatters.Human(ht.BootTime))
	storeData(p, info, fmt.Sprintf("Host %s Procs", ""), t.Formatters.Human(ht.Procs))
	storeData(p, info, fmt.Sprintf("Host %s OS", ""), t.Formatters.Human(ht.OS))
	storeData(p, info, fmt.Sprintf("Host %s Platform", ""), t.Formatters.Human(ht.Platform))
	storeData(p, info, fmt.Sprintf("Host %s Platform Family", ""), t.Formatters.Human(ht.PlatformFamily))
	storeData(p, info, fmt.Sprintf("Host %s Platform Version", ""), t.Formatters.Human(ht.PlatformVersion))
	storeData(p, info, fmt.Sprintf("Host %s Kernel Version", ""), t.Formatters.Human(ht.KernelVersion))
	p.AddColumns(t.Formatters.Human(ht.HostID), t.Formatters.Human(ht.Hostname), t.Formatters.Human(ht.Uptime), t.Formatters.Human(ht.Procs), t.Formatters.Human(ht.OS)+" "+t.Formatters.Human(ht.Platform), t.Formatters.Human(ht.PlatformFamily), t.Formatters.Human(ht.PlatformVersion), t.Formatters.Human(ht.KernelVersion))
	storeData(p, info, fmt.Sprintf("Host %s Virtualization System", ""), t.Formatters.Human(ht.VirtualizationSystem))
	storeData(p, info, fmt.Sprintf("Host %s Virtualization Role", ""), t.Formatters.Human(ht.VirtualizationRole))
	//p,info = storeData(p, info,p,  "Host Info", T.Human( h))

	v, _ := mnt.GetMounts(nil)
	//zz := 0
	p.AddBlankRow()
	p.AddColumnsTitle("Mountpoint", "FSType", "Source", "Free", "Used", "UsedPercent")
	for zz, v := range v {
		//p,info = storeData(p, info,p,  fmt.Sprintf("Mount %d", zz), T.Human( v))
		t.Break()
		storeData(p, info, fmt.Sprintf("Mount %d ID", zz), t.Formatters.Human(v.ID))
		storeData(p, info, fmt.Sprintf("Mount %d Major", zz), t.Formatters.Human(v.Major))
		storeData(p, info, fmt.Sprintf("Mount %d Minor", zz), t.Formatters.Human(v.Minor))
		storeData(p, info, fmt.Sprintf("Mount %d Root", zz), t.Formatters.Human(v.Root))
		storeData(p, info, fmt.Sprintf("Mount %d Parent", zz), t.Formatters.Human(v.Parent))
		storeData(p, info, fmt.Sprintf("Mount %d Mountpoint", zz), t.Formatters.Human(v.Mountpoint))
		storeData(p, info, fmt.Sprintf("Mount %d Options", zz), t.Formatters.Human(v.Options))
		storeData(p, info, fmt.Sprintf("Mount %d Optional", zz), t.Formatters.Human(v.Optional))
		storeData(p, info, fmt.Sprintf("Mount %d FSType", zz), t.Formatters.Human(v.FSType))
		storeData(p, info, fmt.Sprintf("Mount %d Source", zz), t.Formatters.Human(v.Source))
		storeData(p, info, fmt.Sprintf("Mount %d VFSOptions", zz), t.Formatters.Human(v.VFSOptions))

		usage, _ := dsk.Usage(v.Mountpoint)
		//p,info = storeData(p, info,p,  fmt.Sprintf("Disk Usage %d", zz), T.Human( usage))
		storeData(p, info, fmt.Sprintf("Mount %d Total", zz), t.Formatters.Human(usage.Total))
		storeData(p, info, fmt.Sprintf("Mount %d Free", zz), t.Formatters.Human(usage.Free))
		storeData(p, info, fmt.Sprintf("Mount %d Used", zz), t.Formatters.Human(usage.Used))
		storeData(p, info, fmt.Sprintf("Mount %d UsedPercent", zz), t.Formatters.Human(usage.UsedPercent))
		storeData(p, info, fmt.Sprintf("Mount %d InodesTotal", zz), t.Formatters.Human(usage.InodesTotal))
		storeData(p, info, fmt.Sprintf("Mount %d InodesUsed", zz), t.Formatters.Human(usage.InodesUsed))
		storeData(p, info, fmt.Sprintf("Mount %d InodesFree", zz), t.Formatters.Human(usage.InodesFree))
		storeData(p, info, fmt.Sprintf("Mount %d InodesUsedPercent", zz), t.Formatters.Human(usage.InodesUsedPercent))
		p.AddColumns(t.Formatters.Human(v.Mountpoint), t.Formatters.Human(v.FSType), t.Formatters.Human(v.Source), t.Formatters.Human(usage.Free), t.Formatters.Human(usage.Used), t.Formatters.Human(usage.UsedPercent))
	}
	p.Success("Cataloging complete")
	p.Add("Results have been saved to file: "+outputFilename, "", "")
	debugMode = false
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
	p.ResetPrompt()
	p.Display_Actions()
}
