package app

import "fmt"

type Project struct {
	Name      string
	ModName   string
	FrameWork string
	Database  string
}

type AppError struct {
	Message string
}

func (e AppError) Error() string {
	s := fmt.Sprintf(" Error\n󰋽 Message : \n%s", e.Message)
	return s
}
