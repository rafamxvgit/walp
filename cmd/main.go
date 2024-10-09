package main

import (
	"fmt"
	iomem "walp/ioMem"
)

func main() {
	iomem.InitMemFile()
	programMem := iomem.ReadProgMem()
	for _, col := range programMem.Coleções {
		fmt.Printf("%+v\n", col)
	}
}
