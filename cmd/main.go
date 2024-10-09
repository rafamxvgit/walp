package main

import (
	iomem "walp/ioMem"
)

func main() {
	programMem := iomem.ReadProgMem()
	println(programMem.Coleções)
}
