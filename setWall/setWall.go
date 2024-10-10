package setwall

import (
	"os/exec"
)

func Setwall(filePath string) {
	cmdPath := "file://" + filePath
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", cmdPath)

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
