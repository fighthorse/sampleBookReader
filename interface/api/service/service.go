package service

import (
	"github.com/fighthorse/sampleBookReader/interface/api/service/amap"
	"github.com/fighthorse/sampleBookReader/interface/api/service/book"
	"github.com/fighthorse/sampleBookReader/interface/api/service/login"
	"github.com/fighthorse/sampleBookReader/interface/api/service/member"
)

var LoginService *login.Service
var BookService *book.Service
var AmapService *amap.Service
var MemberService *member.Service

func InitService() (err error) {
	LoginService, err = login.New()
	if err != nil {
		panic(err)
	}
	BookService, err = book.New()
	if err != nil {
		panic(err)
	}
	AmapService, err = amap.New()
	if err != nil {
		panic(err)
	}
	MemberService, err = member.New()
	if err != nil {
		panic(err)
	}
	return
}
