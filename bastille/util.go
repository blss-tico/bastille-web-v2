package bastille

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os/exec"
	"strconv"
)

func CommandRunOsUtil(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	log.Println("runOsCommands:", cmd)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("os: %v ,failed: %v\n %s", args, err, result)
	}

	return string(result), nil
}

func InfoOsUtil() (string, error) {
	args := []string{"-a"}
	return CommandRunOsUtil("uname", args...)
}

func MemInfoOsUtil() (string, error) {
	args := []string{"-h", "hw.physmem", "hw.usermem"}
	return CommandRunOsUtil("sysctl", args...)
}

func RandPortUtil() string {
	log.Println("RandPortUtil")
	return strconv.Itoa(8000 + rand.IntN(8200-8000))
}
