/*
*  SPDX-License-Identifier: BSD-3-Clause
*
*  Copyright (c) 2025, Bruno Leonardo Tico) <bruno.ccutp@gmail.com>
*  All rights reserved.
*
*  Redistribution and use in source and binary forms, with or without
*  modification, are permitted provided that the following conditions are met:
*
*  * Redistributions of source code must retain the above copyright notice, this
*    list of conditions and the following disclaimer.
*
*  * Redistributions in binary form must reproduce the above copyright notice,
*    this list of conditions and the following disclaimer in the documentation
*    and/or other materials provided with the distribution.
*
*  * Neither the name of the copyright holder nor the names of its
*    contributors may be used to endorse or promote products derived from
*    this software without specific prior written permission.
*
*  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
*  AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
*  IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
*  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
*  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
*  DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
*  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
*  CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
*  OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
*  OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"bastille-web-v2/docs"
)

var bastille bastilleModel

func init() {
	log.Println("init")

	// check bastille installation
	_, err := runBastilleCommands("-v")
	if err != nil {
		log.Fatalln("error: bastille not installed or not found")
	}

	// get bastille-web secret key
	keyModel = os.Getenv("BW_SCRT")
	if keyModel == "" {
		keyModel = "bastille-web-default-secret-key"
	}

	// load bastille definitions
	bastilleFile, err := os.ReadFile("bastille.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(bastilleFile, &bastille)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}
	log.Println("bastille.json ok")

	// load nodes configuration file
	nodesFile, err := os.ReadFile("nodes.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = json.Unmarshal(nodesFile, &nodesListModel)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}

	log.Println("nodes.json ok", nodesListModel)

	if len(addrModel) == 0 {
		ip1 := GetOutboundIPUtil()
		ip2 := GetLocalIPUtil()

		log.Println("ip1[external lookup]: ", ip1)
		log.Println("ip2[loopback ifaces]: ", ip2)

		addrModel = ip2
		docs.SwaggerInfo.Host = addrModel
		log.Println("addrModel: ", addrModel)
		log.Println("portModel: ", portModel)
	}
}

func setAddrAndPort() string {
	log.Println("setAddrAndPort")

	// get bastille-web default ip addr and port
	addr := os.Getenv("BW_ADDR")
	portModel = os.Getenv("BW_PORT")
	log.Println(".env file: ", addr, portModel)

	if addr == "" {
		addr = "0.0.0.0"
	}

	if portModel == "" {
		portModel = "8007"
	}

	log.Printf("address and port %s:%s", addr, portModel)
	return addr + ":" + portModel
}

func startHttpServer(addrPort string) {
	log.Println("startHttpServer")

	mux := http.NewServeMux()
	bastille := &Bastille{}
	handlerTemplates := &HandlersTemplates{bl: *bastille}
	handlersData := &HandlersData{bl: *bastille}
	routes := &Routes{ht: *handlerTemplates, hd: *handlersData}
	routes.staticRoutes(mux)
	routes.swaggerRoutes(mux)
	routes.templatesRoutes(mux)
	routes.dataRoutes(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: false,
		// Debug: true,
	})
	handler := c.Handler(mux)

	log.Printf("Server starting on http://%s", addrPort)
	if err := http.ListenAndServe(addrPort, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// @title Bastille-Web
// @version 1.0
// @description API interface to FreeBSD Bastille Jails Manager.
// @description Observation: Do not use console, edit, htop and top commands with API. Only for UI Interface.
// @termsOfService http://swagger.io/terms/
// @license.name BSD-3-Clause
// @license.url https://opensource.org/license/bsd-3-clause
// @BasePath /
func main() {
	log.Println("main")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	addrPort := setAddrAndPort()
	startHttpServer(addrPort)
}
