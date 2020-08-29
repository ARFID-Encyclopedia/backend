package main

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type slice []string

type user struct {
	ID           uint64 `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passHash"`
	AccessLevel  int    `json:"accessLevel"`
}

type customClaims struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}

const (
	ANON = iota
	USER
	MOD
	ADMIN
)

var passHashBytes [20]byte = sha1.Sum([]byte("adminpassword"))

func createToken(userid uint64) (string, error) {
	var err error
	atClaims := customClaims{
		ID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "ARFID Encyclopedia",
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyToken(tokenString string) (uint64, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		},
	)
	if err != nil {
		log.Fatal("token error: ", err)
	}
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return 0, errors.New("Couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return 0, errors.New("JWT is expired")
	}
	return claims.ID, nil

}

func login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user user
	json.Unmarshal(reqBody, &user)
	if user.Username != users[1].Username || user.PasswordHash != users[1].PasswordHash {
		fmt.Fprintf(w, `{"Error": "Incorrect username or password"}`)
		return
	}
	token, err := createToken(users[1].ID)
	if err != nil {
		fmt.Fprintf(w, `{"Error": "Error generating token %v"}`, err.Error())
		return
	}
	fmt.Fprint(w, token)
}
