package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"user_credit_card/storage"

	"user_credit_card/types.go"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	user storage.UserStorage
}

func NewUserHandler(user storage.UserStorage) UserHandler {
	return UserHandler{
		user: user,
	}
}

func (handler UserHandler) AutoMigrate(w http.ResponseWriter, r *http.Request) {
	var datas types.User

	err := handler.user.Automigrate(datas)

	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to migrate the datas to db table %v", err))
		json.NewEncoder(w).Encode(e)
	}
}

func (handler UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data types.User
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to decode json..%v", err))
		json.NewEncoder(w).Encode(e)

	}
	err = handler.user.Create(data)
	json.NewEncoder(w).Encode(true)
	json.NewEncoder(w).Encode(data)
}

func (handler UserHandler) GetRecords(w http.ResponseWriter, r *http.Request) {

	var data []types.User
	var err error
	data, err = handler.user.GetRecords()
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to fetch records %v", err))
		json.NewEncoder(w).Encode(e)
	}

	json.NewEncoder(w).Encode(data)

}

func (handler UserHandler) GetRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	var data types.User
	var err error

	data, err = handler.user.GetRecord(name)
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to retreive data %v", err))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(data)
}

func (handler UserHandler) ChangeName(w http.ResponseWriter, r *http.Request) {
	var para types.User
	e := json.NewDecoder(r.Body).Decode(&para)
	if e != nil {
		anErr := errors.New(fmt.Sprintf("Unable to decode json %v", e))
		json.NewEncoder(w).Encode(anErr)
	}
	e = handler.user.ChangeName(para)
	json.NewEncoder(w).Encode("Record updated...")

}
