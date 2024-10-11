package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"
	iomem "walp/ioMem"
	setwall "walp/setWall"
	"walp/utils"
)

func main() {
	programMem := iomem.GetProgMem()

	col1 := programMem.Coleções[1]
	col1.NextImage()
	programMem.Coleções[1] = col1

	col := programMem.Coleções[programMem.SelectedCol]
	dir := utils.Expect(os.ReadDir(col.Path))
	fileName := dir[col.CurrentFile].Name()
	filePath := col.Path + "/" + fileName
	setwall.Setwall(filePath)
	iomem.WriteProgramMem(programMem)
	scheduleNext()
}

func scheduleNext() {
	currentTime := time.Now()
	currentTime = currentTime.Add(time.Minute * 2)
	hour := currentTime.Hour()
	minut := currentTime.Minute()

	hstr := strconv.FormatInt(int64(hour), 10)
	if hour < 10 {
		hstr = "0" + hstr
	}
	mstr := strconv.FormatInt(int64(minut), 10)
	if minut < 10 {
		mstr = "0" + mstr
	}

	newTime := hstr + ":" + mstr

	cmd := exec.Command("at", newTime, "-f", "/home/rmxv/tudo/codes/projetos/walp/cmd/cmd.sh")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
