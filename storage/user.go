package storage

import (
	"context"

	"user_credit_card/types.go"
)

type UserStorage struct {
	client *Client
	ctx    context.Context
}

func NewUser(c *Client, ct context.Context) UserStorage {
	return UserStorage{
		client: c,
		ctx:    ct,
	}

}

func (store UserStorage) Automigrate(data types.User) error {
	return store.client.db.AutoMigrate(&data)
}

func (store UserStorage) Create(data types.User) error {
	save := types.User{
		Name:   data.Name,
		Gender: data.Gender,
		Age:    data.Age,
	}

	return store.client.db.Create(&save).Error
}

func (store UserStorage) GetRecords() ([]types.User, error) {
	var data []types.User

	return data, store.client.db.Find(&data).Error

}

func (store UserStorage) GetRecord(name string) (types.User, error) {
	var data types.User

	return data, store.client.db.First(&data, "name = ?", name).Error
}

func (store UserStorage) ChangeName(data types.User) error {
	return store.client.db.Model(&types.User{}).Where("id = ?", data.ID).Update("name", data.Name).Error
}
