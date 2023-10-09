package api

import "Go-Mall/AppV1/api/user_api"

type Api struct {
	UserApi user_api.UserApi
}

var ApiGroup = new(Api)
