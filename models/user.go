package models

type User struct {
	ID         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Address    string `json:"address" form:"address"`
	Occupation string `json:"occupation" form:"occupation"`
	Phone      string `json:"phone" form:"phone"`
}
