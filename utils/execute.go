package utils

import (
	"os/exec"
)

func ExecuteVcGen(vcGenCmd, command string, parameters ...string) (string, error) {
	args := make([]string, len(parameters)+1)
	args = append(args, command)

	for _, parameter := range parameters {
		args = append(args, parameter)
	}

	return Execute(vcGenCmd, args...)
}

func Execute(appname string, args ...string) (string, error) {
	command := exec.Command(appname, args...)
	output, err := command.Output()

	return string(output), err
}
