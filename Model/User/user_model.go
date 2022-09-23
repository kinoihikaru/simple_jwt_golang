package User

import (
	db "jwt-auth/Db"
)

type User struct {
	Con      *db.DB
	Id       int
	Username string
	Password string
}

func Init(con *db.DB) *User {
	return &User{
		Con: con,
	}
}

func (U *User) GetUserByUserName(user User) (User, error) {
	err := U.Con.GetConnection().Db.QueryRow("select * from users where username = ?", user.Username).Scan(&user.Id, &user.Username, &user.Password)

	return user, err
}
