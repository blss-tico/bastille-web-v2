package config

import (
	"bastille-web-v2/bastille"
	"bastille-web-v2/docs"

	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfigParams() {
	log.Println("loadConfigParams")

	// check bastille installation
	_, err := bastille.RunBastilleCommands("-v")
	if err != nil {
		log.Fatalln("error: bastille not installed on machine or not found")
	}

	err = godotenv.Load(".env")
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

	// load bastille users
	usersFile, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(usersFile, &BwUsers)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}
	log.Println("users.json ok")

	// load bastille definitions
	bastilleFile, err := os.ReadFile("bastille.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(bastilleFile, &BastilleM)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}
	log.Println("bastille.json ok")

	// load nodes configuration file
	nodesFile, err := os.ReadFile("nodes.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(nodesFile, &NodesListModel)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}
	log.Println("nodes.json ok", NodesListModel)

	ip1 := GetOutboundIPUtil()
	log.Println("ip1[external lookup]: ", ip1)

	ip2 := GetLocalIPUtil()
	log.Println("ip2[loopback ifaces]: ", ip2)

	AddrModel = ip2
	docs.SwaggerInfo.Host = AddrModel
	log.Println("addrModel: ", AddrModel)
}
