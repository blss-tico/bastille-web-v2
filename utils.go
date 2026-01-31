package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand/v2"
	"net"
	"net/http"
	"strconv"
	"strings"
)

var templates *template.Template

func renderTemplateUtil(w http.ResponseWriter, name string, data any) {
	log.Println("renderTemplate")

	tmplt := "./templates/" + name
	files := []string{
		"./templates/base.html",
		"./templates/partials/navbar.html",
		"./templates/partials/sidebar.html",
	}
	files = append(files, tmplt)

	var err error
	templates, err = template.ParseFiles(files...)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func respondOkWithJSONUtil(w http.ResponseWriter, payload string) {
	log.Println("respondOkWithJSONUtil")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if payload != "" {
		response := map[string]string{"msg": payload}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]string{"msg": "command executed"}
	json.NewEncoder(w).Encode(response)
}

func RespondErrorWithJSONUtil(w http.ResponseWriter, code int, payload string) {
	log.Println("RespondErrorWithJSONUtil")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	response := map[string]string{"msg": "with errors", "err": payload}
	json.NewEncoder(w).Encode(response)
}

func RandPortUtil() string {
	log.Println("RandPortUtil")
	return strconv.Itoa(8000 + rand.IntN(8200-8000))
}

func getClientIPAddrUtil(r *http.Request) string {
	log.Println("getClientIPAddrUtil")

	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		if header := r.Header.Get(h); header != "" {
			ips := strings.Split(header, ",")
			log.Println("ips: ", ips)
			return strings.TrimSpace(ips[0])
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return ip
}

func GetOutboundIPUtil() string {
	log.Println("GetOutboundIPUtil")

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func GetLocalIPUtil() string {
	log.Println("GetLocalIPUtil")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
