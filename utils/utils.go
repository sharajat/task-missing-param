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

// ExtractBootDetails extracts json data of boot
//func ExtractBootDetails() structure.Boot {
//
//	jsonFile, err := os.Open("introspectedData.json")
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer jsonFile.Close()
//	jsonString, _ := ioutil.ReadAll(jsonFile)
//	result := structure.Data{}
//	json.Unmarshal([]byte(jsonString), &result)
//	return result.HardwareDetails.Boot
//
//}

// ExtractCPUDetails extracts json data of cpu
//func ExtractCPUDetails() structure.CPU {
//
//	jsonFile, err := os.Open("introspectedData.json")
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer jsonFile.Close()
//	jsonString, _ := ioutil.ReadAll(jsonFile)
//	result := structure.Data{}
//	json.Unmarshal([]byte(jsonString), &result)
//	return result.HardwareDetails.CPU
//
//}
//var som result.Data

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

func GetCPUDetails(cpudata result.CPUType) structure.CPU {
	var freq float64
	fmt.Sscanf(cpudata.Frequency, "%f", &freq)
	sort.Strings(cpudata.Flags)
	cpu := structure.CPU{
		Arch:           cpudata.Architecture,
		Model:          cpudata.ModelName,
		ClockMegahertz: structure.ClockSpeed(freq) * structure.MegaHertz,
		Count:          cpudata.Count,
		Flags:          cpudata.Flags,
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

//func getNumaNic(id int, numa []result.NumaNICS) structure.NIC {
//	var nicNames []string
//	numanic := structure.NIC{}
//	for _, nic := range numa{
//		if nic.NumaNodeID == id {
//			nicNames = append(nicNames, nic.Name)
//		}
//	}
//	numanic.NumaNic = nicNames
//	fmt.Println(numanic.NumaNic)
//	return numanic
//}
//
//func getNumacpu(id int, numa []result.NumaCPU) structure.CPU {
//	var cpuSiblings [][]int
//	numacpu := structure.CPU{}
//	for _, cpu := range numa{
//		if cpu.NumaNodeID == id {
//			cpuSiblings = append(cpuSiblings, cpu.ThreadSiblings)
//		}
//	}
//	numacpu.ThreadSiblings = cpuSiblings
//	fmt.Println(numacpu.ThreadSiblings)
//	return numacpu
//}

//func GetNumaDetails(numa result.NumaTopology) {
//	//numaNode := new(structure.NumaNodes)
//	numaNode := structure.NumaNodes{}
//	//numanic := structure.NIC{}
//	for _ , value := range numa.NumaRAM{
//		id := value.NumaNodeID
//		numaNode.NumaNodeId = id
//		var nicNames []string
//		var cpuSiblings [][]int
//		//getNumaNic(id,numa.NumaNics)
//		//getNumacpu(id,numa.NumaCPU)
//		for _, nic := range numa.NumaNics {
//			if nic.NumaNodeID == id {
//				nicNames = append(nicNames, nic.Name)
//			}
//		}
//		for _, cpu := range numa.NumaCPU {
//			if cpu.NumaNodeID == id {
//				cpuSiblings = append(cpuSiblings, cpu.ThreadSiblings)
//			}
//		}
//
//		numaNode.NumaNodeDetails.CPU.ThreadSiblings = cpuSiblings
//		//numanic.NumaNic = nicNames
//		//for _, value := range numaNode.NumaNodeDetails.NIC{
//		//	value.NumaNic = nicNames
//		//}
//
//		fmt.Printf("NUMANODE: %+v", numaNode)
//		fmt.Println("")
//
//		//NodeIDs = append(NodeIDs,value.NumaNodeID)
//	}
//}

//func getnumanics(nics []result.NumaNICS) []structure.NumaNICS {
//	nnics := make([]structure.NumaNICS, len(nics))
//	for i, value := range nics {
//		nnics[i] = structure.NumaNICS{
//			NumaNodeID: value.NumaNodeID,
//			Name:       value.Name,
//		}
//	}
//	return nnics
//}

//func getnumaram(ram []result.NumaRAM) []structure.NumaRAM {
//	nram := make([]structure.NumaRAM, len(ram))
//	for i, value := range ram {
//		nram[i] = structure.NumaRAM{
//			NumaNodeID: value.NumaNodeID,
//			SizeKb:     value.SizeKb,
//		}
//	}
//	return nram
//}
//
//func getnumacpu(cpu []result.NumaCPU) []structure.NumaCPU {
//	ncpu := make([]structure.NumaCPU, len(cpu))
//	for i, value := range cpu {
//		ncpu[i] = structure.NumaCPU{
//			NumaNodeID:     value.NumaNodeID,
//			CPUID:          value.CPUID,
//			ThreadSiblings: value.ThreadSiblings,
//		}
//	}
//	return ncpu
//}
//
func GetNumaDetails(numa result.NumaTopology) []structure.NumaTopology {
	numaNode := make([]structure.NumaTopology, len(numa.NumaRAM))
	for i, value := range numa.NumaRAM {
		id := value.NumaNodeID
		ram := value.SizeKb
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
