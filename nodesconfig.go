package main

import (
	"encoding/json"
	"log"
	"os"
)

var nodesFile = "./nodes.json"
var cfg *nodesModel

var Host string
var Ip string
var Port string

func loadNodesConfig() *nodesModel {
	file, err := os.Open(nodesFile)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var nc nodesModel
	if err := json.NewDecoder(file).Decode(&nc); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	cfg = &nc
	return cfg
}

func setNodesConfig(config *nodesModel) {
	Host = config.Host
	Ip = config.Ip
	Port = config.Port
}
