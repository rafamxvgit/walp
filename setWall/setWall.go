package setwall

import (
	"os/exec"
)

// gsettings set org.gnome.desktop.background picture-uri file:///home/serrano/Pictures/y.jpg
func Setwall(filePath string) {
	cmdPath := "file://" + filePath
	println(cmdPath)
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", cmdPath)

	err := cmd.Run()
	if err != nil {
		panic(err)
	} else {
		print("deu bom")
	}
}
