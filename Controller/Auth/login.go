package auth

import (
	"encoding/json"
	"jwt-auth/Model/User"
)

type UserRespon struct {
	Message string
	Token   string
}

type Auth struct {
	User *User.User
}

func Init(U *User.User) *Auth {
	return &Auth{
		User: U,
	}
}

func (A *Auth) Login(lReq LoginRequest) ([]byte, error) {
	var u = User.User{
		Username: lReq.Username,
		Password: lReq.Password,
	}
	user, err := A.User.GetUserByUserName(u)
	if err != nil {
		return []byte{}, err
	}
	if user.Password != lReq.Password {
		return []byte{}, err
	}

	var token = CreateToken(lReq)
	var urespon = UserRespon{
		Message: "success",
		Token:   token,
	}
	res, _ := json.Marshal(urespon)
	return res, nil
}
