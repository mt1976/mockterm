package catalog

import (
	"fmt"
	"os"

	mnt "github.com/moby/sys/mountinfo"
	term "github.com/mt1976/crt"
	errs "github.com/mt1976/mockterm/errors"
	supt "github.com/mt1976/mockterm/support"
	mem "github.com/shirou/gopsutil/mem"
	cpu "github.com/shirou/gopsutil/v3/cpu"
	dsk "github.com/shirou/gopsutil/v3/disk"
	hst "github.com/shirou/gopsutil/v3/host"
)

// Catalogs the resources available on a given system.

var debugMode bool = false
var crt term.Crt

type info struct {
	data []string
}

func Run(crtIn term.Crt, debug bool, path string) {

	debugMode = debug
	crt = crtIn
	hostname := crtIn.Helpers.GetHostName()
	outputFilename := "catalog_" + hostname + "_" + supt.GetTimeStamp() + ".info"

	info := NewInfo()
	//X := T

	info.breakData("Cataloging system resources")
	//fmt.Println(crt.PR("Output file = "+outputFilename, T))
	crt.Print("Output file : " + outputFilename)
	crt.Break()
	info.storeData("Hostname", crt.Helpers.GetHostName())
	info.storeData("Machine Name", crt.Helpers.GetSytemInfo())
	info.storeData("Username", crt.Helpers.GetUsername())
	info.storeData("Current Path", path)

	m, _ := mem.VirtualMemory()
	info.storeData("Total Memory", crt.Formatters.Human(m.Total))
	info.storeData("Available Memory", crt.Formatters.Human(m.Available))
	info.storeData("Used Memory", crt.Formatters.Human(m.Used))
	info.storeData("Used Percent", fmt.Sprintf("%f", m.UsedPercent))
	info.storeData("Free Memory", crt.Formatters.Human(m.Free))
	c, _ := cpu.Info()
	//range through each cpu and store the cpu details

	for _, cpu := range c {
		crt.Break()
		info.storeData(fmt.Sprintf("CPU %d", cpu.CPU), crt.Formatters.Human(cpu.CPU))
		info.storeData(fmt.Sprintf("CPU %d Model", cpu.CPU), crt.Formatters.Human(cpu.ModelName))
		info.storeData(fmt.Sprintf("CPU %d Cores", cpu.CPU), crt.Formatters.Human(cpu.Cores))
		info.storeData(fmt.Sprintf("CPU %d Mhz", cpu.CPU), crt.Formatters.Human(cpu.Mhz))
		info.storeData(fmt.Sprintf("CPU %d Cache Size", cpu.CPU), crt.Formatters.Human(cpu.CacheSize))
		info.storeData(fmt.Sprintf("CPU %d Flags", cpu.CPU), crt.Formatters.Human(cpu.Flags))
		info.storeData(fmt.Sprintf("CPU %d Stepping", cpu.CPU), crt.Formatters.Human(cpu.Stepping))
		info.storeData(fmt.Sprintf("CPU %d Vendor ID", cpu.CPU), crt.Formatters.Human(cpu.VendorID))
		info.storeData(fmt.Sprintf("CPU %d Family", cpu.CPU), crt.Formatters.Human(cpu.Family))
		info.storeData(fmt.Sprintf("CPU %d Model", cpu.CPU), crt.Formatters.Human(cpu.Model))
		info.storeData(fmt.Sprintf("CPU %d Physical ID", cpu.CPU), crt.Formatters.Human(cpu.PhysicalID))
		info.storeData(fmt.Sprintf("CPU %d Core ID", cpu.CPU), crt.Formatters.Human(cpu.CoreID))
		info.storeData(fmt.Sprintf("CPU %d Microcode", cpu.CPU), crt.Formatters.Human(cpu.Microcode))
		info.storeData(fmt.Sprintf("CPU %d Model Name", cpu.CPU), crt.Formatters.Human(cpu.ModelName))
	}

	ht, _ := hst.Info()
	//range through each host and print host info
	crt.Break()
	info.storeData(fmt.Sprintf("Host %s", "ID"), crt.Formatters.Human(ht.HostID))
	info.storeData(fmt.Sprintf("Host %s Hostname", ""), crt.Formatters.Human(ht.Hostname))
	info.storeData(fmt.Sprintf("Host %s Uptime", ""), crt.Formatters.Human(ht.Uptime))
	info.storeData(fmt.Sprintf("Host %s Boot Time", ""), crt.Formatters.Human(ht.BootTime))
	info.storeData(fmt.Sprintf("Host %s Procs", ""), crt.Formatters.Human(ht.Procs))
	info.storeData(fmt.Sprintf("Host %s OS", ""), crt.Formatters.Human(ht.OS))
	info.storeData(fmt.Sprintf("Host %s Platform", ""), crt.Formatters.Human(ht.Platform))
	info.storeData(fmt.Sprintf("Host %s Platform Family", ""), crt.Formatters.Human(ht.PlatformFamily))
	info.storeData(fmt.Sprintf("Host %s Platform Version", ""), crt.Formatters.Human(ht.PlatformVersion))
	info.storeData(fmt.Sprintf("Host %s Kernel Version", ""), crt.Formatters.Human(ht.KernelVersion))
	info.storeData(fmt.Sprintf("Host %s Virtualization System", ""), crt.Formatters.Human(ht.VirtualizationSystem))
	info.storeData(fmt.Sprintf("Host %s Virtualization Role", ""), crt.Formatters.Human(ht.VirtualizationRole))
	//info.storeData( T,  "Host Info", T.Human( h))

	v, _ := mnt.GetMounts(nil)
	//zz := 0
	for zz, v := range v {
		//info.storeData( T,  fmt.Sprintf("Mount %d", zz), T.Human( v))
		crt.Break()
		info.storeData(fmt.Sprintf("Mount %d ID", zz), crt.Formatters.Human(v.ID))
		info.storeData(fmt.Sprintf("Mount %d Major", zz), crt.Formatters.Human(v.Major))
		info.storeData(fmt.Sprintf("Mount %d Minor", zz), crt.Formatters.Human(v.Minor))
		info.storeData(fmt.Sprintf("Mount %d Root", zz), crt.Formatters.Human(v.Root))
		info.storeData(fmt.Sprintf("Mount %d Parent", zz), crt.Formatters.Human(v.Parent))
		info.storeData(fmt.Sprintf("Mount %d Mountpoint", zz), crt.Formatters.Human(v.Mountpoint))
		info.storeData(fmt.Sprintf("Mount %d Options", zz), crt.Formatters.Human(v.Options))
		info.storeData(fmt.Sprintf("Mount %d Optional", zz), crt.Formatters.Human(v.Optional))
		info.storeData(fmt.Sprintf("Mount %d FSType", zz), crt.Formatters.Human(v.FSType))
		info.storeData(fmt.Sprintf("Mount %d Source", zz), crt.Formatters.Human(v.Source))
		info.storeData(fmt.Sprintf("Mount %d VFSOptions", zz), crt.Formatters.Human(v.VFSOptions))

		usage, _ := dsk.Usage(v.Mountpoint)
		//info.storeData( T,  fmt.Sprintf("Disk Usage %d", zz), T.Human( usage))
		info.storeData(fmt.Sprintf("Mount %d Total", zz), crt.Formatters.Human(usage.Total))
		info.storeData(fmt.Sprintf("Mount %d Free", zz), crt.Formatters.Human(usage.Free))
		info.storeData(fmt.Sprintf("Mount %d Used", zz), crt.Formatters.Human(usage.Used))
		info.storeData(fmt.Sprintf("Mount %d UsedPercent", zz), crt.Formatters.Human(usage.UsedPercent))
		info.storeData(fmt.Sprintf("Mount %d InodesTotal", zz), crt.Formatters.Human(usage.InodesTotal))
		info.storeData(fmt.Sprintf("Mount %d InodesUsed", zz), crt.Formatters.Human(usage.InodesUsed))
		info.storeData(fmt.Sprintf("Mount %d InodesFree", zz), crt.Formatters.Human(usage.InodesFree))
		info.storeData(fmt.Sprintf("Mount %d InodesUsedPercent", zz), crt.Formatters.Human(usage.InodesUsedPercent))

	}
	if !debugMode {
		// Open output file
		file, err := openFile(outputFilename)
		if err != nil {
			return
		}
		defer file.Close()
		err = writeStringSliceToFile(file, info.data)
		if err != nil {
			return
		}
	}
}

