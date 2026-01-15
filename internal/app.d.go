package app

import "fmt"

type Project struct {
	Name        string
	Description string
	ModName     string
	FrameWork   string
	Database    string
	License     string
	Author      string
	Year        string
}

type AppError struct {
	Message string
}

func (e AppError) Error() string {
	s := fmt.Sprintf(" Error\n󰋽 Message : \n%s", e.Message)
	return s
}
