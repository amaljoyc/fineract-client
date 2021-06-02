package util

import (
	"io/ioutil"
)

func Read(filename string) []byte  {
	data, err := ioutil.ReadFile("json/" + filename)
	if err != nil {
		panic(err)
	}
	return data
}
