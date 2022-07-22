package api

import (
	"fmt"
	"reflect"

	db "github.com/cyanrad/university/db/sqlc"
	"github.com/cyanrad/university/util"
	"github.com/golang/mock/gomock"
)

type eqCreateUserMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUser(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserMatcher{arg, password}
}
