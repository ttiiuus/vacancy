package register

type registerStruct struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewRegister(login string, pass string, id string) *registerStruct {
	return &registerStruct{
		ID:       id,
		Login:    login,
		Password: pass,
	}
}