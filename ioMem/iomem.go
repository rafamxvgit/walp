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
	SelectedCol int
	Coleções    []Collection
}

// Lê a memoria e retorna a memória do programa
func GetProgMem() ProgramMemory {
	file := openMemFile()
	defer file.Close()
	fileData := utils.Expect(io.ReadAll(file))

	var ProgramMemBuffer ProgramMemory
	json.Unmarshal(fileData, &ProgramMemBuffer)
	return ProgramMemBuffer
}

// escreve a memória do programa
func WriteProgramMem(progMem ProgramMemory) {
	file := createMemFile()
	defer file.Close()
	data := utils.Expect(json.Marshal(progMem))
	file.Write(data)
}

// cria ou recria um  novo arquivo mem.json
func createMemFile() *os.File {
	println("try create " + progMemName)
	file, err := os.Create(programMemoryPath())
	if err == nil {
		println("created " + progMemName)
		return file
	}
	panic(err)
}

// abre o arquivo de memória do programa
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

// creates an empty memory file
func initializeMemFile(file *os.File) {
	newProgramMem := ProgramMemory{}
	data := utils.Expect(json.Marshal(newProgramMem))
	file.Write(data)
}

// Retorna o caminho do diretório do executável do programa
func programMemoryPath() string {
	exe := utils.Expect(os.Executable())
	exeDir := path.Dir(exe)
	progMemPath := exeDir + "/" + progMemName
	return progMemPath
}

// this is just a debug function
func InitMemFile() {
	file := createMemFile()
	defer file.Close()

	initialProgMemValue := ProgramMemory{}
	coll1 := Collection{
		Nome: "light",
		Path: "/home/rmxv/tudo/images/wallpapers/light",
	}
	coll2 := Collection{
		Nome: "dark",
		Path: "/home/rmxv/tudo/images/wallpapers/dark",
	}
	initialProgMemValue.Coleções = append(initialProgMemValue.Coleções, coll1, coll2)

	initialProgMemValue.SelectedCol = 1

	data := utils.Expect(json.Marshal(initialProgMemValue))
	utils.Expect(file.Write(data))
}

// seleciona uma nova imagem para a coleção
func (col *Collection) NextImage() {
	dir := utils.Expect(os.ReadDir(col.Path))
	col.CurrentFile++
	if col.CurrentFile >= len(dir) {
		col.CurrentFile = 0
	}
}
