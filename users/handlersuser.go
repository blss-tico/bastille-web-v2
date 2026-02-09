package users

import (
	"bastille-web-v2/config"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type HandlersUser struct{}

func (hu *HandlersUser) register(w http.ResponseWriter, r *http.Request) {
	log.Println("registerHandler")

	var user config.UsersModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	hashedPassword, _ := config.HashPasswordUtil(user.Password)
	user.Password = hashedPassword
	user.ID = len(config.BwUsers) + 1

	err = RegisterUserToFile(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hu *HandlersUser) login(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandler")

	var creds config.UsersModel
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var storedUser config.UsersModel
	for _, u := range config.BwUsers {
		if u.Username == creds.Username {
			storedUser = u
		}
	}

	if storedUser.Username == "" || !config.CheckPasswordHashUtil(creds.Password, storedUser.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &claimsModel{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessString, err := accessToken.SignedString(config.JwtKeyModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshExpiration := time.Now().Add(24 * time.Hour)
	refreshClaims := &claimsModel{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpiration.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString(config.RefreshKeyModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessCookie := &http.Cookie{
		Name:     "bw-actk",
		Value:    accessString,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	refreshCookie := &http.Cookie{
		Name:     "bw-rftk",
		Value:    refreshString,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
	//w.Write([]byte("JWT set in HttpOnly cookie with SameSite=Strict"))

	err = json.NewEncoder(w).Encode(map[string]string{
		"bw_actk": accessString,
		"bw_rftk": refreshString,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hu *HandlersUser) logout(w http.ResponseWriter, r *http.Request) {
	log.Println("logoutHandler")

	accessCookie := &http.Cookie{
		Name:     "bw-actk",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	refreshCookie := &http.Cookie{
		Name:     "bw-rftk",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
	w.Write([]byte("Logout"))
}

func (hu *HandlersUser) refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("refreshHandler")

	var req map[string]string
	_ = json.NewDecoder(r.Body).Decode(&req)

	refreshToken := req["bw-rftk"]
	claims := &claimsModel{}

	tkn, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return config.RefreshKeyModel, nil
	})

	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := newToken.SignedString(config.JwtKeyModel)

	json.NewEncoder(w).Encode(map[string]string{
		"access_token": tokenString,
	})
}

func (hu *HandlersUser) getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsersHandler")
	json.NewEncoder(w).Encode(config.BwUsers)
}

func (hu *HandlersUser) updateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("updateUserHandler")

	id := r.PathValue("id")
	var updated config.UsersModel
	_ = json.NewDecoder(r.Body).Decode(&updated)

	for i, u := range config.BwUsers {
		if fmt.Sprintf("%d", u.ID) == id {
			config.BwUsers[i].Username = updated.Username
			if updated.Password != "" {
				hashed, _ := config.HashPasswordUtil(updated.Password)
				config.BwUsers[i].Password = hashed
			}
			json.NewEncoder(w).Encode(config.BwUsers[i])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (hu *HandlersUser) deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteUserHandler")

	id := r.PathValue("id")
	for i, u := range config.BwUsers {
		if fmt.Sprintf("%d", u.ID) == id {
			config.BwUsers = append(config.BwUsers[:i], config.BwUsers[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

/*
import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Password string
}

type RefreshToken struct {
    ID        uint   `gorm:"primaryKey"`
    Token     string `gorm:"unique"`
    Username  string
    ExpiresAt int64
}

func initDB() {
    var err error
    db, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&User{}, &RefreshToken{})
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds User
    _ = json.NewDecoder(r.Body).Decode(&creds)

    var storedUser User
    if err := db.Where("username = ?", creds.Username).First(&storedUser).Error; err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    if !checkPasswordHash(creds.Password, storedUser.Password) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(15 * time.Minute)
    claims := &Claims{Username: creds.Username, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString(jwtKey)

    refreshExpiration := time.Now().Add(24 * time.Hour)
    refreshClaims := &Claims{Username: creds.Username, StandardClaims: jwt.StandardClaims{ExpiresAt: refreshExpiration.Unix()}}
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    refreshString, _ := refreshToken.SignedString(refreshKey)

    // Store refresh token in DB
    db.Create(&RefreshToken{Token: refreshString, Username: creds.Username, ExpiresAt: refreshExpiration.Unix()})

    json.NewEncoder(w).Encode(map[string]string{
        "access_token":  tokenString,
        "refresh_token": refreshString,
    })
}

func Refresh(w http.ResponseWriter, r *http.Request) {
    var req map[string]string
    _ = json.NewDecoder(r.Body).Decode(&req)
    refreshToken := req["refresh_token"]

    var stored RefreshToken
    if err := db.Where("token = ?", refreshToken).First(&stored).Error; err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    claims := &Claims{}
    tkn, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
        return refreshKey, nil
    })

    if err != nil || !tkn.Valid || stored.ExpiresAt < time.Now().Unix() {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(15 * time.Minute)
    claims.ExpiresAt = expirationTime.Unix()
    newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := newToken.SignedString(jwtKey)

    json.NewEncoder(w).Encode(map[string]string{"access_token": tokenString})
}

func Logout(w http.ResponseWriter, r *http.Request) {
    var req map[string]string
    _ = json.NewDecoder(r.Body).Decode(&req)
    refreshToken := req["refresh_token"]

    db.Where("token = ?", refreshToken).Delete(&RefreshToken{})
    w.WriteHeader(http.StatusNoContent)
}

*/
