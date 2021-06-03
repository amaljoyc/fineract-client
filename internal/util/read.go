package util

import (
	"embed"
)

//go:embed json/*
var jsondir embed.FS

func Read(filename string) []byte  {
	data, err := jsondir.ReadFile("json/" + filename)
	if err != nil {
		panic(err)
	}
	return data
}
