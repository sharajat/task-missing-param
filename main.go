package main

import (
	"fmt"
	"task-missing-param/result"
	"task-missing-param/utils"
)

func getHardwareDetails(data result.Data) {
	Firmware := utils.GetFirmwareDetails(data.Extra.Firmware)
	SystemVendor := utils.GetSystemVendorDetails(data.Inventory.SystemVendor)
	RAMMebibytes := data.MemoryMB
	NIC := utils.GetNICDetails(data.Inventory.Interfaces, data.AllInterfaces, data.Extra.Network)
	Storage := utils.GetStorageDetails(data.Inventory.Disks)
	Numa := utils.GetNumaDetails(data.NumaTopology)

	CPU := utils.GetCPUDetails(data.Inventory.CPU, Numa)
	Hostname := data.Inventory.Hostname
	//CurrentBootMode := data.Inventory.Boot.CurrentBootMode
	fmt.Println("")
	fmt.Println("")

	fmt.Println("")

	fmt.Println("")

	fmt.Println(Firmware)
	fmt.Println(SystemVendor)
	fmt.Println(RAMMebibytes)
	fmt.Println(NIC)
	fmt.Println(Storage)
	fmt.Printf("%+v", CPU)
	fmt.Println(Hostname)
	fmt.Println(Numa)
}

func main() {
	fmt.Println("hello world")
	introspectiondata := utils.Gopherextract()
	getHardwareDetails(introspectiondata)
}
