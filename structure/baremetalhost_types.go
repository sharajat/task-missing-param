package structure

// Data ironic introspected data
type Data struct {
	HardwareDetails HardwareDetails `json:"HardwareDetails"`
}

// HardwareDetails config details
type HardwareDetails struct {
	Boot         Boot         `json:"boot"`
	SystemVendor SystemVendor `json:"system_vendor"`
	Memory       Memory       `json:"memory"`
	CPU          CPU          `json:"cpu"`
}

// Boot boot details
type Boot struct {
	CurrentBootMode string `json:"current_boot_mode"`
	PxeInterface    string `json:"pxe_interface"`
}

// SystemVendor vendor details
type SystemVendor struct {
	SerialNumber string `json:"serial_number"`
	ProductName  string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
}

// CPU cpu details
type CPU struct {
	Count        int          `json:"count"`
	Frequency    string       `json:"frequency"`
	Flags        []string     `json:"flags"`
	ModelName    string       `json:"model_name"`
	Architecture string       `json:"architecture"`
	NumaTopology NumaTopology `json:"numa_topology"`
}

// NumaTopology structure of all
type NumaTopology struct {
	NNics []NumaNICS `json:"nics"`
	NRam  []NumaRAM  `json:"ram"`
	NCPU  []NumaCPU  `json:"cpus"`
}

// NumaNICS nic details
type NumaNICS struct {
	NumaNodeID int    `json:"numa_node"`
	Name       string `json:"name"`
}

// NumaRAM ram details
type NumaRAM struct {
	NumaNodeID int `json:"numa_node"`
	SizeKb     int `json:"size_kb"`
}

// NumaCPU numacpu details
type NumaCPU struct {
	NumaNodeID     int   `json:"numa_node"`
	CPUID          int   `json:"cpu"`
	ThreadSiblings []int `json:"thread_siblings"`
}

// Memory memory details
type Memory struct {
	PhysicalMb int   `json:"physical_mb"`
	Total      int64 `json:"total"`
}
