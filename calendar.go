package gdate

import "os/exec"

func Calendar() (string, error) {
	cmd := exec.Command("cal")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
