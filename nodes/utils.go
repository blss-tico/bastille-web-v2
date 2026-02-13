package nodes

import (
	"bastille-web-v2/config"

	"log"

	"encoding/json"
	"fmt"
	"os"
)

var fileName = "./nodes.json"

func RegisterNodeToFile(node config.NodesModel) error {
	log.Println("RegisterNodeToFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error register, reading file: %v", err)
	}

	var nodes []config.NodesModel
	if err := json.Unmarshal(fileBytes, &nodes); err != nil {
		return fmt.Errorf("Error register, unmarshalling JSON: %v", err)
	}

	for _, n := range nodes {
		if fmt.Sprintf("%s", n.Nodeip) == node.Nodeip {
			return fmt.Errorf("Error register, node %s with ip %s exists", node.Nodename, node.Nodeip)
		}
	}

	nodes = append(nodes, node)
	updatedBytes, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		return fmt.Errorf("Error register, marshalling JSON: %v", err)
	}

	if err := os.WriteFile(fileName, updatedBytes, 0644); err != nil {
		return fmt.Errorf("Error register, writing file: %v", err)
	}

	return nil
}

func LoadAllNodesFromFile() ([]AllNodes, error) {
	log.Println("LoadAllNodesFromFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return []AllNodes{}, fmt.Errorf("Error load all nodes, reading file: %v", err)
	}

	if len(fileBytes) == 0 {
		return []AllNodes{}, fmt.Errorf("Error load all nodes, file is empty")
	}

	var nodes []config.NodesModel
	if err := json.Unmarshal(fileBytes, &nodes); err != nil {
		return []AllNodes{}, fmt.Errorf("Error load all nodes, unmarshalling JSON: %v", err)
	}

	allNodes := []AllNodes{}
	for _, n := range nodes {
		allNodes = append(allNodes, AllNodes{
			Nodename: n.Nodename,
			Nodeip:   n.Nodeip,
			Nodeport: n.Nodeport,
		})
	}

	return allNodes, nil
}

func UpdateNodeToFile(nodename string, node config.NodesModel) error {
	log.Println("UpdateNodeToFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error update, failed to read file: %w", err)
	}

	if len(fileBytes) == 0 {
		return fmt.Errorf("Error update, nodes file is empty")
	}

	var nodes []config.NodesModel
	if err := json.Unmarshal(fileBytes, &nodes); err != nil {
		return fmt.Errorf("Error update, failed to unmarshal JSON: %w", err)
	}

	log.Println(nodes, node, nodename)
	index := -1
	for i, n := range nodes {
		if fmt.Sprintf("%s", n.Nodename) == nodename {
			nodes[i].Nodename = node.Nodename
			nodes[i].Nodeip = node.Nodeip
			nodes[i].Nodeport = node.Nodeport
			index = i
		}
	}

	if index == -1 {
		return fmt.Errorf("Error update, node %s not found", nodename)
	}

	updatedData, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		return fmt.Errorf("Error update, failed to marshal updated JSON: %w", err)
	}

	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		return fmt.Errorf("Error update, failed to write file: %w", err)
	}

	return nil
}

func DeleteNodeFromFile(nodename string) error {
	log.Println("DeleteUserFromFile")

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error delete, failed to read file: %w", err)
	}

	if len(fileBytes) == 0 {
		return fmt.Errorf("Error delete, nodes file is empty")
	}

	var nodes []config.NodesModel
	if err := json.Unmarshal(fileBytes, &nodes); err != nil {
		return fmt.Errorf("Error delete, failed to unmarshal JSON: %w", err)
	}

	index := -1
	for i, n := range nodes {
		if fmt.Sprintf("%s", n.Nodename) == nodename {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Error delete, user %s not found", nodename)
	}

	nodes = append(nodes[:index], nodes[index+1:]...)
	updatedData, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		return fmt.Errorf("Error delete, failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		return fmt.Errorf("Error delete, failed to write file: %w", err)
	}

	return nil
}
