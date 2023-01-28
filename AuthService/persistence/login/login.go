package login

import (
	"auth-service/entity"
)

type LoginService struct{}

func NewUserService() LoginService {
	return LoginService{}
}

func (u *LoginService) SaveUser(login entity.Login) {

}

func (u *LoginService) GetUser() {

}

func (u *LoginService) GetUserByEmailAndPassword() {

}
