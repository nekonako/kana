package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Token string `json:"token"`
}

func newConfig() *config {
	return new(config)
}

func readConfig() (config *config, err error) {

	fmt.Println("reading config....")

	file, err := ioutil.ReadFile("./config.json")
	returnIfErr(err)

	config = newConfig()

	err = json.Unmarshal(file, &config)
	returnIfErr(err)
	fmt.Println(config.Token)

	return

}
