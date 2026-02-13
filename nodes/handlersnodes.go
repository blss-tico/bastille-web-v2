package nodes

import (
	"bastille-web-v2/config"

	"encoding/json"
	"log"
	"net/http"
)

type HandlersNodes struct{}

func (hn *HandlersNodes) createNodes(w http.ResponseWriter, r *http.Request) {
	log.Println("createNodesHandler")

	var node config.NodesModel
	err := json.NewDecoder(r.Body).Decode(&node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = RegisterNodeToFile(node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Node struct {
		Nodename string `json:"nodename"`
	}

	cNode := Node{Nodename: node.Nodename}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(cNode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hn *HandlersNodes) getNodes(w http.ResponseWriter, r *http.Request) {
	log.Println("getNodesHandler")

	allNodes, err := LoadAllNodesFromFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allNodes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hn *HandlersNodes) updateNodes(w http.ResponseWriter, r *http.Request) {
	log.Println("updateNodesHandler")

	nodename := r.PathValue("nodename")
	var updated config.NodesModel
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = UpdateNodeToFile(nodename, updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (hn *HandlersNodes) deleteNodes(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteNodesHandler")

	nodename := r.PathValue("nodename")
	err := DeleteNodeFromFile(nodename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
