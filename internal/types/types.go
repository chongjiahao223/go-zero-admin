// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResq struct {
	Code         string `json:"code"`
	Message      string `json:"message"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	AccessToken  string `json:"accessToken"`
	AccessExipre int64  `json:"accessExipre"`
}
