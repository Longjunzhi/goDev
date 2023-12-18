package model

const UserMediaTableName = "media_users"

type UserMedia struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	UnionId string `json:"union_id"`
}

func (u *UserMedia) TableName() string {
	return UserMediaTableName
}
