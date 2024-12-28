package utils

import (
	"errors"
	"os"
)

func CheckValidProjectName(name string) error {
	dirs, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	switch name {
	case "":
		return errors.New("No null name, sorry")
	default:
		for _, dir := range dirs {
			if dir.Name() == name && dir.IsDir() {
				return errors.New("Project already exists")
			}
		}
	}
	return nil
}
