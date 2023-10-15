package file

import (
	"fmt"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	fmt.Printf("WARN %s", err)
	return false
}
