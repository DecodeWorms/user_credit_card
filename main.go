package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"user_credit_card/config"
	"user_credit_card/handlers"
	"user_credit_card/storage"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var client *storage.Client
var use storage.UserStorage
var card storage.CardStorage

func init() {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	var db *gorm.DB

	cfg := config.Config{
		DatabaseHost:     host,
		DatabaseName:     dbname,
		DatabaseUserName: user,
		DatabasePort:     port,
	}
	ctx := context.Background()
	client = storage.NewClient(db, ctx, cfg)
	use = storage.NewUser(client, ctx)
	card = storage.NewCardStore(client, ctx)

}

func main() {
	userHandler := handlers.NewUserHandler(use)
	cardHandler := handlers.NewCardHandler(card)
	router := mux.NewRouter()
	router.HandleFunc("/user/table", userHandler.AutoMigrate).Methods("POST")
	router.HandleFunc("/user/create", userHandler.Create).Methods("POST")
	router.HandleFunc("/user/firstrecord", userHandler.GetRecords).Methods("GET")
	router.HandleFunc("/user/records", userHandler.GetRecords).Methods("GET")
	router.HandleFunc("/user/record/{name}", userHandler.GetRecord).Methods("GET")
	router.HandleFunc("/user/update", userHandler.ChangeName).Methods("PUT")

	router.HandleFunc("/card/table", cardHandler.AutoMigrate).Methods("POST")
	router.HandleFunc("/card/create", cardHandler.Create).Methods("POST")
	router.HandleFunc("/card/cards/{number}", cardHandler.Card).Methods("GET")
	router.HandleFunc("/card/cards", cardHandler.Cards).Methods("GET")
	router.HandleFunc("/card/user_card", cardHandler.SelectUsernameAndCardType).Methods("GET")
	router.HandleFunc("/card/user_cards/{name}", cardHandler.SelectUsernameAndCardTypeUsingId).Methods("GET")
	router.HandleFunc("/card/cardupdate", cardHandler.ChangeCradNumber).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
