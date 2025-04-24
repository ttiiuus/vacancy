package login

import (
	"fmt"
)

type LoginStruct struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewLogin(login string, pass string, id string) *LoginStruct {
	return &LoginStruct{
		ID:       id,
		Login:    login,
		Password: pass,
	}
}

func (v *LoginStruct) PrintVacancy() {
	fmt.Println(v)
}
