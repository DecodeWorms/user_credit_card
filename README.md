USER CREDIT CARD PROJECT
This app is built by golang,postgres server technologies,the purpose of this project is to practice one-one tables relation.

The method below fetches datas from both table users and cards and merge them if the users.id = cards.id are the same

storage/cards
func (cards CardStorage) SelectUsernameAndCardType() ([]types.User_Card, error) {
	var data []types.User_Card
	return data, cards.client.db.Table("users").Select("users.name,cards.number,cards.card_type").Joins("left join cards on users.id = cards.user_id").Scan(&data).Error

}

storage/cards
The method below also fetches datas from table users and cards but this "WHERE clause" is involve in order to fetch a single row from both table and merge them using "name" as a constraint

func (cards CardStorage) SelectUsernameAndCardTypeUsingId(n string) (types.User_Card, error) {
	var data types.User_Card

	return data, cards.client.db.Table("users").Select("users.name,cards.number,cards.card_type").Joins("left join cards on users.id = cards.user_id").Where("users.name  = ?", n).Scan(&data).Error
}

And also i scanned the result of query to the DB in to a struct (types.user_card)

NOTE:= Pls if the way i implemented those functions are not good ,kindly correct me.