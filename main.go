package main

import (
	"bastille-web-v2/api"
	"bastille-web-v2/bastille"
	"bastille-web-v2/config"
	"bastille-web-v2/nodes"
	"bastille-web-v2/users"
	"bastille-web-v2/web"

	"log"
	"net/http"

	"github.com/rs/cors"
)

func init() {
	log.Println("init")

	_, err := bastille.RunBastilleCommands("-v")
	if err != nil {
		log.Fatalln("error: bastille not installed on machine or not found")
	}
}

func startHttpServer() {
	log.Println("startHttpServer")

	config.LoadConfigParams()
	addrPort := config.BwAddrModel + ":" + config.BwPortModel

	mux := http.NewServeMux()
	bastille := bastille.NewBastille()

	handlersData := &api.HandlersData{Bl: *bastille}
	apiRoutes := api.NewRoutes(*handlersData)
	apiRoutes.SwaggerRoutes(mux)
	apiRoutes.DataRoutes(mux)

	handlerTemplates := &web.HandlersTemplates{Bl: *bastille}
	webRoutes := web.NewRoutes(*handlerTemplates)
	webRoutes.StaticRoutes(mux)
	webRoutes.TemplatesRoutes(mux)

	handlerUsers := &users.HandlersUser{}
	userRoutes := users.NewRoutes(*handlerUsers)
	userRoutes.UserRoutes(mux)

	handlerNodes := &nodes.HandlersNodes{}
	nodesRoutes := nodes.NewRoutes(*handlerNodes)
	nodesRoutes.NodesRoutes(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
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
	startHttpServer()
}
