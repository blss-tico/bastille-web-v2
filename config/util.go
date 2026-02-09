package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserToFile(user UsersModel) error {
	fileBytes, err := os.ReadFile("users.json")
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	var users []UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	users = append(users, user)
	updatedBytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}

	if err := os.WriteFile(string(fileBytes), updatedBytes, 0644); err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}

func UpdateUserToFile(user UsersModel) error {
	fileBytes, err := os.ReadFile("users.json")
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var users []UsersModel
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for i, u := range users {
		if fmt.Sprintf("%d", u.ID) == strconv.Itoa(user.ID) {
			users[i].Username = user.Username
			if user.Password != "" {
				hashed, _ := HashPasswordUtil(user.Password)
				users[i].Password = hashed
			}
		}
	}

	updatedData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal updated JSON: %w", err)
	}

	if err := os.WriteFile("users.json", updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

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

func HashPasswordUtil(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHashUtil(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
