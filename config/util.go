package config

import (
	"log"
	"net"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

	var ip string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println("ip4: ", ipnet.IP.String())
				ip = ipnet.IP.String()
			}

			if ipnet.IP.To16() != nil {
				log.Println("ip6: ", ipnet.IP.String())
			}
		}
	}

	return ip
}

func HashPasswordUtil(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHashUtil(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
