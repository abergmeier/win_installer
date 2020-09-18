package config

import (
	"io"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func ReadYaml(r io.Reader) []interface{} {

	content, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []interface{}
	err = yaml.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return tasks
}

func ReadYamlFile(filename string) []interface{} {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []interface{}
	err = yaml.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return tasks
}
