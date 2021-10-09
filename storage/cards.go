package storage

import (
	"context"

	"user_credit_card/types.go"
)

type CardStorage struct {
	client *Client
	ctx    context.Context
}

func NewCardStore(c *Client, ctx context.Context) CardStorage {
	return CardStorage{
		client: c,
		ctx:    ctx,
	}
}

func (cards CardStorage) AutoMigrate(data types.Card) error {
	return cards.client.db.AutoMigrate(&data)
}

func (cards CardStorage) Create(data types.Card) error {
	save := types.Card{
		CardType: data.CardType,
		Number:   data.Number,
		UserID:   data.UserID,
	}
	return cards.client.db.Create(&save).Error
}

func (cards CardStorage) Cards() ([]types.Card, error) {
	var data []types.Card

	return data, cards.client.db.Find(&data).Error
}

func (cards CardStorage) Card(cardN string) (types.Card, error) {
	var data types.Card
	return data, cards.client.db.Find(&data, &types.Card{Number: cardN}).Error
}

func (cards CardStorage) ChangeCradNumber(data types.Card) error {
	return cards.client.db.Model(&types.Card{}).Where("number = ?", data.Number).Update("card_type", data.CardType).Error
}

func (cards CardStorage) SelectUsernameAndCardType() ([]types.User_Card, error) {
	var data []types.User_Card
	return data, cards.client.db.Table("users").Select("users.name,cards.number,cards.card_type").Joins("left join cards on users.id = cards.user_id").Scan(&data).Error

}

func (cards CardStorage) SelectUsernameAndCardTypeUsingId(n string) (types.User_Card, error) {
	var data types.User_Card

	return data, cards.client.db.Table("users").Select("users.name,cards.number,cards.card_type").Joins("left join cards on users.id = cards.user_id").Where("users.name  = ?", n).Scan(&data).Error
}
