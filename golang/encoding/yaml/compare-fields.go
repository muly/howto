package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Name string    `yaml:"name"`
	Dev  variables `yaml:"dev"`
	QA   variables `yaml:"qa"`
}

type variables struct {
	Variables map[string]string `yaml:"variables"`
}

func main() {
	f, err := os.Open("./sample1.yaml")
	if err != nil {
		log.Println("error opening file:", err)
		return
	}
	defer f.Close()

	var c config

	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		log.Println("error decoding yaml:", err)
		return
	}
	log.Println(c)

	results := []string{}

	for devK := range c.Dev.Variables {
		_, ok := c.QA.Variables[devK]
		if !ok {
			results = append(results, fmt.Sprintf("%s: variable in dev, missing in qa", devK))
			continue
		}
	}
	// TODO: need to check for duplicate variables in the same env

	fmt.Println("comparision results:")
	for _, r := range results {
		fmt.Println(r)
	}

}
