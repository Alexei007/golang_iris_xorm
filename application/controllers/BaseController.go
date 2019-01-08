package controllers

import "github.com/kataras/iris/sessions"

var Sess *sessions.Sessions

func init() {
	Sess = sessions.New(sessions.Config{Cookie:"GOSESSID",AllowReclaim:true})
}