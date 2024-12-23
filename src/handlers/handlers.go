package handlers

import (
	"fmt"
	"my-go-web-server/src/connections"
	"my-go-web-server/src/models"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	dbName := os.Getenv("DB_NAME")
	message := fmt.Sprintf("Welcome to my Go web server! DB_NAME is %s", dbName)
	w.Write([]byte(message))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	db := connections.ConnectToDB()
	var user models.DBUser
	db.First(&user, "Name = ?", username)
	if user.Name == "" {
		user = models.DBUser{Name: username}
		fmt.Fprintf(w, "Hello, %s!", username)
		db.Create(&user)
	} else {
		fmt.Fprintf(w, "Hello, %s! Didn't see you since %s!", username, user.CreatedAt.Format("2006-01-02 15:04:05"))
		db.Save(&user)
	}
}
