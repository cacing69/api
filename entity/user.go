package entity

import "github.com/rocketlaunchr/dbq/v2"

type User struct {
	Id   int64  `json:"id" dbq:"user_id"`
	Name string `json:"name" dbq:"user_name"`
}

func UserSingleOption() *dbq.Options {
	return &dbq.Options{
		ConcreteStruct: User{},
		SingleResult:   true,
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
}
