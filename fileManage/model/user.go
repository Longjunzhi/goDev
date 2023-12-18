package model

const UserTableName = "users"

type User struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	UnionId string `json:"union_id"`
}

func NewUser() (user *User) {
	return &User{}
}

func (u *User) TableName() string {
	return UserTableName
}
