//go:build darwin

/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  05/04/22, 5:58 PM
 */

package open

import (
	"os/exec"
)

func open(input string) *exec.Cmd {
	return exec.Command("open", input)
}

func openWith(input string, appName string) *exec.Cmd {
	return exec.Command("open", "-a", appName, input)
}
