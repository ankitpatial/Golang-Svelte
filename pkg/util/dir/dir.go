package dir

import (
	"errors"
	"os"
)

func Create(dirName string) error {
	err := os.MkdirAll(dirName, 0750)
	if err == nil {
		return nil
	}

	if os.IsExist(err) {
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}

		return nil
	}

	return err
}
