package main

import (
	"encoding/json"
	"fmt"
	"task-missing-param/result"
	"task-missing-param/structure"
	"task-missing-param/utils"
)

func getHardwareDetails(data result.Data) *structure.HardwareDetails {
	details := new(structure.HardwareDetails)
	details.Firmware = utils.GetFirmwareDetails(data.Extra.Firmware)
	details.SystemVendor = utils.GetSystemVendorDetails(data.Inventory.SystemVendor)
	details.RAMMebibytes = data.MemoryMB
	details.NIC = utils.GetNICDetails(data.Inventory.Interfaces, data.AllInterfaces, data.Extra.Network)
	details.Storage = utils.GetStorageDetails(data.Inventory.Disks)
	details.CPU = utils.GetCPUDetails(data.Inventory.CPU, data.NumaTopology)
	details.Hostname = data.Inventory.Hostname
	details.CurrentBootMode = data.Inventory.Boot.CurrentBootMode
	return details
}
func jsonmarshal(data *structure.HardwareDetails) {
	job, _ := json.Marshal(data)
	fmt.Println(string(job))
}

func main() {
	fmt.Println("hello world")
	introspectiondata := utils.Gopherextract()
	rest := getHardwareDetails(introspectiondata)
	jsonmarshal(rest)
}
