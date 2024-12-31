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
		return errors.New("Hehe nice try! Enter a project name")
	default:
		for _, dir := range dirs {
			if dir.Name() == name && dir.IsDir() {
				return errors.New("Directory already exists!")
			}
		}
	}
	return nil
}
