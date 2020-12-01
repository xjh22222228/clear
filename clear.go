package clear

import (
	"os"
	"os/exec"
	"runtime"
)

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func Clear()  {
	if isWindows() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		os.Stdout.WriteString("\033[H\033[2J")
	}
}
