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

type CardHandler struct {
	card storage.CardStorage
}

func NewCardHandler(c storage.CardStorage) CardHandler {
	return CardHandler{
		card: c,
	}
}

func (cd CardHandler) AutoMigrate(w http.ResponseWriter, r *http.Request) {
	var data types.Card

	err := cd.card.AutoMigrate(data)

	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to migrate datas to DB table.. %V", err))
		json.NewEncoder(w).Encode(e)
	}
}

func (cd CardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data types.Card
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to decode json %v", err))
		json.NewEncoder(w).Encode(e)

	}

	err = cd.card.Create(data)
	json.NewEncoder(w).Encode(err)
}

func (cd CardHandler) Card(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	num := params["number"]
	var data types.Card
	var err error
	data, err = cd.card.Card(num)

	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to retrieve data %v", err))
		json.NewEncoder(w).Encode(e)
	}

	json.NewEncoder(w).Encode(data)
}

func (cd CardHandler) Cards(w http.ResponseWriter, r *http.Request) {
	var data []types.Card
	var err error

	data, err = cd.card.Cards()
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to retrieve records from DB.. %v", err))
		json.NewEncoder(w).Encode(e)
	}

	json.NewEncoder(w).Encode(data)
}

func (cd CardHandler) ChangeCradNumber(w http.ResponseWriter, r *http.Request) {
	var err error
	var data types.Card
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to decode json %v", err))
		json.NewEncoder(w).Encode(e)
	}

	err = cd.card.ChangeCradNumber(data)
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to decode json %v", err))
		json.NewEncoder(w).Encode(e)
	}

	json.NewEncoder(w).Encode("Successfully updated..")
}

func (cd CardHandler) SelectUsernameAndCardType(w http.ResponseWriter, r *http.Request) {
	var data []types.User_Card
	var err error

	data, err = cd.card.SelectUsernameAndCardType()
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to retreive data %v", data))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode(data)

}

func (cd CardHandler) SelectUsernameAndCardTypeUsingId(w http.ResponseWriter, r *http.Request) {
	var err error
	// var d types.User
	// err = json.NewDecoder(r.Body).Decode(&d)
	params := mux.Vars(r)
	nam := params["name"]
	var data types.User_Card

	data, err = cd.card.SelectUsernameAndCardTypeUsingId(nam)
	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to fetch data from server %v", err))
		json.NewEncoder(w).Encode(e)
	}
	json.NewEncoder(w).Encode("Fetching datas..")
	json.NewEncoder(w).Encode(data)
}
