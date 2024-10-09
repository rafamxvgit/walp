package iomem

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"walp/utils"
)

const progMemName string = "mem.json"

type Collection struct {
	Nome        string //o nome da coleção
	Path        string //o caminho da coleção
	CurrentFile int    //o indice da última imagem mostrada
}

type ProgramMemory struct {
	Coleções []Collection
}

func ReadProgMem() ProgramMemory {
	file := openMemFile()
	defer file.Close()
	fileData := utils.Expect(io.ReadAll(file))

	var ProgramMemBuffer ProgramMemory
	json.Unmarshal(fileData, &ProgramMemBuffer)
	return ProgramMemBuffer
}

func openMemFile() *os.File {
	println("try open " + progMemName)
	file, err := os.Open(programMemoryPath())
	if err == nil {
		println("opened " + progMemName)
		return file
	}
	println("open " + progMemName + " failed")
	println("try create " + progMemName)

	file, err = os.Create(programMemoryPath())

	if err == nil {
		println("created " + progMemName)
		println("initializing " + progMemName)
		initializeMemFile(file)
		file.Close()
		return openMemFile()
	}

	panic("could not open " + progMemName)
}

func initializeMemFile(file *os.File) {
	newProgramMem := ProgramMemory{}
	data := utils.Expect(json.Marshal(newProgramMem))
	file.Write(data)
}

/*
Retorna o caminho do diretório do executável do programa
*/
func programMemoryPath() string {
	exe := utils.Expect(os.Executable())
	exeDir := path.Dir(exe)
	progMemPath := exeDir + "/" + progMemName
	return progMemPath
}
