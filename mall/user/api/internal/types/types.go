// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserAddReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type UserAddRes struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserDelReq struct {
	Id int64 `json:"id"`
}

type UserDelRes struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type UserEditReq struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type UserEditRes struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type UserQueryReq struct {
	Username string `json:"username"`
}

type UserQueryRes struct {
	UserList []*User `json:"products"`
}