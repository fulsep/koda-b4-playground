package lib

import (
	"os"
	"os/exec"
	"runtime"
)

func CLS(){
	switch runtime.GOOS{
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}