func (i *info) breakData(title string) {
	i.data = append(i.data, crt.Format("", ""))
	i.data = append(i.data, crt.Format(title, ""))
	i.data = append(i.data, crt.Format("", ""))
	crt.Print(crt.Formatters.Bold(title))
	crt.Break()
}

func NewInfo() info {
	return info{}
}

func (i *info) storeData(title string, data string) {

	// Padd title to 15 characters
	title = fmt.Sprintf("%-30s", title)

	i.data = append(i.data, crt.Format(title+": "+data, ""))
	//fmt.Println(support.PR(title+": "+support.BOLD+data+support.RESET, T))
	crt.Print(title + ": " + crt.Formatters.Bold(data))
}

func openFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//fmt.Printf("%s Error opening file %s: %v\n", crt.CHnormal, filename, err)
		crt.Error(errs.ErrOpeningFile, crt.Formatters.Bold(filename), err.Error())
		return nil, err
	}
	return file, nil
}

func writeStringSliceToFile(file *os.File, info []string) error {
	for _, line := range info {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			//fmt.Printf("%s Error writing to file %s: %v\n", crt.CHnormal, file.Name(), err)
			crt.Error(errs.ErrWritingFile, crt.Formatters.Bold(file.Name()), err.Error())
			return err
		}
	}
	return nil
}
