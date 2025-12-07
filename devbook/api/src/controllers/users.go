package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewUsersRepository(db)
	userId, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Fprintf(w, "new user created, Id: %d", userId)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
