//go:build linux

/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  05/04/22, 5:55 PM
 */

package open

import (
	"os/exec"
)

// http://sources.debian.net/src/xdg-utils

func open(input string) *exec.Cmd {
	return exec.Command("xdg-open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command(appName, input)
}
