package term

import (
	"os/exec"
	"runtime"
)

// Checks and returns if the command is available
// helpfull gist: https://gist.github.com/miguelmota/ed4ec562b8cd1781e7b20151b37de8a0
func CheckCmd(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// returns if the platform is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
