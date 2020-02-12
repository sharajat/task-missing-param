package result

// Data represents the full introspection data collected.
// The format and contents of the stored data depends on the ramdisk used
// and plugins enabled both in the ramdisk and in inspector itself.
// This structure has been provided for basic compatibility but it
// will need extensions
type Data struct {
	AllInterfaces map[string]BaseInterfaceType `json:"all_interfaces"`
	BootInterface string                       `json:"boot_interface"`
	CPUArch       string                       `json:"cpu_arch"`
	CPUs          int                          `json:"cpus"`
	Error         string                       `json:"error"`
	Interfaces    map[string]BaseInterfaceType `json:"interfaces"`
	Inventory     InventoryType                `json:"inventory"`
	IPMIAddress   string                       `json:"ipmi_address"`
	LocalGB       int                          `json:"local_gb"`
	MACs          []string                     `json:"macs"`
	MemoryMB      int                          `json:"memory_mb"`
	RootDisk      RootDiskType                 `json:"root_disk"`
	Extra         ExtraHardwareDataType        `json:"extra"`
	NumaTopology  NumaTopology                 `json:"numa_topology"`
}

// Sub Types defined under Data and deeper in the structure

type BaseInterfaceType struct {
	ClientID      string                 `json:"client_id"`
	IP            string                 `json:"ip"`
	MAC           string                 `json:"mac"`
	PXE           bool                   `json:"pxe"`
	LLDPProcessed map[string]interface{} `json:"lldp_processed"`
}

type BootInfoType struct {
	CurrentBootMode string `json:"current_boot_mode"`
	PXEInterface    string `json:"pxe_interface"`
}

type CPUType struct {
	Architecture string   `json:"architecture"`
	Count        int      `json:"count"`
	Flags        []string `json:"flags"`
	Frequency    string   `json:"frequency"`
	ModelName    string   `json:"model_name"`
}

type LLDPTLVType struct {
	Type  int
	Value string
}

type InterfaceType struct {
	BIOSDevName string        `json:"biosdevname"`
	ClientID    string        `json:"client_id"`
	HasCarrier  bool          `json:"has_carrier"`
	IPV4Address string        `json:"ipv4_address"`
	IPV6Address string        `json:"ipv6_address"`
	LLDP        []LLDPTLVType `json:"lldp"`
	MACAddress  string        `json:"mac_address"`
	Name        string        `json:"name"`
	Product     string        `json:"product"`
	Vendor      string        `json:"vendor"`
}

type InventoryType struct {
	BmcAddress   string           `json:"bmc_address"`
	Boot         BootInfoType     `json:"boot"`
	CPU          CPUType          `json:"cpu"`
	Disks        []RootDiskType   `json:"disks"`
	Interfaces   []InterfaceType  `json:"interfaces"`
	Memory       MemoryType       `json:"memory"`
	SystemVendor SystemVendorType `json:"system_vendor"`
	Hostname     string           `json:"hostname"`
}

type MemoryType struct {
	PhysicalMb int `json:"physical_mb"`
	Total      int `json:"total"`
}

type RootDiskType struct {
	Hctl               string `json:"hctl"`
	Model              string `json:"model"`
	Name               string `json:"name"`
	ByPath             string `json:"by_path"`
	Rotational         bool   `json:"rotational"`
	Serial             string `json:"serial"`
	Size               int    `json:"size"`
	Vendor             string `json:"vendor"`
	Wwn                string `json:"wwn"`
	WwnVendorExtension string `json:"wwn_vendor_extension"`
	WwnWithExtension   string `json:"wwn_with_extension"`
}

type SystemVendorType struct {
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"product_name"`
	SerialNumber string `json:"serial_number"`
}

type ExtraHardwareData map[string]interface{}

type ExtraHardwareDataSection map[string]ExtraHardwareData

type ExtraHardwareDataType struct {
	CPU      ExtraHardwareDataSection `json:"cpu"`
	Disk     ExtraHardwareDataSection `json:"disk"`
	Firmware ExtraHardwareDataSection `json:"firmware"`
	IPMI     ExtraHardwareDataSection `json:"ipmi"`
	Memory   ExtraHardwareDataSection `json:"memory"`
	Network  ExtraHardwareDataSection `json:"network"`
	System   ExtraHardwareDataSection `json:"system"`
}

type NumaTopology struct {
	NumaNics []NumaNICS `json:"nics"`
	NumaRAM  []NumaRAM  `json:"ram"`
	NumaCPU  []NumaCPU  `json:"cpus"`
}
type NumaNICS struct {
	NumaNodeID int    `json:"numa_node"`
	Name       string `json:"name"`
}
type NumaRAM struct {
	NumaNodeID int `json:"numa_node"`
	SizeKb     int `json:"size_kb"`
}
type NumaCPU struct {
	NumaNodeID     int   `json:"numa_node"`
	CPUID          int   `json:"cpu"`
	ThreadSiblings []int `json:"thread_siblings"`
}
