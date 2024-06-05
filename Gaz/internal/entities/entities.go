package entities

import "time"

type User struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Birthday string `json:"birthday" db:"birthday"`
}

type UserForDb struct {
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	Birthday time.Time `json:"birthday" db:"birthdate"`
}

type SubscribeReq struct {
	Username string `json:"username"`
}

type BirthDay struct {
	UserWithBirth string `db:"username"`
	Email         string `db:"email"`
}
