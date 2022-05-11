package util

import (
	"io/ioutil"
	"log"
)

var root_path = "./scripts/"

func LoadSQL(name string) string {
	body, err := ioutil.ReadFile(root_path + name)
	if err != nil {
		log.Fatalf("Unable to load %s", name)
	}
	return string(body)
}
