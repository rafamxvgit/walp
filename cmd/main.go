package main

import (
	"os"
	iomem "walp/ioMem"
	setwall "walp/setWall"
	"walp/utils"
)

func main() {
	programMem := iomem.GetProgMem()

	col1 := programMem.Coleções[1]
	col1.NextImage()
	programMem.Coleções[1] = col1

	iomem.WriteProgramMem(programMem)
	programMem = iomem.GetProgMem()

	col := programMem.Coleções[programMem.SelectedCol]
	dir := utils.Expect(os.ReadDir(col.Path))
	fileName := dir[col.CurrentFile].Name()
	filePath := col.Path + "/" + fileName
	setwall.Setwall(filePath)
}
