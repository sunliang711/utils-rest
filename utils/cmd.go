package utils

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func Run(command string, input []byte) (output []byte, err error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	if input != nil {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		stdin.Write(input)
		stdin.Close()
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// stderr, err := cmd.StderrPipe()
	// if err != nil {
	// 	return nil, err
	// }

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	output, err = ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}
	// errorOutput, err = ioutil.ReadAll(stderr)
	// if err != nil {
	// 	return nil,  err
	// }

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return
}

func Run2(command string, input []byte) (output []byte, err error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	if input != nil {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		stdin.Write(input)
		stdin.Close()
	}

	return cmd.CombinedOutput()

}

func GetSolcVersion(cmd string) (string, error) {
	output, err := Run(cmd, nil)
	if err != nil {
		return "", err
	}
	splits := strings.Split(string(output), "Version: ")
	if len(splits) < 2 {
		return "", errors.New("can not find version info")
	}
	return strings.TrimSpace(splits[1]), nil
}
