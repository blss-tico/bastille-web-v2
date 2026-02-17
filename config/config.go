package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfigParams() {
	log.Println("loadConfigParams")

	LoadEnvVarsFile()
	LoadUsersFile()
	LoadBastilleFile()
	LoadNodesFile()
}

func LoadEnvVarsFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	BwAddrModel = os.Getenv("BW_ADDR")
	BwPortModel = os.Getenv("BW_PORT")

	if BwAddrModel == "" {
		BwAddrModel = "0.0.0.0"
	}

	if BwPortModel == "" {
		BwPortModel = "8007"
	}
	log.Printf("address and port %s:%s", BwAddrModel, BwPortModel)

	// get bastille-web secret key
	KeyModel = os.Getenv("BW_SCRT")
	if KeyModel == "" {
		log.Fatalln("error: bastille-web secret key not loaded")
	}

	// get bastille-web jwt key
	JwtKeyModel = []byte(os.Getenv("BW_JWT_KEY"))
	if len(JwtKeyModel) == 0 {
		log.Fatalln("error: bastille-web jwt secret key not loaded")
	}

	// get bastille-web refresh jwt key
	RefreshKeyModel = []byte(os.Getenv("BW_REFRESH_KEY"))
	if len(RefreshKeyModel) == 0 {
		log.Fatalln("error: bastille-web refresh jwt secret key not loaded")
	}
}

func LoadUsersFile() {
	usersFile, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(usersFile, &BwUsers)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}
	log.Println("users.json ok")
}

func LoadBastilleFile() {
	bastilleFile, err := os.ReadFile("bastille.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(bastilleFile, &BastilleM)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}
	log.Println("bastille.json ok")
}

func LoadNodesFile() {
	nodesFile, err := os.ReadFile("nodes.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if len(nodesFile) > 0 {
		err = json.Unmarshal(nodesFile, &NodesListModel)
		if err != nil {
			log.Fatalf("error unmarshaling JSON: %v", err)
		}
		log.Println("nodes.json ok", NodesListModel)
	} else {
		log.Println("nodes.json is empty or incorrect", NodesListModel)
	}
}
