package config

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func GetClientIPAddrUtil(r *http.Request) string {
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
