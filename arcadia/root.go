package arcadia

import (
	"os/exec"
	"strings"
)

func Root() (string, error) {
	out, err := exec.Command("arc", "root").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
