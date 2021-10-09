package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"user_credit_card/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewClient(db *gorm.DB, ctx context.Context, config config.Config) *Client {
	log.Println("Connecting to DB....")
	var err error

	uri := fmt.Sprintf("host=%s dbname=%s user=%s port=%s", config.DatabaseHost, config.DatabaseName, config.DatabaseUserName, config.DatabasePort)

	db, err = gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		e := errors.New(fmt.Sprintf("Unable to connect to DB %V", err))
		fmt.Println(e)
	}

	log.Println("Connected to DB..")

	return &Client{
		db: db,
	}

}
