package iotool

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetDataFromFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}
