package main

import (
	"fmt"
	"task-missing-param/utils"
)

func main() {
	fmt.Println("hello world")
	BootDetails := utils.ExtractBootDetails()
	CPUDetails := utils.ExtractCPUDetails()
	fmt.Printf("\nBoot Details-:\nCurrentBootMode: %s\nPxeInterface: %s\n", BootDetails.CurrentBootMode, BootDetails.PxeInterface)
	fmt.Printf("\nCPU Details-:\ncount: %d\nfrequency: %s\nflags:  %s\nmodel_name: %s\narchitecture: %s\n", CPUDetails.Count, CPUDetails.Frequency, CPUDetails.Flags, CPUDetails.ModelName, CPUDetails.Architecture)
	fmt.Println(CPUDetails.NumaTopology)
}
