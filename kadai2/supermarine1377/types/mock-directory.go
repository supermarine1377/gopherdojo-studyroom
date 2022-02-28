package types

import "supermarine1377/myio"

type MockDirectory struct {
	Name string
}

func (md MockDirectory) Read() ([]myio.Myimage, error) {
	mock := []myio.Myimage{
		{FileName: "test.jpg", Path: "./images/", Extension: "jpg"},
		{FileName: "test.png", Path: "./images/", Extension: "png"},
	}
	return mock, nil
}

func (md MockDirectory) Write([]myio.Myimage) error {
	return nil
}
