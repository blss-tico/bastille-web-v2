package users

import (
	"bastille-web-v2/config"
	"crypto/rand"

	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oklog/ulid/v2"
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

	if user.Username == "admin" {
		http.Error(w, "error: user admin cannot be created", http.StatusInternalServerError)
		return
	}

	hashedPassword, _ := config.HashPasswordUtil(user.Password)
	user.Password = hashedPassword
	entropy := ulid.Monotonic(rand.Reader, 0)
	id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	user.ID = id.String()

	err = RegisterUserToFile(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type User struct {
		Id string `json:"id"`
	}

	cUser := User{Id: id.String()}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(cUser)
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

	err = LoadUserFromFile(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)
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
		Name:     "bw-session",
		Value:    accessString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, accessCookie)

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
		Name:     "bw-session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	}

	http.SetCookie(w, accessCookie)

	err := json.NewEncoder(w).Encode(map[string]string{
		"bw_actk": "",
		"bw_rftk": "",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hu *HandlersUser) refreshTkApi(w http.ResponseWriter, r *http.Request) {
	log.Println("refreshTkApiHandler")

	var req map[string]string
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken := req["bw_rftk"]
	claims := &claimsModel{}
	tkn, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return config.RefreshKeyModel, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessString, _ := newToken.SignedString(config.JwtKeyModel)

	json.NewEncoder(w).Encode(map[string]string{
		"bw_actk": accessString,
	})
}

func (hu *HandlersUser) getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsersHandler")

	allUsers, err := LoadAllUserFromFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hu *HandlersUser) updateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("updateUserHandler")

	username := r.PathValue("username")
	var updated config.UsersModel
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = UpdateUserToFile(username, updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (hu *HandlersUser) deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteUserHandler")

	username := r.PathValue("username")
	err := DeleteUserFromFile(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
