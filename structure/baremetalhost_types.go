package structure

// ClockSpeed is a clock speed in MHz
type ClockSpeed float64

// ClockSpeed multipliers
const (
	MegaHertz ClockSpeed = 1.0
	GigaHertz            = 1000 * MegaHertz
)

// Capacity is a disk size in Bytes
type Capacity int64

// Capacity multipliers
const (
	Byte     Capacity = 1
	KibiByte          = Byte * 1024
	KiloByte          = Byte * 1000
	MebiByte          = KibiByte * 1024
	MegaByte          = KiloByte * 1000
	GibiByte          = MebiByte * 1024
	GigaByte          = MegaByte * 1000
	TebiByte          = GibiByte * 1024
	TeraByte          = GigaByte * 1000
)

// CPU describes one processor on the host.
type CPU struct {
	Arch           string         `json:"arch"`
	Model          string         `json:"model"`
	ClockMegahertz ClockSpeed     `json:"clockMegahertz"`
	Flags          []string       `json:"flags"`
	Count          int            `json:"count"`
	NumaTopology   []NumaTopology `json:"numa_topology"`
}

// Storage describes one storage device (disk, SSD, etc.) on the host.
type Storage struct {
	// A name for the disk, e.g. "disk 1 (boot)"
	Name string `json:"name"`

	// Whether this disk represents rotational storage
	Rotational bool `json:"rotational"`

	// The size of the disk in Bytes
	SizeBytes Capacity `json:"sizeBytes"`

	// The name of the vendor of the device
	Vendor string `json:"vendor,omitempty"`

	// Hardware model
	Model string `json:"model,omitempty"`

	// The serial number of the device
	SerialNumber string `json:"serialNumber"`

	// The WWN of the device
	WWN string `json:"wwn,omitempty"`

	// The WWN Vendor extension of the device
	WWNVendorExtension string `json:"wwnVendorExtension,omitempty"`

	// The WWN with the extension
	WWNWithExtension string `json:"wwnWithExtension,omitempty"`

	// The SCSI location of the device
	HCTL string `json:"hctl,omitempty"`
}

// VLANID is a 12-bit 802.1Q VLAN identifier
type VLANID int32

// VLAN represents the name and ID of a VLAN
type VLAN struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4094
	ID VLANID `json:"id"`

	Name string `json:"name,omitempty"`
}

// NIC describes one network interface on the host.
type NIC struct {
	// The name of the NIC, e.g. "nic-1"
	Name string `json:"name"`

	// The name of the model, e.g. "virt-io"
	Model string `json:"model"`

	// The device MAC addr
	// +kubebuilder:validation:Pattern=`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`
	MAC string `json:"mac"`

	// The IP address of the device
	IP string `json:"ip"`

	// The speed of the device
	SpeedGbps int `json:"speedGbps"`

	// The VLANs available
	VLANs []VLAN `json:"vlans,omitempty"`

	// The untagged VLAN ID
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4094
	VLANID VLANID `json:"vlanId"`

	// Whether the NIC is PXE Bootable
	PXE bool `json:"pxe"`
}

// Firmware describes the firmware on the host.
type Firmware struct {
	// The BIOS for this firmware
	BIOS BIOS `json:"bios"`
}

// BIOS describes the BIOS version on the host.
type BIOS struct {
	// The release/build date for this BIOS
	Date string `json:"date"`

	// The vendor name for this BIOS
	Vendor string `json:"vendor"`

	// The version of the BIOS
	Version string `json:"version"`
}

// HardwareSystemVendor stores details about the whole hardware system.
type HardwareSystemVendor struct {
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"productName"`
	SerialNumber string `json:"serialNumber"`
}

// HardwareDetails collects all of the information about hardware
// discovered on the host.
type HardwareDetails struct {
	SystemVendor    HardwareSystemVendor `json:"systemVendor"`
	Firmware        Firmware             `json:"firmware"`
	RAMMebibytes    int                  `json:"ramMebibytes"`
	NIC             []NIC                `json:"nics"`
	Storage         []Storage            `json:"storage"`
	CPU             CPU                  `json:"cpu"`
	Hostname        string               `json:"hostname"`
	CurrentBootMode string               `json:"currentbootmode"`
	//NumaTopology   []NumaTopology   `json:"numa_topology"`
	//NodeDetails  NodeDetails
	//NumaNodes    []NumaNodes
}
type NumaTopology struct {
	NumaNodeID  int         `json:"numa_node"`
	Numadetails Numadetails `json:"numa_topology"`
}

type Numadetails struct {
	Nics           []string `json:"nics"`
	RAM            int      `json:"ram"`
	ThreadSiblings [][]int  `json:"cpus"`
}
type NumaNICS struct {
	Name []string `json:"name"`
}
type NumaRAM struct {
	SizeKb int `json:"size_kb"`
}
type NumaCPU struct {
	ThreadSiblings [][]int `json:"thread_siblings"`
}

//type NodeDetails struct {
//	CPU          CPU    `json:"cpu"`
//	NIC          []NIC  	`json:"nics"`
//	//NumaNics	[]string
//	RAMMebibytes int    `json:"ramMebibytes"`
//}
//
//type NumaNodes struct {
//	NumaNodeId      int
//	NumaNodeDetails NodeDetails
//}
