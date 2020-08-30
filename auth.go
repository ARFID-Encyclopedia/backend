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
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passHash"`
	AccessLevel  int    `json:"accessLevel"`
}

type customClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

const (
	ANON = iota
	USER
	MOD
	ADMIN
)

var passHashBytes [20]byte = sha1.Sum([]byte("adminpassword"))

func createToken(userid string) (string, error) {
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

func verifyToken(tokenString string) (string, error) {
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
		return "", errors.New("Couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return "", errors.New("JWT is expired")
	}
	return claims.ID, nil

}

func login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var attemptedCredentials user
	json.Unmarshal(reqBody, &attemptedCredentials)
	log.Printf("attempted Credentials: %v", attemptedCredentials.Username)
	var dbuser user
	log.Printf("looking up user: %v", userMap[attemptedCredentials.Username])
	if err := db.Read("users", userMap[attemptedCredentials.Username], &dbuser); err != nil {
		log.Println("Error user not found", err)
		fmt.Fprintf(w, `{"Error": "User not found, Incorrect username?"}`)
		return
	}
	if attemptedCredentials.Username != dbuser.Username || attemptedCredentials.PasswordHash != dbuser.PasswordHash {
		fmt.Fprintf(w, `{"Error": "Incorrect username or password"}`)
		return
	}
	token, err := createToken(dbuser.ID)
	if err != nil {
		fmt.Fprintf(w, `{"Error": "Error generating token %v"}`, err.Error())
		return
	}
	fmt.Fprint(w, token)
}
