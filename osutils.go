package main

import (
	"fmt"
	"log"
	"os/exec"
)

func commandRunOsUtil(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	log.Println("runOsCommands:", cmd)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("os: %v ,failed: %v\n %s", args, err, result)
	}

	return string(result), nil
}

func infoOsUtil() (string, error) {
	args := []string{"-a"}
	return commandRunOsUtil("uname", args...)
}

func memInfoOsUtil() (string, error) {
	args := []string{"-h", "hw.physmem", "hw.usermem"}
	return commandRunOsUtil("sysctl", args...)
}
