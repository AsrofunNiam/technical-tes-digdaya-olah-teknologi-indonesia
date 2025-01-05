package service

import "github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/web"

type UserService interface {
	Login(identity, password, userAgent, remoteAddress *string, request *web.UserLoginRequest) web.TokenResponse
}
