package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"yaml-typed-parser/model"
)

func main() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	config := model.Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(config)
}
