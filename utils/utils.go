package utils

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"task-missing-param/result"
	"task-missing-param/structure"
	//"fmt"
	"io/ioutil"
	"os"
)

func Gopherextract() result.Data {
	jsonFile, _ := os.Open("introspectedData.json")
	defer jsonFile.Close()
	jsonString, _ := ioutil.ReadAll(jsonFile)
	result := result.Data{}
	json.Unmarshal([]byte(jsonString), &result)
	return result
}

func getVLANs(intf result.BaseInterfaceType) (vlans []structure.VLAN, vlanid structure.VLANID) {
	if intf.LLDPProcessed == nil {
		return
	}
	if spvs, ok := intf.LLDPProcessed["switch_port_vlans"]; ok {
		if data, ok := spvs.([]map[string]interface{}); ok {
			vlans = make([]structure.VLAN, len(data))
			for i, vlan := range data {
				vid, _ := vlan["id"].(int)
				name, _ := vlan["name"].(string)
				vlans[i] = structure.VLAN{
					ID:   structure.VLANID(vid),
					Name: name,
				}
			}
		}
	}
	if vid, ok := intf.LLDPProcessed["switch_port_untagged_vlan_id"].(int); ok {
		vlanid = structure.VLANID(vid)
	}
	return
}

func getNICSpeedGbps(intfExtradata result.ExtraHardwareData) (speedGbps int) {
	if speed, ok := intfExtradata["speed"].(string); ok {
		if strings.HasSuffix(speed, "Gbps") {
			fmt.Sscanf(speed, "%d", &speedGbps)
		}
	}
	return
}

func GetSystemVendorDetails(vendor result.SystemVendorType) structure.HardwareSystemVendor {
	return structure.HardwareSystemVendor{
		Manufacturer: vendor.Manufacturer,
		ProductName:  vendor.ProductName,
		SerialNumber: vendor.SerialNumber,
	}
}

func GetCPUDetails(cpudata result.CPUType, topology result.NumaTopology) structure.CPU {
	var freq float64
	fmt.Sscanf(cpudata.Frequency, "%f", &freq)
	sort.Strings(cpudata.Flags)
	cpu := structure.CPU{
		Arch:           cpudata.Architecture,
		Model:          cpudata.ModelName,
		ClockMegahertz: structure.ClockSpeed(freq) * structure.MegaHertz,
		Count:          cpudata.Count,
		Flags:          cpudata.Flags,
		NumaTopology:   GetNumaDetails(topology),
	}
	return cpu
}

func GetFirmwareDetails(firmwaredata result.ExtraHardwareDataSection) structure.Firmware {

	// handle bios optionally
	var bios structure.BIOS

	if biosdata, ok := firmwaredata["bios"]; ok {
		// we do not know if all fields will be supplied
		// as this is not a structured response
		// so we must handle each field conditionally
		bios.Vendor, _ = biosdata["vendor"].(string)
		bios.Version, _ = biosdata["version"].(string)
		bios.Date, _ = biosdata["date"].(string)
	}

	return structure.Firmware{
		BIOS: bios,
	}

}

func GetNICDetails(ifdata []result.InterfaceType,
	basedata map[string]result.BaseInterfaceType,
	extradata result.ExtraHardwareDataSection) []structure.NIC {
	nics := make([]structure.NIC, len(ifdata))
	for i, intf := range ifdata {
		baseIntf := basedata[intf.Name]
		vlans, vlanid := getVLANs(baseIntf)

		nics[i] = structure.NIC{
			Name: intf.Name,
			Model: strings.TrimLeft(fmt.Sprintf("%s %s",
				intf.Vendor, intf.Product), " "),
			MAC:       intf.MACAddress,
			IP:        intf.IPV4Address,
			VLANs:     vlans,
			VLANID:    vlanid,
			SpeedGbps: getNICSpeedGbps(extradata[intf.Name]),
			PXE:       baseIntf.PXE,
		}
	}
	return nics
}

func GetStorageDetails(diskdata []result.RootDiskType) []structure.Storage {
	storage := make([]structure.Storage, len(diskdata))
	for i, disk := range diskdata {
		storage[i] = structure.Storage{
			Name:               disk.Name,
			Rotational:         disk.Rotational,
			SizeBytes:          structure.Capacity(disk.Size),
			Vendor:             disk.Vendor,
			Model:              disk.Model,
			SerialNumber:       disk.Serial,
			WWN:                disk.Wwn,
			WWNVendorExtension: disk.WwnVendorExtension,
			WWNWithExtension:   disk.WwnWithExtension,
			HCTL:               disk.Hctl,
		}
	}
	return storage
}

func GetNumaDetails(numa result.NumaTopology) []structure.NumaTopology {
	numaNode := make([]structure.NumaTopology, len(numa.NumaRAM))
	for i, value := range numa.NumaRAM {
		id, ram := value.NumaNodeID, value.SizeKb
		var nicNames []string
		var cpuSiblings [][]int
		for _, nic := range numa.NumaNics {
			if nic.NumaNodeID == id {
				nicNames = append(nicNames, nic.Name)
			}
		}
		for _, cpu := range numa.NumaCPU {
			if cpu.NumaNodeID == id {
				cpuSiblings = append(cpuSiblings, cpu.ThreadSiblings)
			}
		}
		numaNode[i].NumaNodeID = id
		numaNode[i].Numadetails.ThreadSiblings = cpuSiblings
		numaNode[i].Numadetails.Nics = nicNames
		numaNode[i].Numadetails.RAM = ram
	}
	return numaNode

}
